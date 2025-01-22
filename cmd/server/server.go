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

	s := retro.NewServer(":8080", 1*time.Second, &retro.State{
		Users: map[string]*retro.User{
			"1": {
				Name: "Raymond",
			},
		},
		Stickies: map[string]*retro.Sticky{
			"idsticky": {
				Owner:   "1",
				X:       989898,
				Y:       -121212,
				Content: "je contient du texte",
			},
		}})

	go s.Run(ctx)
	slog.Info("server started")

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		slog.Error("server shutdown error", slog.Any("error", err))
	}
}
