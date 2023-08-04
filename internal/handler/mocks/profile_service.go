// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/distuurbia/profile/internal/model"

	uuid "github.com/google/uuid"
)

// ProfileService is an autogenerated mock type for the ProfileService type
type ProfileService struct {
	mock.Mock
}

// AddRefreshToken provides a mock function with given fields: ctx, refreshToken, profileID
func (_m *ProfileService) AddRefreshToken(ctx context.Context, refreshToken []byte, profileID uuid.UUID) error {
	ret := _m.Called(ctx, refreshToken, profileID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, uuid.UUID) error); ok {
		r0 = rf(ctx, refreshToken, profileID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateProfile provides a mock function with given fields: ctx, profile
func (_m *ProfileService) CreateProfile(ctx context.Context, profile *model.Profile) error {
	ret := _m.Called(ctx, profile)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Profile) error); ok {
		r0 = rf(ctx, profile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteProfile provides a mock function with given fields: ctx, profileID
func (_m *ProfileService) DeleteProfile(ctx context.Context, profileID uuid.UUID) error {
	ret := _m.Called(ctx, profileID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, profileID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPasswordAndIDByUsername provides a mock function with given fields: ctx, username
func (_m *ProfileService) GetPasswordAndIDByUsername(ctx context.Context, username string) (uuid.UUID, []byte, error) {
	ret := _m.Called(ctx, username)

	var r0 uuid.UUID
	var r1 []byte
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (uuid.UUID, []byte, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) uuid.UUID); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) []byte); ok {
		r1 = rf(ctx, username)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, username)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetRefreshTokenByID provides a mock function with given fields: ctx, profileID
func (_m *ProfileService) GetRefreshTokenByID(ctx context.Context, profileID uuid.UUID) ([]byte, error) {
	ret := _m.Called(ctx, profileID)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]byte, error)); ok {
		return rf(ctx, profileID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []byte); ok {
		r0 = rf(ctx, profileID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, profileID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProfileService creates a new instance of ProfileService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProfileService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProfileService {
	mock := &ProfileService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
