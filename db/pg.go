package db

import (
	"time"

	"WebService/config"
	"github.com/go-pg/pg/v9"
	_ "github.com/lib/pq"
)

const Timeout = 5

type PgDB struct {
	*pg.DB
}

func Dial(cfg *config.Config) (*PgDB, error) {
	pgOpts := &pg.Options{}

	var err error

	if cfg.PgURL != "" {
		pgOpts, err = pg.ParseURL(cfg.PgURL)
		if err != nil {
			return nil, err
		}
	}
	if cfg.PgProto != "" {
		pgOpts.Network = cfg.PgProto
	}
	if cfg.PgAddr != "" {
		pgOpts.Addr = cfg.PgAddr
	}
	if cfg.PgDb != "" {
		pgOpts.Database = cfg.PgDb
	}
	if cfg.PgUser != "" {
		pgOpts.User = cfg.PgUser
	}
	if cfg.PgPassword != "" {
		pgOpts.Password = cfg.PgPassword
	}

	pgDB := pg.Connect(pgOpts)

	_, err = pgDB.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}
	if Timeout > 0 {
		pgDB.WithTimeout(time.Second * time.Duration(Timeout))
	}

	return &PgDB{pgDB}, nil
}
