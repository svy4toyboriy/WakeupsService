package config

import (
	"sync"
)

type Config struct {
	PgURL      string
	PgProto    string
	PgAddr     string
	PgDb       string
	PgUser     string
	PgPassword string
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	return &Config{
		PgURL:      "postgres://postgres:postgres@localhost:5432/info?sslmode=disable",
		PgProto:    "tcp",
		PgAddr:     "localhost:5432",
		PgDb:       "Info",
		PgUser:     "postgres",
		PgPassword: "12345",
	}
}
