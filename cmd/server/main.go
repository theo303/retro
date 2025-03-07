package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/theo303/retro"
)

func main() {
	ctx := context.Background()
	ctx, _ = signal.NotifyContext(ctx, os.Interrupt)

	s := retro.NewServer(":8080", 10*time.Millisecond, &retro.State{
		Users: map[string]*retro.User{
			"1": {
				Name: "Raymond",
			},
		},
		Stickies: map[string]*retro.Sticky{
			"idsticky": {
				Owner:   "1",
				X:       100,
				Y:       100,
				Content: "je contient du texte11",
			},
			"idsticky2": {
				Owner:   "1",
				X:       200,
				Y:       200,
				Content: "je contient du texte22",
			},
		},
	})

	go s.Run(ctx)
	slog.Info("server started")

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		slog.Error("server shutdown error", slog.Any("error", err))
	}
}
