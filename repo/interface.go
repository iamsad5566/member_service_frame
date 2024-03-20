package repo

import (
	"context"
	"member_service_frame/object"
	"time"
)

type UserRepoInterface interface {
	Register(usr *object.User) bool
	GetPassword(usr *object.User) (string, error)
	UpdateLastLogin(usr *object.User) error
	UpdatePassword(usr *object.User) bool
	DeleteAccount(usr *object.User) bool
	CheckExistsID(usr *object.User) bool
}

type LoginTimeInterface interface {
	SetLoginTime(ctx context.Context, account string) bool
	GetLoginTime(ctx context.Context, account string) time.Time
}
