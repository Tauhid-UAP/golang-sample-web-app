package models

import "time"

type User struct {
	ID string
	Email string
	FirstName string
	LastName string
	PasswordHash string
	ProfileImage *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
