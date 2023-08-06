package repository

import (
	"context"
	"testing"

	"github.com/distuurbia/profile/internal/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateProfile(t *testing.T) {
	err := r.CreateProfile(context.Background(), &testProfile)
	require.NoError(t, err)

	var readProfile model.Profile
	err = r.pool.QueryRow(context.Background(), "SELECT username, password, refreshToken, age, country FROM profiles WHERE ID = $1",
		testProfile.ID).Scan(&readProfile.Username, &readProfile.Password, &readProfile.RefreshToken, &readProfile.Age, &readProfile.Country)
	require.NoError(t, err)
	require.Equal(t, testProfile.Username, readProfile.Username)
	require.Equal(t, testProfile.Password, readProfile.Password)
	require.Equal(t, testProfile.RefreshToken, readProfile.RefreshToken)
	require.Equal(t, testProfile.Age, readProfile.Age)
	require.Equal(t, testProfile.Country, readProfile.Country)
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

func TestGetRefreshTokenByID(t *testing.T) {
	testProfile.ID = uuid.New()
	testProfile.Username = "Vlados"
	err := r.CreateProfile(context.Background(), &testProfile)
	require.NoError(t, err)
	hashedRefresh, err := r.GetRefreshTokenByID(context.Background(), testProfile.ID)
	require.NoError(t, err)
	require.Equal(t, testProfile.RefreshToken, hashedRefresh)
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

func TestDeleteProfile(t *testing.T) {
	testProfile.ID = uuid.New()
	testProfile.Username = "Volodmir"
	err := r.CreateProfile(context.Background(), &testProfile)
	require.NoError(t, err)

	err = r.DeleteProfile(context.Background(), testProfile.ID)
	require.NoError(t, err)
}
