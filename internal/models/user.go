package models

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}