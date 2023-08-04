package repository

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateProfile(t *testing.T) {
	err := r.CreateProfile(context.Background(), &testProfile)
	require.NoError(t, err)
}

func TestCreateProfileNilID(t *testing.T) {
	testProfile.ID = uuid.Nil
	err := r.CreateProfile(context.Background(), &testProfile)
	require.Error(t, err)
}

func TestCreateProfileExistingID(t *testing.T) {
	_ = r.CreateProfile(context.Background(), &testProfile)
	testProfile.Username = "Vova"
	err := r.CreateProfile(context.Background(), &testProfile)
	require.Error(t, err)
}

func TestCreateProfileExistingUsername(t *testing.T) {
	testProfile.ID = uuid.New()
	testProfile.Username = "Vladimir"
	_ = r.CreateProfile(context.Background(), &testProfile)
	testProfile.ID = uuid.New()
	err := r.CreateProfile(context.Background(), &testProfile)
	require.Error(t, err)
}

func TestGetPasswordAndIDByUsername(t *testing.T) {
	testProfile.ID = uuid.New()
	testProfile.Username = "Volodya"
	err := r.CreateProfile(context.Background(), &testProfile)
	require.NoError(t, err)
	testID, testPsw, err := r.GetPasswordAndIDByUsername(context.Background(), testProfile.Username)
	require.NoError(t, err)
	require.Equal(t, testID, testProfile.ID)
	require.Equal(t, testPsw, testPsw)
}

func TestGetPasswordAndIDByNotExistingUsername(t *testing.T) {
	_, _, err := r.GetPasswordAndIDByUsername(context.Background(), "Nobody")
	require.Error(t, err)
}

func TestGetRefreshTokenByID(t *testing.T) {
	testProfile.ID = uuid.New()
	testProfile.Username = "Vlados"
	err := r.CreateProfile(context.Background(), &testProfile)
	require.NoError(t, err)
	hashedRefresh, err := r.GetRefreshTokenByID(context.Background(), testProfile.ID)
	require.NoError(t, err)
	require.Equal(t, testProfile.RefreshToken, hashedRefresh)
}

func TestGetRefreshTokenByNotExistingID(t *testing.T) {
	_, err := r.GetRefreshTokenByID(context.Background(), uuid.New())
	require.Error(t, err)
}

func TestAddRefreshToken(t *testing.T) {
	testProfile.ID = uuid.New()
	testProfile.Username = "Vladlen"
	err := r.CreateProfile(context.Background(), &testProfile)
	require.NoError(t, err)

	newRT := []byte("NewRT")
	err = r.AddRefreshToken(context.Background(), newRT, testProfile.ID)
	require.NoError(t, err)

	testRT, err := r.GetRefreshTokenByID(context.Background(), testProfile.ID)
	require.NoError(t, err)
	require.Equal(t, newRT, testRT)
}

func TestAddRefreshTokenWithNotExistingID(t *testing.T) {
	newRT := []byte("NewRT")
	err := r.AddRefreshToken(context.Background(), newRT, uuid.New())
	require.Error(t, err)
}

func TestDeleteProfile(t *testing.T) {
	testProfile.ID = uuid.New()
	testProfile.Username = "Volodmir"
	err := r.CreateProfile(context.Background(), &testProfile)
	require.NoError(t, err)

	err = r.DeleteProfile(context.Background(), testProfile.ID)
	require.NoError(t, err)
}

func TestDeleteProfileWithNotExistingID(t *testing.T) {
	err := r.DeleteProfile(context.Background(), uuid.New())
	require.Error(t, err)
}
