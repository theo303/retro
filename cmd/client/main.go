package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/theo303/retro"
	"google.golang.org/protobuf/proto"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/"}
	slog.Info(fmt.Sprintf("connecting to %s", u.String()))

	header := http.Header{}
	header.Add(retro.UserNameHeaderKey, "cli-client")
	c, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		slog.Error("dial error", slog.Any("error", err))
		os.Exit(1)
	}
	defer c.Close()

	done := make(chan struct{})

	s := new(retro.State)
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				slog.Error("read error", slog.Any("error", err))
				return
			}
			if err := proto.Unmarshal(message, s); err != nil {
				slog.Error("proto unmarshal error", slog.Any("error", err))
				return
			}
			slog.Info(fmt.Sprintf("%+v", s))
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			slog.Info("interrupt")

			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				slog.Error("write error", slog.Any("error", err))
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
