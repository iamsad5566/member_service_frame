package model

import (
	"context"
	"time"

	"github.com/iamsad5566/member_service_frame/config"
	"github.com/iamsad5566/member_service_frame/object"
	"github.com/iamsad5566/member_service_frame/repo"
	"github.com/iamsad5566/member_service_frame/util"

	"github.com/golang-jwt/jwt/v5"
)

const (
	Pass int = iota
	LoginExpired
	TokenExpired
	WrongToken
)

func CertifyToken(loginTimeRepo repo.LoginTimeInterface, ctx context.Context, token string) (int, string) {
	id, account, err := util.CertificateToken(token)
	valid, getLoginTimeErr := loginStillValid(loginTimeRepo, ctx, account)
	if !valid && getLoginTimeErr == nil {
		return LoginExpired, ""
	}

	if err != nil && err.Error() == jwt.ErrTokenExpired.Error() {
		return TokenExpired, util.GenerateToken(object.NewUser(id, account))
	} else if err != nil {
		return WrongToken, ""
	}

	return Pass, ""
}

func loginStillValid(loginTimeRepo repo.LoginTimeInterface, ctx context.Context, account string) (bool, error) {
	var lastTimeLogin, err = loginTimeRepo.GetLoginTime(ctx, account)
	if err != nil {
		return false, err
	}
	var now time.Time = time.Now()
	var duration time.Duration = now.Sub(lastTimeLogin)
	return duration.Hours() < float64(config.Setting.GetValidLogin())*24.0, nil
}

func CertifyOauthAccount(loginTimeRepo repo.LoginTimeInterface, ctx context.Context, account string) int {
	var lastTimeLogin, err = loginTimeRepo.GetLoginTime(ctx, account)
	if err != nil {
		return WrongToken
	}
	var now time.Time = time.Now()
	var duration time.Duration = now.Sub(lastTimeLogin)
	if duration.Hours() < 24.0 {
		return Pass
	} else {
		return LoginExpired
	}
}
