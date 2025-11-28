package database

import (
	"context"
	"fmt"
	"log"

	"diov.local/krasivo/internal/configuration"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool
var ctx = context.Background()

func Init(c configuration.Configuration) {

	var err error
	pool, err = pgxpool.New(ctx, c.DatabaseUrl)

	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	// Verify the connection
	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Unable to ping database:", err)
	}

	fmt.Println("Connected to PostgreSQL database!")
}
