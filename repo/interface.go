package repo

import (
	"context"
	"member_service_frame/object"
	"time"
)

type UserRepoInterface interface {
	Register(usr *object.User) (bool, error)
	GetPassword(usr *object.User) (string, error)
	UpdateLastTimeLogin(usr *object.User) (bool, error)
	UpdatePassword(usr *object.User) (bool, error)
	DeleteAccount(usr *object.User) (bool, error)
	CheckExistsID(usr *object.User) (bool, error)
}

type LoginTimeInterface interface {
	SetLoginTime(ctx context.Context, account string) bool
	GetLoginTime(ctx context.Context, account string) time.Time
}
