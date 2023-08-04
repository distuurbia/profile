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

func TestCreateProfileNilProfile(t *testing.T) {
	r := new(mocks.ProfileRepository)
	s := NewProfileService(r, &cfg)

	err := s.CreateProfile(context.Background(), nil)
	require.Error(t, err)
}

func TestCreateProfileNilID(t *testing.T) {
	r := new(mocks.ProfileRepository)
	r.On("CreateProfile", mock.Anything, mock.AnythingOfType("*model.Profile")).Return(nil)

	s := NewProfileService(r, &cfg)

	testProfile.ID = uuid.Nil
	err := s.CreateProfile(context.Background(), &testProfile)
	require.Error(t, err)
}

func TestCreateProfileEmptyUsername(t *testing.T) {
	r := new(mocks.ProfileRepository)
	r.On("CreateProfile", mock.Anything, mock.AnythingOfType("*model.Profile")).Return(nil)

	s := NewProfileService(r, &cfg)

	testProfile.ID = uuid.New()
	testProfile.Username = ""
	err := s.CreateProfile(context.Background(), &testProfile)
	testProfile.Username = "Volodya"
	require.Error(t, err)
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

func TestGetPasswordAndIDByEmptyUsername(t *testing.T) {
	r := new(mocks.ProfileRepository)

	r.On("GetPasswordAndIDByUsername", mock.Anything, mock.AnythingOfType("string")).
		Return(testProfile.ID, nil, nil)

	s := NewProfileService(r, &cfg)

	_, _, err := s.GetPasswordAndIDByUsername(context.Background(), "")
	require.Error(t, err)
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

func TestGetRefreshTokenByIDNil(t *testing.T) {
	r := new(mocks.ProfileRepository)
	s := NewProfileService(r, &cfg)

	_, err := s.GetRefreshTokenByID(context.Background(), uuid.Nil)
	require.Error(t, err)
}

func TestAddRefreshToken(t *testing.T) {
	r := new(mocks.ProfileRepository)

	r.On("AddRefreshToken", mock.Anything, mock.AnythingOfType("[]uint8"), mock.AnythingOfType("uuid.UUID")).
		Return(nil)

	s := NewProfileService(r, &cfg)

	err := s.AddRefreshToken(context.Background(), testProfile.RefreshToken, testProfile.ID)
	require.NoError(t, err)
}

func TestAddRefreshTokenNilToken(t *testing.T) {
	r := new(mocks.ProfileRepository)

	r.On("AddRefreshToken", mock.Anything, mock.AnythingOfType("[]uint8"), mock.AnythingOfType("uuid.UUID")).
		Return(nil)

	s := NewProfileService(r, &cfg)

	err := s.AddRefreshToken(context.Background(), nil, testProfile.ID)
	require.Error(t, err)
}

func TestAddRefreshTokenNilID(t *testing.T) {
	r := new(mocks.ProfileRepository)

	r.On("AddRefreshToken", mock.Anything, mock.AnythingOfType("[]uint8"), mock.AnythingOfType("uuid.UUID")).
		Return(nil)

	s := NewProfileService(r, &cfg)

	err := s.AddRefreshToken(context.Background(), testProfile.RefreshToken, uuid.Nil)
	require.Error(t, err)
}

func TestDeleteProfile(t *testing.T) {
	r := new(mocks.ProfileRepository)
	r.On("DeleteProfile", mock.Anything, mock.AnythingOfType("uuid.UUID")).
		Return(nil)
	s := NewProfileService(r, &cfg)
	err := s.DeleteProfile(context.Background(), uuid.New())
	require.NoError(t, err)
}

func TestDeleteProfileNilID(t *testing.T) {
	r := new(mocks.ProfileRepository)
	s := NewProfileService(r, &cfg)

	err := s.DeleteProfile(context.Background(), uuid.Nil)
	require.Error(t, err)
}
