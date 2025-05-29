package main

import (
	syslog "log"

	"github.com/Makovey/booking_auth/internal/app"
	"github.com/Makovey/booking_utils/pkg/config/env"
	"github.com/Makovey/booking_utils/pkg/logger/slog"
)

func main() {
	fn := "main"

	log := slog.NewLogger()
	cfg, err := env.NewAuthConfig(log)
	if err != nil {
		syslog.Fatalf("[%s]: %w", fn, err)
	}

	appl := app.NewApp(log, cfg)
	if err = appl.Run(); err != nil {
		syslog.Fatalf("[%s]: %v", fn, err)
	}
}
