package retro

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	sync "sync"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/exp/rand"
	"google.golang.org/protobuf/proto"
)

const (
	UserNameHeaderKey = "X-User-Name"
	maxUserNameLength = 20
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(_ *http.Request) bool {
		return true
	},
}

type serverState struct {
	*State
	lastHeight int64
	mutex      sync.Mutex
}

type Server struct {
	*http.Server
	conns map[string]*websocket.Conn

	tickRate time.Duration

	serverState
}

func NewServer(addr string, tickRate time.Duration, state *State) *Server {
	if state == nil {
		state = new(State)
	}
	s := &Server{
		Server: &http.Server{
			Addr: addr,
		},
		conns:       make(map[string]*websocket.Conn),
		tickRate:    tickRate,
		serverState: serverState{State: state},
	}
	http.HandleFunc("/connect", s.connect)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	return s
}

func (s *Server) Run(ctx context.Context) {
	go func() {
		if err := s.ListenAndServe(); err != nil {
			slog.Info("server closed", slog.Any("error", err))
		}
	}()

	ticker := time.NewTicker(s.tickRate)
	userIDTicker := time.NewTicker(s.tickRate * 100)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			data, err := proto.Marshal(s.State)
			if err != nil {
				slog.Error("proto marshal error", slog.Any("error", err))
			}
			for userID, conn := range s.conns {
				if err := conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
					slog.Error("write state message error", slog.Any("error", err), slog.Any("userID", userID))
				}
			}
		case <-userIDTicker.C:
			for userID, conn := range s.conns {
				if err := conn.WriteMessage(websocket.TextMessage, []byte(userID)); err != nil {
					slog.Error("write userID message error", slog.Any("error", err), slog.Any("userID", userID))
				}
			}
		}
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (s *Server) addConn(c *websocket.Conn, name string) string {
	id := "u-" + randString()
	s.conns[id] = c
	s.State.Users[id] = &User{Name: name}
	return id
}

func (s *Server) connect(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "connection error")
		slog.Error("upgrade connection error", slog.Any("error", err))
		return
	}
	defer c.Close()

	userName := r.URL.Query().Get("name")
	userName = strings.TrimSpace(userName)
	userName = userName[:min(maxUserNameLength, len(userName))]

	userID := s.addConn(c, userName)
	log := slog.With(slog.Any("userID", userID), slog.Any("userName", userName))
	log.Info("client connected")

	defer delete(s.conns, userID)

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
				log.Error("read error", slog.Any("error", err))
			}
			log.Info("client disconnected")
			return
		}
		if mt == websocket.BinaryMessage {
			a := new(Action)
			if err := proto.Unmarshal(msg, a); err != nil {
				log.Error("proto unmarshal error", slog.Any("error", err))
				continue
			}
			if err := s.applyAction(a, userID); err != nil {
				log.Warn("apply action error", slog.Any("error", err))
				continue
			}
		} else {
			log.Warn("unsupported message type read", slog.Any("messageType", mt), slog.Any("message", fmt.Sprintf("%v", msg)))
		}
	}
}

func (s *Server) applyAction(a *Action, userID string) error {
	if s.State == nil {
		return errors.New("state is nil")
	}
	if a == nil {
		return errors.New("action is nil")
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	switch a.Action.(type) {
	case *Action_Select:
		selectAction := a.Action.(*Action_Select).Select
		sticky, ok := s.State.Stickies[selectAction.StickyID]
		if !ok {
			return errors.New("select: sticky not found")
		}
		if err := selection(s.State, sticky, selectAction.StickyID, userID); err != nil {
			return fmt.Errorf("select: %w", err)
		}
	case *Action_Add:
		addAction := a.Action.(*Action_Add).Add
		stickyID := "s-" + randString()
		s.lastHeight++
		sticky := &Sticky{
			X:      addAction.X,
			Y:      addAction.Y,
			Height: s.lastHeight,
		}
		s.State.Stickies[stickyID] = sticky
		if err := selection(s.State, sticky, stickyID, userID); err != nil {
			return fmt.Errorf("add: %w", err)
		}
	case *Action_Move:
		moveAction := a.Action.(*Action_Move).Move
		sticky, ok := s.State.Stickies[moveAction.StickyID]
		if !ok {
			return errors.New("move: sticky not found")
		}
		if sticky.SelectedBy == nil {
			if err := selection(s.State, sticky, moveAction.StickyID, userID); err != nil {
				return fmt.Errorf("move: %w", err)
			}
		} else if *sticky.SelectedBy != userID {
			return fmt.Errorf("move: sticky is selected by another user (%s)", *sticky.SelectedBy)
		}
		sticky.X = moveAction.X
		sticky.Y = moveAction.Y
	case *Action_Edit:
		editAction := a.Action.(*Action_Edit).Edit
		sticky, ok := s.State.Stickies[editAction.StickyID]
		if !ok {
			return errors.New("edit: sticky not found")
		}
		if sticky.SelectedBy == nil {
			if err := selection(s.State, sticky, editAction.StickyID, userID); err != nil {
				return fmt.Errorf("move: %w", err)
			}
		} else if *sticky.SelectedBy != userID {
			return fmt.Errorf("edit: sticky is selected by another user (%s)", *sticky.SelectedBy)
		}
		sticky.Content = editAction.Content
	default:
		return fmt.Errorf("unknown action type (%T)", a.Action)
	}
	return nil
}

func selection(s *State, sticky *Sticky, stickyID, userID string) error {
	user, ok := s.Users[userID]
	if !ok {
		return errors.New("user not found")
	}
	if user.HasSelected != nil {
		selected, ok := s.Stickies[*user.HasSelected]
		if ok {
			selected.SelectedBy = nil
		}
	}
	sticky.SelectedBy = &userID
	user.HasSelected = &stickyID
	return nil
}
