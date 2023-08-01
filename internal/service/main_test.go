package service

import (
	"os"
	"testing"

	"github.com/caarlos0/env"
	"github.com/distuurbia/profile/internal/config"
	"github.com/distuurbia/profile/internal/model"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	cfg config.Config
	testProfile = model.Profile{
		ID: uuid.New(),
		Password: []byte("password"),
		RefreshToken: []byte("refreshToken"),
		Username: "Volodya",
		Country: "Belarus",
		Age: 27,

	}
)
func TestMain(m *testing.M){
	if err := env.Parse(&cfg); err != nil {
		logrus.Fatalf("main -> %v", err)
	}
	exitCode := m.Run()
	os.Exit(exitCode)
}