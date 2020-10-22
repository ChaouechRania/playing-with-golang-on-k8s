package store

import (
	"time"
)

// User represents a registered user on the platform
type User struct {
	ID         string `gorm:"primary_key;"`
	Email      string `gorm:"index"`
	Password   string
	RememberMe bool
	CreatedAt  *time.Time `sql:"index"`
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}

//GetProdsRequest query object for orgs
type GetProdsRequest struct {
	UserID string
}
