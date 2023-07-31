package main

import (
	"context"
	"fmt"

	"github.com/caarlos0/env"
	"github.com/distuurbia/profile/internal/config"
	"github.com/distuurbia/profile/internal/model"
	"github.com/distuurbia/profile/internal/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func connectPostgres(cfg *config.Config) (*pgxpool.Pool, error) {
	conf, err := pgxpool.ParseConfig(cfg.PostgresPath)
	if err != nil {
		return nil, fmt.Errorf("error in method pgxpool.ParseConfig: %v", err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		return nil, fmt.Errorf("error in method pgxpool.NewWithConfig: %v", err)
	}
	return pool, nil
}

func main() {
	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		logrus.Fatalf("main -> %v", err)
	}
	pool, err := connectPostgres(&cfg)
	if err != nil {
		logrus.Fatalf("main -> %v", err)
	}
	r := repository.NewProfileRepository(pool)
	prof := model.Profile{
		ID: uuid.New(),
		Username: "vladimir",
		Password: []byte("1234"),
	}
	r.SignUp(context.Background(), &prof)
}