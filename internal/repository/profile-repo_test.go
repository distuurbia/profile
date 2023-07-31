package repository

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestSignUp(t *testing.T) {
	err := r.SignUp(context.Background(), &testProfile)
	require.NoError(t, err)
}

func TestSignUpNilID(t *testing.T){
	testProfile.ID = uuid.Nil
	err := r.SignUp(context.Background(), &testProfile)
	require.Error(t, err)
}

func TestSignUpExistingID(t *testing.T){
	_ = r.SignUp(context.Background(), &testProfile)
	testProfile.Username = "smth"
	err := r.SignUp(context.Background(), &testProfile)
	require.Error(t, err)
}


func TestSignUpExistingUsername(t *testing.T) {
	testProfile.ID = uuid.New()
	testProfile.Username = "Vladimir"
	_ = r.SignUp(context.Background(), &testProfile)
	testProfile.ID = uuid.New()
	err := r.SignUp(context.Background(), &testProfile)
	require.Error(t, err)
}

