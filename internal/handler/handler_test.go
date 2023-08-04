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

func TestCreateProfile(t *testing.T) {
	s := new(mocks.ProfileService)

	s.On("CreateProfile", mock.Anything, mock.AnythingOfType("*model.Profile")).Return(nil)

	h := NewProfileHandler(s, validate)
	_, err := h.CreateProfile(context.Background(), &protocol.CreateProfileRequest{Profile: &testProtoProfile})

	require.NoError(t, err)
}

func TestCreateProfileFailedValidation(t *testing.T) {
	s := new(mocks.ProfileService)
	h := NewProfileHandler(s, validate)

	testProtoProfile.Username = "f"
	_, err := h.CreateProfile(context.Background(), &protocol.CreateProfileRequest{Profile: &testProtoProfile})

	require.Error(t, err)
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

func TestGetPasswordAndIDByEmptyUsername(t *testing.T) {
	s := new(mocks.ProfileService)

	h := NewProfileHandler(s, validate)

	_, err := h.GetPasswordAndIDByUsername(context.Background(), &protocol.GetPasswordAndIDByUsernameRequest{Username: ""})
	require.Error(t, err)
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

func TestGetRefreshTokenByIDNil(t *testing.T) {
	s := new(mocks.ProfileService)
	h := NewProfileHandler(s, validate)

	_, err := h.GetRefreshTokenByID(context.Background(), &protocol.GetRefreshTokenByIDRequest{Id: uuid.Nil.String()})
	require.Error(t, err)
}

func TestAddRefreshToken(t *testing.T) {
	s := new(mocks.ProfileService)

	s.On("AddRefreshToken", mock.Anything, mock.AnythingOfType("[]uint8"), mock.AnythingOfType("uuid.UUID")).
		Return(nil)

	h := NewProfileHandler(s, validate)

	_, err := h.AddRefreshToken(context.Background(), &protocol.AddRefreshTokenRequest{Id: testProfile.ID.String()})
	require.NoError(t, err)
}

func TestAddRefreshTokenNilID(t *testing.T) {
	s := new(mocks.ProfileService)

	s.On("AddRefreshToken", mock.Anything, mock.AnythingOfType("[]uint8"), mock.AnythingOfType("uuid.UUID")).
		Return(nil)

	h := NewProfileHandler(s, validate)

	_, err := h.AddRefreshToken(context.Background(), &protocol.AddRefreshTokenRequest{Id: uuid.Nil.String()})
	require.Error(t, err)
}

func TestDeleteProfile(t *testing.T) {
	s := new(mocks.ProfileService)

	s.On("DeleteProfile", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil)

	h := NewProfileHandler(s, validate)

	_, err := h.DeleteProfile(context.Background(), &protocol.DeleteProfileRequest{Id: uuid.New().String()})

	require.NoError(t, err)
}

func TestDeleteProfileSendNotUUID(t *testing.T) {
	s := new(mocks.ProfileService)

	s.On("DeleteProfile", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil)

	h := NewProfileHandler(s, validate)

	_, err := h.DeleteProfile(context.Background(), &protocol.DeleteProfileRequest{Id: "not uuid"})

	require.Error(t, err)
}

func TestDeleteProfileNilID(t *testing.T) {
	s := new(mocks.ProfileService)

	s.On("DeleteProfile", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil)

	h := NewProfileHandler(s, validate)

	_, err := h.DeleteProfile(context.Background(), &protocol.DeleteProfileRequest{Id: uuid.Nil.String()})

	require.Error(t, err)
}
