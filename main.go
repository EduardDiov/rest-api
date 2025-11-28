package main

import (
	"log/slog"

	"diov.local/krasivo/internal/configuration"
	"diov.local/krasivo/internal/database"
	"diov.local/krasivo/internal/logger"

	http_server "diov.local/krasivo/internal/http-server"
)

func main() {
	c := configuration.MustLoad()
	log := logger.Start(c.Env)

	// database.ExampleDB(c.DatabaseUrl)

	log.Info(
		"starting bot: krasivo",
		slog.String("env", c.Env),
		slog.String("version", "0.1"),
	)

	database.Init(*c)

	http_server.Start(c, log)

	// client := bybit.NewBybitHttpClient(c.Security.Key, c.Security.Secret, bybit.WithBaseURL(bybit.MAINNET))
	// market.GetServerTime(client)
}
