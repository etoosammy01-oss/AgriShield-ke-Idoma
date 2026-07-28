package models

import "time"

type Farmer struct {
	ID           int
	FullName     string
	Phone        string
	PasswordHash string
	Location     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
