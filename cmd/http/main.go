package main

import (
	"context"
	"log"

	"ca/cmd/http/app"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func main() {
	// gracefull shutdown

	cfg := NewConfig()

	r := chi.NewRouter()

	ctx := context.Background()

	db, err := pgx.Connect(ctx, cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close(context.Background())

	// init logger slog

	app.Run(
		r,
		db,
		// cfg,
		// logger
	)
}
