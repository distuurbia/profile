package service

import (
	"context"
	"testing"

	"github.com/distuurbia/profile/internal/service/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateProfile(t *testing.T) {
	r := new(mocks.ProfileRepository)
	r.On("CreateProfile", mock.Anything, mock.AnythingOfType("*model.Profile")).Return(nil)

	s := NewProfileService(r, &cfg)

	err := s.CreateProfile(context.Background(), &testProfile)
	require.NoError(t, err)
}

func TestGetPasswordAndIDByUsername(t *testing.T) {
	r := new(mocks.ProfileRepository)

	r.On("GetPasswordAndIDByUsername", mock.Anything, mock.AnythingOfType("string")).
		Return(testProfile.ID, []byte("pass"), nil)

	s := NewProfileService(r, &cfg)

	profileID, hashedPassword, err := s.GetPasswordAndIDByUsername(context.Background(), testProfile.Username)
	require.NoError(t, err)
	require.Equal(t, len([]byte("pass")), len(hashedPassword))
	require.Equal(t, testProfile.ID, profileID)
}

func TestGetRefreshTokenByID(t *testing.T) {
	r := new(mocks.ProfileRepository)

	r.On("GetRefreshTokenByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).
		Return([]byte("token"), nil)

	s := NewProfileService(r, &cfg)

	hashedToken, err := s.GetRefreshTokenByID(context.Background(), testProfile.ID)
	require.NoError(t, err)
	require.Equal(t, len([]byte("token")), len(hashedToken))
}

func TestAddRefreshToken(t *testing.T) {
	r := new(mocks.ProfileRepository)

	r.On("AddRefreshToken", mock.Anything, mock.AnythingOfType("[]uint8"), mock.AnythingOfType("uuid.UUID")).
		Return(nil)

	s := NewProfileService(r, &cfg)

	err := s.AddRefreshToken(context.Background(), testProfile.RefreshToken, testProfile.ID)
	require.NoError(t, err)
}

func TestDeleteProfile(t *testing.T) {
	r := new(mocks.ProfileRepository)
	r.On("DeleteProfile", mock.Anything, mock.AnythingOfType("uuid.UUID")).
		Return(nil)
	s := NewProfileService(r, &cfg)
	err := s.DeleteProfile(context.Background(), uuid.New())
	require.NoError(t, err)
}
