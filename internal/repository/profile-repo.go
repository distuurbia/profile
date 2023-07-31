// Package repository contains methods working with db
package repository

import (
	"context"
	"fmt"

	"github.com/distuurbia/profile/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ProfileRepository contains pgxpool
type ProfileRepository struct {
	pool *pgxpool.Pool
}

// NewProfileRepository creates an object of *ProfileRepository
func NewProfileRepository(pool *pgxpool.Pool) *ProfileRepository{
	return &ProfileRepository{pool: pool}
}

// SugnUp creates the row in db with fields of model.Profile
func (r *ProfileRepository) SignUp(ctx context.Context, profile *model.Profile) error {
	if profile.ID == uuid.Nil{
		return fmt.Errorf("ProfileRepository -> SignUp -> error: failed to use uuid")
	}
	var count int
	err := r.pool.QueryRow(ctx, "SELECT COUNT(*) FROM profiles WHERE username = $1", profile.Username).Scan(&count)
	if err != nil {
		return fmt.Errorf("ProfileRepository -> SignUp -> QueryRow -> %w", err)
	} 
	if count > 0 {
		return fmt.Errorf("ProfileRepository -> SignUp -> QueryRow -> error: profile with such username already exists")
	}

	_, err = r.pool.Exec(ctx, "INSERT into profiles (id, username, password, refreshToken, country, age) VALUES($1, $2, $3, $4, $5, $6)",
		profile.ID, profile.Username, profile.Password, profile.RefreshToken, profile.Country, profile.Age)
	if err != nil {
		return fmt.Errorf("ProfileRepository -> SignUp -> Exec -> %w", err)
	}

	return nil
}

// GetPasswordAndIDByUsername returns hash of the password and id from profiles table
func (r *ProfileRepository) GetPasswordAndIDByUsername(ctx context.Context, username string) (id uuid.UUID, password []byte, err error) {
	err = r.pool.QueryRow(ctx, "SELECT id, password FROM profiles WHERE username = $1", username).Scan(&id, &password)
	if err != nil {
		return uuid.Nil, nil, fmt.Errorf("ProfileRepository -> GetPasswordAndIDByUserName: %w", err)
	}

	return id, password, nil
}

// GetRefreshTokenByID returnes refreshToken from profiles table from excact row by id
func (r *ProfileRepository) GetRefreshTokenByID(ctx context.Context, id uuid.UUID) (hashedRefresh string, err error) {
	err = r.pool.QueryRow(ctx, "SELECT refreshToken FROM profiles WHERE id = $1", id).Scan(&hashedRefresh)
	if err != nil {
		return "", fmt.Errorf("ProfileRepository -> GetRefreshTokenByName: %w", err)
	}

	return hashedRefresh, nil
}

// AddRefreshToken adds refreshToken to profiles table in excact row by id
func (r *ProfileRepository) AddRefreshToken(ctx context.Context, refreshToken string, id uuid.UUID) error {
	_, err := r.pool.Exec(ctx, "UPDATE prfoles SET refreshtoken = $1 WHERE id = $2", refreshToken, id)
	if err != nil {
		return fmt.Errorf("ProfileRepository -> AddRefreshToken: %w", err)
	}

	return nil
}

// DeleteProfile deletes exact row from profiles table
func (r *ProfileRepository) DeleteProfile(ctx context.Context, id uuid.UUID) error {
	res, err := r.pool.Exec(ctx, "DELETE FROM profiles WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("ProfileRepository -> DeleteProfile -> error: %w", err)
	}
	if res.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}