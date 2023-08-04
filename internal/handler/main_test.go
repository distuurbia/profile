package handler

import (
	"os"
	"testing"

	"github.com/distuurbia/profile/internal/model"
	protocol "github.com/distuurbia/profile/protocol/profile"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

var (
	validate    *validator.Validate
	testProfile = model.Profile{
		ID:           uuid.New(),
		Username:     "Vladimir",
		Password:     []byte("1234"),
		RefreshToken: []byte("someToken"),
		Country:      "Belarus",
		Age:          27,
	}
	testProtoProfile = protocol.Profile{
		Username:     testProfile.Username,
		Password:     []byte("1234"),
		RefreshToken: []byte("someToken"),
		Country:      "Belarus",
		Age:          27,
	}
)

func TestMain(m *testing.M) {
	validate = validator.New()
	exitCode := m.Run()
	os.Exit(exitCode)
}
