package retro

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/exp/rand"
	"google.golang.org/protobuf/proto"
)

const UserNameHeaderKey = "X-User-Name"
const maxUserNameLength = 20

var upgrader = websocket.Upgrader{}

type Server struct {
	*http.Server
	conns map[string]*websocket.Conn

	tickRate time.Duration

	state *State
}

func NewServer(addr string, tickRate time.Duration, state *State) *Server {
	if state == nil {
		state = new(State)
	}
	s := &Server{
		Server: &http.Server{
			Addr: addr,
		},
		conns:    make(map[string]*websocket.Conn),
		tickRate: tickRate,
		state:    state,
	}
	http.HandleFunc("/", s.connect)
	return s
}

func (s *Server) Run(ctx context.Context) {
	go func() {
		if err := s.ListenAndServe(); err != nil {
			slog.Info("server closed", slog.Any("error", err))
		}
	}()

	ticker := time.NewTicker(s.tickRate)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			data, err := proto.Marshal(s.state)
			if err != nil {
				slog.Error("proto marshal error", slog.Any("error", err))
			}
			for userID, conn := range s.conns {
				if err := conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
					slog.Error("write message error", slog.Any("error", err), slog.Any("userID", userID))
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
	id := randString()
	s.conns[id] = c
	s.state.Users = append(s.state.Users, &User{Name: name, Id: id})
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

	userName := r.Header.Get(UserNameHeaderKey)
	userName = strings.TrimSpace(userName)
	userName = userName[:min(maxUserNameLength, len(userName))]
	id := s.addConn(c, userName)
	slog.Info("client connected", slog.Any("id", id), slog.Any("userName", userName))

	defer delete(s.conns, id)

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
				slog.Error("read error", slog.Any("error", err))
			}
			slog.Info("client disconnected", slog.Any("id", id), slog.Any("userName", userName))
			return
		}
		slog.Info("message read", slog.Any("messageType", mt), slog.Any("message", fmt.Sprintf("%v", msg)))
	}
}
