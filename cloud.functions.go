package Wakeups

import (
	"WebService/config"
	"WebService/db"
	"WebService/db/migrations"
	"WebService/server"
	"context"
	"log"
)

var s *server.Server

func init() {
	ctx := context.Background()

	cfg := config.Get()

	pgDB, err := db.Dial(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := migrations.Run(pgDB); err != nil {
		log.Fatal(err)
	}

	s = server.Init(ctx, cfg, pgDB)
}
