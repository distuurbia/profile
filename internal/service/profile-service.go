// Package service contains the bisnes logic of app
package service

import (
	"context"
	"fmt"

	"github.com/distuurbia/profile/internal/config"
	"github.com/distuurbia/profile/internal/model"
	"github.com/google/uuid"
)

// ProfileRepository is an interface of repository.ProfileRepository and contains its methods
type ProfileRepository interface {
	CreateProfile(ctx context.Context, profile *model.Profile) error
	GetPasswordAndIDByUsername(ctx context.Context, username string) (profileID uuid.UUID, password []byte, err error)
	GetRefreshTokenByID(ctx context.Context, profileID uuid.UUID) (hashedRefresh []byte, err error)
	AddRefreshToken(ctx context.Context, refreshToken []byte, profileID uuid.UUID) error
	DeleteProfile(ctx context.Context, profileID uuid.UUID) error
}

// ProfileService contains an object of ProfileRepository and config with env variables
type ProfileService struct {
	r   ProfileRepository
	cfg *config.Config
}

// NewProfileService creates *ProfileSevice object filles it and returns
func NewProfileService(r ProfileRepository, cfg *config.Config) *ProfileService {
	return &ProfileService{r: r, cfg: cfg}
}

// CreateProfile calls lower method of ProfileRepository CreateProfile
func (s *ProfileService) CreateProfile(ctx context.Context, profile *model.Profile) (err error) {
	if profile == nil || profile.ID == uuid.Nil || profile.Username == "" || len(profile.Password) == 0 {
		return fmt.Errorf("ProfileService -> GenerateTokenPair -> error: profile is empty")
	}
	err = s.r.CreateProfile(ctx, profile)
	if err != nil {
		return fmt.Errorf("ProfileService -> %w", err)
	}
	return nil
}

// GetPasswordAndIDByUsername calls lower method of ProfileRepository GetPasswordAndIDByUsername
func (s *ProfileService) GetPasswordAndIDByUsername(ctx context.Context, username string) (profileID uuid.UUID, password []byte, err error) {
	if username == "" {
		return uuid.Nil, nil, fmt.Errorf("ProfileService -> GetPasswordAndIDByUsername -> error: username is empty")
	}
	profileID, hashedPassword, err := s.r.GetPasswordAndIDByUsername(ctx, username)
	if err != nil {
		return uuid.Nil, nil, fmt.Errorf("ProfileService ->  GetPasswordAndIDByUsername -> %w", err)
	}
	return profileID, hashedPassword, nil
}

// GetRefreshTokenByID calls lower method of ProfileRepository GetRefreshTokenByID
func (s *ProfileService) GetRefreshTokenByID(ctx context.Context, profileID uuid.UUID) (hashedRefresh []byte, err error) {
	if profileID == uuid.Nil {
		return nil, fmt.Errorf("ProfileService -> GetRefreshTokenByID -> error: failed to use uuid")
	}
	hashedRefresh, err = s.r.GetRefreshTokenByID(ctx, profileID)
	if err != nil {
		return nil, fmt.Errorf("ProfileService ->  GetRefreshTokenByID -> %w", err)
	}
	return hashedRefresh, nil
}

// AddRefreshToken calls lower method of ProfileRepository AddRefreshToken
func (s *ProfileService) AddRefreshToken(ctx context.Context, refreshToken []byte, profileID uuid.UUID) error {
	if profileID == uuid.Nil {
		return fmt.Errorf("ProfileService -> GetRefreshTokenByID -> error: failed to use uuid")
	}
	if len(refreshToken) == 0 {
		return fmt.Errorf("ProfileService -> GetRefreshTokenByID -> error: refreshToken is empty")
	}
	err := s.r.AddRefreshToken(ctx, refreshToken, profileID)
	if err != nil {
		return fmt.Errorf("ProfileService -> GetRefreshTokenByID -> %w", err)
	}
	return nil
}

// DeleteProfile calls lower method of ProfileRepository DeleteProfile
func (s *ProfileService) DeleteProfile(ctx context.Context, profileID uuid.UUID) error {
	if profileID == uuid.Nil {
		return fmt.Errorf("ProfileService -> DeleteProfile -> error: profileID is nil")
	}
	err := s.r.DeleteProfile(ctx, profileID)
	if err != nil {
		return fmt.Errorf("ProfileService -> DeleteProfile -> %w", err)
	}
	return nil
}
