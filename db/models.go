// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID                pgtype.UUID        `json:"id"`
	FullName          string             `json:"full_name"`
	Email             string             `json:"email"`
	PhoneNumber       string             `json:"phone_number"`
	ProfilePictureUrl string             `json:"profile_picture_url"`
	Username          string             `json:"username"`
	HashedPassword    string             `json:"hashed_password"`
	CreatedAt         pgtype.Timestamptz `json:"created_at"`
}
