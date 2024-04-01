package model

import (
	"context"
	"log"
	"member_service_frame/config"
	"member_service_frame/object"
	"member_service_frame/repo"
	"member_service_frame/util"
	"time"

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
	if err != nil && err.Error() == jwt.ErrTokenExpired.Error() {
		return TokenExpired, util.GenerateToken(object.NewUser(id, account))
	} else if err != nil {
		return WrongToken, ""
	}

	if !loginStillValid(loginTimeRepo, ctx, account) {
		return LoginExpired, ""
	}
	return Pass, ""
}

func loginStillValid(loginTimeRepo repo.LoginTimeInterface, ctx context.Context, account string) bool {
	var lastTimeLogin, err = loginTimeRepo.GetLoginTime(ctx, account)
	if err != nil {
		log.Println(err)
		return false
	}
	var now time.Time = time.Now()
	var duration time.Duration = now.Sub(lastTimeLogin)
	return duration.Hours() < float64(config.Setting.GetValidLogin())*24.0
}
