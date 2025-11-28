package http_server

import (
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"diov.local/krasivo/internal/configuration"
	"diov.local/krasivo/internal/http-server/router"
)

func Start(c *configuration.Configuration, log *slog.Logger) {
	r := router.Start()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{
		Addr:         c.Address,
		Handler:      &r,
		ReadTimeout:  c.HTTPServer.Timeout,
		WriteTimeout: c.HTTPServer.Timeout,
		IdleTimeout:  c.HTTPServer.IdleTimeout,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Error("failed to start server")
			log.Error(err.Error())
		}
	}()

	log.Info("server started")

	<-done
	log.Info("stopping server")
}
