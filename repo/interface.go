package repo

import (
	"context"
	"time"

	"github.com/iamsad5566/member_service_frame/object"
)

// UserRepoInterface defines the interface for user repository operations.
type UserRepoInterface interface {
	// Register registers a new user.
	// It takes a pointer to a User object as input and returns a boolean indicating success and an error, if any.
	Register(usr *object.User) (bool, error)

	// GetPassword retrieves the password for a user.
	// It takes a pointer to a User object as input and returns the password as a string and an error, if any.
	GetPassword(usr *object.User) (string, error)

	// UpdateLastTimeLogin updates the last time a user logged in.
	// It takes a pointer to a User object as input and returns a boolean indicating success and an error, if any.
	UpdateLastTimeLogin(usr *object.User) (bool, error)

	// UpdatePassword updates the password for a user.
	// It takes a pointer to a User object as input and returns a boolean indicating success and an error, if any.
	UpdatePassword(usr *object.User) (bool, error)

	// DeleteAccount deletes a user account.
	// It takes a pointer to a User object as input and returns a boolean indicating success and an error, if any.
	DeleteAccount(usr *object.User) (bool, error)

	// CheckExistsID checks if a user ID exists.
	// It takes a pointer to a User object as input and returns a boolean indicating whether the ID exists and an error, if any.
	CheckExistsID(usr *object.User) (bool, error)

	// CreateTable creates the user table in the database.
	CreateTable() (bool, error)
}

// LoginTimeInterface represents an interface for managing login time information.
type LoginTimeInterface interface {
	// SetLoginTime sets the login time for the specified account.
	// It returns a boolean indicating whether the login time was successfully set,
	// and an error if any occurred.
	SetLoginTime(ctx context.Context, account string) (bool, error)

	// GetLoginTime retrieves the login time for the specified account.
	// It returns the login time as a time.Time value,
	// and an error if any occurred.
	GetLoginTime(ctx context.Context, account string) (time.Time, error)
}
