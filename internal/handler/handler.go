// Package handler contains methods that handle requests and send them to service part
package handler

import (
	"context"
	"fmt"

	"github.com/distuurbia/profile/internal/model"
	protocol "github.com/distuurbia/profile/protocol/profile"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// ProfileService is an interface that contains methods of service part
type ProfileService interface {
	CreateProfile(ctx context.Context, profile *model.Profile) error
	GetPasswordAndIDByUsername(ctx context.Context, username string) (profileID uuid.UUID, password []byte, err error)
	GetRefreshTokenByID(ctx context.Context, profileID uuid.UUID) (hashedRefresh []byte, err error)
	AddRefreshToken(ctx context.Context, refreshToken []byte, profileID uuid.UUID) error
	DeleteProfile(ctx context.Context, profileID uuid.UUID) error
}

// ProfileHandler is a structure of handler that contains an object implemented ProfileService interface and validator
type ProfileHandler struct {
	s        ProfileService
	validate *validator.Validate
	protocol.UnimplementedProfileServiceServer
}

// NewProfileHandler creates an onject of *ProfileHandler fulfilled with provided fields
func NewProfileHandler(s ProfileService, validate *validator.Validate) *ProfileHandler {
	return &ProfileHandler{s: s, validate: validate}
}

// CreateProfile validates fields of the request and send them to the service
func (h *ProfileHandler) CreateProfile(ctx context.Context, req *protocol.CreateProfileRequest) (*protocol.CreateProfileResponse, error) {
	var profile = model.Profile{
		ID:       uuid.New(),
		Age:      req.Profile.Age,
		Country:  req.Profile.Country,
		Username: req.Profile.Username,
		Password: req.Profile.Password,
	}
	err := h.validate.StructCtx(ctx, profile)
	if err != nil {
		logrus.Errorf("ProfileHandler -> CreateProfile -> %v", err)
		return &protocol.CreateProfileResponse{}, err
	}
	err = h.s.CreateProfile(ctx, &profile)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username":      profile.Username,
			"Password":      profile.Password,
			"Refresh Token": profile.RefreshToken,
			"Age":           profile.Age,
			"Country":       profile.Country,
			"ID":            profile.ID,
		}).Errorf("ProfileHandler -> CreateProfile -> %v", err)
		return &protocol.CreateProfileResponse{}, err
	}
	return &protocol.CreateProfileResponse{}, nil
}

// GetPasswordAndIDByUsername validates username from request and sends it lower to the service
func (h *ProfileHandler) GetPasswordAndIDByUsername(ctx context.Context, req *protocol.GetPasswordAndIDByUsernameRequest) (
	*protocol.GetPasswordAndIDByUsernameResponse, error) {
	err := h.validate.VarCtx(ctx, req.Username, "required,min=4,max=20")
	if err != nil {
		logrus.Errorf("ProfileHandler -> GetPasswordAndIDByUsername -> %v", err)
		return &protocol.GetPasswordAndIDByUsernameResponse{}, err
	}

	id, password, err := h.s.GetPasswordAndIDByUsername(ctx, req.Username)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": req.Username,
		}).Errorf("ProfileHandler -> GetPasswordAndIDByUsername -> %v", err)
		return &protocol.GetPasswordAndIDByUsernameResponse{}, nil
	}
	return &protocol.GetPasswordAndIDByUsernameResponse{Id: id.String(), Password: password}, nil
}

// ValidationID validate given in and parses it to uuid.UUID type
func (h *ProfileHandler) ValidationID(ctx context.Context, id string) (uuid.UUID, error) {
	err := h.validate.VarCtx(ctx, id, "required,uuid")
	if err != nil {
		logrus.Errorf("ValidationID -> %v", err)
		return uuid.Nil, err
	}

	profileID, err := uuid.Parse(id)
	if err != nil {
		logrus.Errorf("ValidationID -> %v", err)
		return uuid.Nil, err
	}

	if profileID == uuid.Nil {
		logrus.Errorf("ValidationID -> error: failed to use uuid")
		return uuid.Nil, fmt.Errorf("ValidationID -> error: failed to use uuid")
	}
	return profileID, nil
}

// GetRefreshTokenByID validates id from request and sends it lower to the service
func (h *ProfileHandler) GetRefreshTokenByID(ctx context.Context, req *protocol.GetRefreshTokenByIDRequest) (
	*protocol.GetRefreshTokenByIDResponse, error) {
	profileID, err := h.ValidationID(ctx, req.Id)
	if err != nil {
		logrus.Errorf("ProfileHandler -> GetRefreshTokenByID %v", err)
		return &protocol.GetRefreshTokenByIDResponse{}, err
	}
	hashedRefresh, err := h.s.GetRefreshTokenByID(ctx, profileID)
	if err != nil {
		logrus.Errorf("ProfileHandler -> GetRefreshTokenByID %v", err)
		return &protocol.GetRefreshTokenByIDResponse{}, err
	}
	return &protocol.GetRefreshTokenByIDResponse{HashedRefresh: hashedRefresh}, nil
}

// AddRefreshToken validates id from request and sends it lower to the service
func (h *ProfileHandler) AddRefreshToken(ctx context.Context, req *protocol.AddRefreshTokenRequest) (
	*protocol.AddRefreshTokenResponse, error) {
	profileID, err := h.ValidationID(ctx, req.Id)
	if err != nil {
		logrus.Errorf("ProfileHandler -> AddRefreshToken %v", err)
		return &protocol.AddRefreshTokenResponse{}, err
	}
	err = h.s.AddRefreshToken(ctx, req.HashedRefresh, profileID)
	if err != nil {
		logrus.Errorf("ProfileHandler -> AddRefreshToken %v", err)
		return &protocol.AddRefreshTokenResponse{}, err
	}
	return &protocol.AddRefreshTokenResponse{}, nil
}

// DeleteProfile deletes exact profile by id from db using lower levels of microservice
func (h *ProfileHandler) DeleteProfile(ctx context.Context, req *protocol.DeleteProfileRequest) (*protocol.DeleteProfileResponse, error) {
	profileID, err := h.ValidationID(ctx, req.Id)
	if err != nil {
		logrus.Errorf("ProfileHandler -> DeleteProfile %v", err)
		return &protocol.DeleteProfileResponse{}, err
	}
	err = h.s.DeleteProfile(ctx, profileID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"id": req.Id,
		}).Errorf("ProfileHandler -> DeleteProfile -> %v", err)
		return &protocol.DeleteProfileResponse{}, err
	}
	return &protocol.DeleteProfileResponse{}, nil
}
