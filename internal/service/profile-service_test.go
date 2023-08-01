package service

import (
	"context"
	"crypto/sha256"
	"testing"

	"github.com/distuurbia/profile/internal/model"
	"github.com/distuurbia/profile/internal/service/mocks"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestHashBytes(t *testing.T) {
	r := new(mocks.ProfileRepository)
	s := NewProfileService(r, &cfg)

	testBytes := []byte("test")
	_, err := s.HashBytes(testBytes)

	require.NoError(t, err)
}

func TestHashEmptyBytes(t *testing.T) {
	r := new(mocks.ProfileRepository)
	s := NewProfileService(r, &cfg)

	testBytes := make([]byte, 0)
	_, err := s.HashBytes(testBytes)

	require.Error(t, err)
}

func TestCompareHashAndBytes(t *testing.T) {
	r := new(mocks.ProfileRepository)
	s := NewProfileService(r, &cfg)

	testBytes := []byte("test")
	hashedBytes, err := s.HashBytes(testBytes)
	require.NoError(t, err)

	isEqual, err := s.CompareHashAndBytes(hashedBytes, testBytes)
	require.NoError(t, err)

	require.True(t, isEqual)
}

func TestCompareDifferentHashAndBytes(t *testing.T) {
	r := new(mocks.ProfileRepository)
	s := NewProfileService(r, &cfg)

	testBytes := []byte("test")
	hashedBytes, err := s.HashBytes(testBytes)
	require.NoError(t, err)

	testAnotherBytes := []byte("anotherTest")
	isEqual, err := s.CompareHashAndBytes(hashedBytes, testAnotherBytes)
	require.Error(t, err)

	require.True(t, !isEqual)
}

func TestGenerateJWT(t *testing.T) {
	r := new(mocks.ProfileRepository)
	s := NewProfileService(r, &cfg)

	_, err := s.GenerateJWTToken(accessTokenExpiration, uuid.New())
	require.NoError(t, err)
}

func TestGenerateJWTNilID(t *testing.T) {
	r := new(mocks.ProfileRepository)
	s := NewProfileService(r, &cfg)

	_, err := s.GenerateJWTToken(accessTokenExpiration, uuid.Nil)
	require.Error(t, err)
}

func TestGenerateTokenPair(t *testing.T) {
	r := new(mocks.ProfileRepository)
	s := NewProfileService(r, &cfg)

	_, err := s.GenerateTokenPair(uuid.New())
	require.NoError(t, err)
}

func TestGenerateTokenPairNilID(t *testing.T) {
	r := new(mocks.ProfileRepository)
	s := NewProfileService(r, &cfg)

	_, err := s.GenerateTokenPair(uuid.Nil)
	require.Error(t, err)
}

func TestValidateToken(t *testing.T) {
	r := new(mocks.ProfileRepository)
	s := NewProfileService(r, &cfg)

	token, err := s.GenerateJWTToken(accessTokenExpiration, uuid.New())
	require.NoError(t, err)

	_, err = s.ValidateToken(token)
	require.NoError(t, err)
}

func TestValidateTokenWithAnotherSignKey(t *testing.T) {
	r := new(mocks.ProfileRepository)
	s := NewProfileService(r, &cfg)

	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("anotherKey"))
	require.NoError(t, err)
	_, err = s.ValidateToken(tokenString)
	require.Error(t, err)
}

func TestSignUp(t *testing.T) {
	r := new(mocks.ProfileRepository)
	r.On("SignUp", mock.Anything, mock.AnythingOfType("*model.Profile")).Return(nil)

	s := NewProfileService(r, &cfg)

	err := s.SignUp(context.Background(), &testProfile)
	require.NoError(t, err)
}

func TestSignUpProfileNil(t *testing.T) {
	r := new(mocks.ProfileRepository)
	r.On("SignUp", mock.Anything, mock.AnythingOfType("*model.Profile")).Return(nil)

	s := NewProfileService(r, &cfg)

	err := s.SignUp(context.Background(), nil)
	require.Error(t, err)
}

func TestSignUpProfileNilID(t *testing.T) {
	r := new(mocks.ProfileRepository)
	r.On("SignUp", mock.Anything, mock.AnythingOfType("*model.Profile")).Return(nil)

	s := NewProfileService(r, &cfg)

	testProfile.ID = uuid.Nil
	err := s.SignUp(context.Background(), &testProfile)
	require.Error(t, err)
}

func TestSignUpProfileEmptyUsername(t *testing.T) {
	r := new(mocks.ProfileRepository)
	r.On("SignUp", mock.Anything, mock.AnythingOfType("*model.Profile")).Return(nil)

	s := NewProfileService(r, &cfg)

	testProfile.ID = uuid.New()
	testProfile.Username = ""
	err := s.SignUp(context.Background(), &testProfile)
	testProfile.Username = "Volodya"
	require.Error(t, err)
}

func TestLogin(t *testing.T){
	r := new(mocks.ProfileRepository)

	hashedbytes, err := bcrypt.GenerateFromPassword(testProfile.Password, bcryptCost)
	require.NoError(t, err)

	r.On("GetPasswordAndIDByUsername", mock.Anything, mock.AnythingOfType("string")).
	Return(testProfile.ID, hashedbytes, nil)
	r.On("AddRefreshToken", mock.Anything, mock.AnythingOfType("[]uint8"), mock.AnythingOfType("uuid.UUID")).
	Return(nil)

	s := NewProfileService(r, &cfg)

	_, err = s.Login(context.Background(), testProfile.Username, testProfile.Password)
	require.NoError(t, err)
}

func TestLoginEmptyUsername(t *testing.T){
	r := new(mocks.ProfileRepository)

	hashedbytes, err := bcrypt.GenerateFromPassword(testProfile.Password, bcryptCost)
	require.NoError(t, err)

	r.On("GetPasswordAndIDByUsername", mock.Anything, mock.AnythingOfType("string")).
	Return(testProfile.ID, hashedbytes, nil)
	r.On("AddRefreshToken", mock.Anything, mock.AnythingOfType("[]uint8"), mock.AnythingOfType("uuid.UUID")).
	Return(nil)

	s := NewProfileService(r, &cfg)

	_, err = s.Login(context.Background(), "", testProfile.Password)
	require.Error(t, err)
}

func TestLoginEmptyPassword(t *testing.T){
	r := new(mocks.ProfileRepository)

	hashedbytes, err := bcrypt.GenerateFromPassword(testProfile.Password, bcryptCost)
	require.NoError(t, err)

	r.On("GetPasswordAndIDByUsername", mock.Anything, mock.AnythingOfType("string")).
	Return(testProfile.ID, hashedbytes, nil)
	r.On("AddRefreshToken", mock.Anything, mock.AnythingOfType("[]uint8"), mock.AnythingOfType("uuid.UUID")).
	Return(nil)

	s := NewProfileService(r, &cfg)
	testPassword := make([]byte, 0)
	_, err = s.Login(context.Background(), testProfile.Username, testPassword)
	require.Error(t, err)
}

func TestTokensIDCompare(t *testing.T) {
	r := new(mocks.ProfileRepository)
	s := NewProfileService(r, &cfg)
	tokenPair, err := s.GenerateTokenPair(testProfile.ID)
	require.NoError(t, err)
	id, err := s.TokensIDCompare(tokenPair)
	require.NoError(t, err)
	require.Equal(t, testProfile.ID, id)
}

func TestTokensIDCompareDifferentIDs(t *testing.T){
	r := new(mocks.ProfileRepository)
	s := NewProfileService(r, &cfg)

	var tokenPair = model.TokenPair{}
	var err error
	tokenPair.AccessToken, err = s.GenerateJWTToken(accessTokenExpiration, uuid.New())
	require.NoError(t, err)
	tokenPair.RefreshToken, err = s.GenerateJWTToken(refreshTokenExpiration, uuid.New())
	require.NoError(t, err)
	_, err = s.TokensIDCompare(&tokenPair)
	require.Error(t, err)
}

func TestRefresh(t *testing.T){
	r := new(mocks.ProfileRepository)
	s := NewProfileService(r, &cfg)

	tokenPair, err := s.GenerateTokenPair(testProfile.ID)
	require.NoError(t, err)
	sum := sha256.Sum256([]byte(tokenPair.RefreshToken))

	hashedbytes, err := bcrypt.GenerateFromPassword(sum[:], bcryptCost)
	require.NoError(t, err)

	r.On("GetRefreshTokenByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).
	Return(hashedbytes, nil)
	r.On("AddRefreshToken", mock.Anything, mock.AnythingOfType("[]uint8"), mock.AnythingOfType("uuid.UUID")).
	Return(nil)

	_, err = s.Refresh(context.Background(), tokenPair)
	require.NoError(t, err)
}

func TestRefreshTokenPairNil(t *testing.T){
	r := new(mocks.ProfileRepository)
	s := NewProfileService(r, &cfg)

	_, err := s.Refresh(context.Background(), nil)
	require.Error(t, err)
}

func TestRefreshTokenPairEmpty(t *testing.T){
	r := new(mocks.ProfileRepository)
	s := NewProfileService(r, &cfg)
	tokenPair := model.TokenPair{
		AccessToken: "",
		RefreshToken: "",
	}
	_, err := s.Refresh(context.Background(), &tokenPair)
	require.Error(t, err)
}

