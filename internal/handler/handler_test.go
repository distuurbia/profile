package handler

import (
	"context"
	"testing"

	"github.com/distuurbia/profile/internal/handler/mocks"
	protocol "github.com/distuurbia/profile/protocol/profile"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestValidationID(t *testing.T) {
	s := new(mocks.ProfileService)
	h := NewProfileHandler(s, validate)

	testID := uuid.New()
	validatedID, err := h.ValidationID(context.Background(), testID.String())
	require.NoError(t, err)
	require.Equal(t, testID, validatedID)
}

func TestCreateProfile(t *testing.T) {
	s := new(mocks.ProfileService)

	s.On("CreateProfile", mock.Anything, mock.AnythingOfType("*model.Profile")).Return(nil)

	h := NewProfileHandler(s, validate)
	_, err := h.CreateProfile(context.Background(), &protocol.CreateProfileRequest{Profile: &testProtoProfile})

	require.NoError(t, err)
}

func TestGetPasswordAndIDByUsername(t *testing.T) {
	s := new(mocks.ProfileService)

	s.On("GetPasswordAndIDByUsername", mock.Anything, mock.AnythingOfType("string")).
		Return(testProfile.ID, []byte("pass"), nil)

	h := NewProfileHandler(s, validate)

	resp, err := h.GetPasswordAndIDByUsername(context.Background(), &protocol.GetPasswordAndIDByUsernameRequest{Username: testProfile.Username})
	require.NoError(t, err)
	require.Equal(t, len([]byte("pass")), len(resp.Password))
	require.Equal(t, testProfile.ID.String(), resp.Id)
}

func TestGetRefreshTokenByID(t *testing.T) {
	s := new(mocks.ProfileService)

	s.On("GetRefreshTokenByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).
		Return([]byte("token"), nil)

	h := NewProfileHandler(s, validate)

	resp, err := h.GetRefreshTokenByID(context.Background(), &protocol.GetRefreshTokenByIDRequest{Id: testProfile.ID.String()})
	require.NoError(t, err)
	require.Equal(t, len([]byte("token")), len(resp.HashedRefresh))
}

func TestAddRefreshToken(t *testing.T) {
	s := new(mocks.ProfileService)

	s.On("AddRefreshToken", mock.Anything, mock.AnythingOfType("[]uint8"), mock.AnythingOfType("uuid.UUID")).
		Return(nil)

	h := NewProfileHandler(s, validate)

	_, err := h.AddRefreshToken(context.Background(), &protocol.AddRefreshTokenRequest{Id: testProfile.ID.String()})
	require.NoError(t, err)
}

func TestDeleteProfile(t *testing.T) {
	s := new(mocks.ProfileService)

	s.On("DeleteProfile", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil)

	h := NewProfileHandler(s, validate)

	_, err := h.DeleteProfile(context.Background(), &protocol.DeleteProfileRequest{Id: uuid.New().String()})

	require.NoError(t, err)
}
