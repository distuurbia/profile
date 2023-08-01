// Package service contains the bisnes logic of app
package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/distuurbia/profile/internal/config"
	"github.com/distuurbia/profile/internal/model"
	"github.com/google/uuid"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Expiration time for an access and a refresh tokens and bcryptCost
const (
	accessTokenExpiration  = 15 * time.Minute
	refreshTokenExpiration = 72 * time.Hour
	bcryptCost             = 14
)

// ProfileRepository is an interface of repository.ProfileRepository and contains its methods
type ProfileRepository interface {
	SignUp(ctx context.Context, profile *model.Profile) error
	GetPasswordAndIDByUsername(ctx context.Context, username string) (id uuid.UUID, password []byte, err error)
	GetRefreshTokenByID(ctx context.Context, id uuid.UUID) (hashedRefresh []byte, err error)
	AddRefreshToken(ctx context.Context, refreshToken []byte, id uuid.UUID) error
	DeleteProfile(ctx context.Context, id uuid.UUID) error
}

// ProfileService contains an object of ProfileRepository and config with env variables
type ProfileService struct {
	r ProfileRepository
	cfg *config.Config
}

// NewProfileService creates *ProfileSevice object filles it and returns
func NewProfileService(r ProfileRepository, cfg *config.Config) *ProfileService {
	return &ProfileService{r: r, cfg: cfg}
}

// HashBytes makes from bytes hashed value
func (s *ProfileService) HashBytes(bytes []byte) ([]byte, error) {
	if len(bytes) == 0{
		return nil, fmt.Errorf("error: gained bytes are empty")
	}
	hashedbytes, err := bcrypt.GenerateFromPassword(bytes, bcryptCost)
	if err != nil {
		return nil, fmt.Errorf("GenerateFromPassword: %w", err)
	}
	return hashedbytes, nil
}

// CompareHashAndBytes compares given hashedBytes and bytes
func (s *ProfileService) CompareHashAndBytes(hashedBytes, bytes []byte) (bool, error) {
	if len(hashedBytes) == 0 || len(bytes) == 0{
		return false, fmt.Errorf("error: gained bytes are empty")
	}
	err := bcrypt.CompareHashAndPassword(hashedBytes, bytes)
	if err != nil {
		return false, fmt.Errorf("CompareHashAndPassword: %w", err)
	}
	return true, nil
}

// GenerateJWTToken generates JWT token with given expiration and profile id
func (s *ProfileService) GenerateJWTToken(expiration time.Duration, id uuid.UUID) (string, error) {
	if id == uuid.Nil {
		return "", fmt.Errorf("error: id is nil")
	}
	claims := &jwt.MapClaims{
		"exp": time.Now().Add(expiration).Unix(),
		"id":  id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.cfg.SecretKey))
	if err != nil {
		return "", fmt.Errorf("token.SignedString: %w", err)
	}
	return tokenString, nil
}

// GenerateTokenPair generates pair of access and refresh tokens
func (s *ProfileService) GenerateTokenPair(id uuid.UUID) (*model.TokenPair, error) {
	if id == uuid.Nil {
		return nil, fmt.Errorf("ProfileService -> GenerateTokenPair -> error: id is nil")
	}
	accessToken, err := s.GenerateJWTToken(accessTokenExpiration, id)
	if err != nil {
		return nil, fmt.Errorf("ProfileService -> GenerateTokenPair -> accessToken -> %w", err)
	}
	refreshToken, err := s.GenerateJWTToken(refreshTokenExpiration, id)
	if err != nil {
		return nil, fmt.Errorf("ProfileService -> GenerateTokenPair -> refreshToken -> %w", err)
	}
	return &model.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// SignUp hashes password and send profile to SignUp method of ProfileRepository
func (s *ProfileService) SignUp(ctx context.Context, profile *model.Profile) (err error) {
	if profile == nil || profile.ID == uuid.Nil || len(profile.Username) == 0 {
		return fmt.Errorf("ProfileService -> GenerateTokenPair -> error: profile is empty")
	}
	profile.Password, err = s.HashBytes(profile.Password)
	if err != nil {
		return fmt.Errorf("ProfileService -> HashPassword -> %w", err)
	}
	err = s.r.SignUp(ctx, profile)
	if err != nil {
		return fmt.Errorf("ProfileService -> %w", err)
	}
	return nil
}

// Login checks if profile with such username exists, compare given password and hashed password in db then generates access and refresh tokens
func (s *ProfileService) Login(ctx context.Context, username string, password []byte) (*model.TokenPair, error){
	if len(username) == 0 || len(password) == 0{
		return nil, fmt.Errorf("ProfileService -> Login -> error: username or password is empty")
	}
	id, hashedPassword, err := s.r.GetPasswordAndIDByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("ProfileService ->  Login -> %w", err)
	}
	verified, err := s.CompareHashAndBytes(hashedPassword, password)
	if err != nil || !verified {
		return nil, fmt.Errorf("ProfileService ->  Login -> %w", err)
	}
	tokenPair, err := s.GenerateTokenPair(id)
	if err != nil {
		return nil, fmt.Errorf("ProfileService ->  Login -> %w", err)
	}
	sum := sha256.Sum256([]byte(tokenPair.RefreshToken))
	hashedRefreshToken, err := s.HashBytes(sum[:])
	if err != nil {
		return nil, fmt.Errorf("ProfileService -> Login -> %w", err)
	}

	err = s.r.AddRefreshToken(ctx, hashedRefreshToken, id)
	if err != nil {
		return nil, fmt.Errorf("ProfileService ->  Login -> %w", err)
	}
	return tokenPair, nil
}

// ValidateToken parses tokenString and checks if signing method is ok and return jwt token with filled Valid field
func (s *ProfileService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("ValidateToken - > error: unexpected sign method")
		}
		return []byte(s.cfg.SecretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("ValidateToken -> %w", err)
	}
	return token, nil
}

// TokensIDCompare compares IDs from refresh and access token for being equal
func (s *ProfileService) TokensIDCompare(tokenPair *model.TokenPair) (uuid.UUID, error) {
	accessToken, err := s.ValidateToken(tokenPair.AccessToken)
	if err != nil {
		return uuid.Nil, fmt.Errorf("ProfileService -> TokensIDCompare -> %w", err)
	}
	var accessID uuid.UUID
	var uuidID uuid.UUID
	if claims, ok := accessToken.Claims.(jwt.MapClaims); ok && accessToken.Valid {
		uuidID, err = uuid.Parse(claims["id"].(string))
		if err != nil {
			return uuid.Nil, fmt.Errorf("ProfileService -> TokensIDCompare -> %w", err)
		}
		accessID = uuidID
	}
	refreshToken, err := s.ValidateToken(tokenPair.RefreshToken)
	if err != nil {
		return uuid.Nil, fmt.Errorf("ProfileService -> TokensIDCompare -> %w", err)
	}
	var refreshID uuid.UUID
	if claims, ok := refreshToken.Claims.(jwt.MapClaims); ok && refreshToken.Valid {
		exp := claims["exp"].(float64)
		uuidID, err = uuid.Parse(claims["id"].(string))
		if err != nil {
			return uuid.Nil, fmt.Errorf("ProfileService -> TokensIDCompare -> %w", err)
		}
		refreshID = uuidID
		if exp < float64(time.Now().Unix()) {
			return uuid.Nil, fmt.Errorf("ProfileService ->  TokensIDCompare -> %w", err)
		}
	}
	if accessID != refreshID {
		return uuid.Nil, fmt.Errorf("ProfileService -> error: profile ID in acess token doesn't equal profile ID in refresh token")
	}
	return accessID, nil
}

// Refresh is a method of ProfileService that refeshes access token and refresh token
func (s *ProfileService) Refresh(ctx context.Context, tokenPair *model.TokenPair) (*model.TokenPair, error) {
	if tokenPair == nil || len(tokenPair.AccessToken) == 0 || len(tokenPair.RefreshToken) == 0 {
		return nil, fmt.Errorf("ProfileService -> Refresh -> error: tokenPair is empty")
	}
	id, err := s.TokensIDCompare(tokenPair)
	if err != nil {
		return nil, fmt.Errorf("ProfileService -> Refresh ->  %w", err)
	}
	hash, err := s.r.GetRefreshTokenByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("ProfileService ->  Refresh -> %w", err)
	}
	sum := sha256.Sum256([]byte(tokenPair.RefreshToken))
	verified, err := s.CompareHashAndBytes([]byte(hash), sum[:])
	if err != nil || !verified {
		return nil, fmt.Errorf("ProfileService ->  Refresh -> CompareHashAndBytes -> error: refreshToken invalid")
	}
	tokenPair, err = s.GenerateTokenPair(id)
	if err != nil {
		return nil, fmt.Errorf("ProfileService ->  Refresh ->  %w", err)
	}
	sum = sha256.Sum256([]byte(tokenPair.RefreshToken))
	hashedRefreshToken, err := s.HashBytes(sum[:])
	if err != nil {
		return nil, fmt.Errorf("ProfileService -> Refresh ->  %w", err)
	}
	err = s.r.AddRefreshToken(context.Background(), hashedRefreshToken, id)
	if err != nil {
		return nil, fmt.Errorf("ProfileService ->  Refresh -> %w", err)
	}
	return tokenPair, nil
}
