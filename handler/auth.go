package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/iamsad5566/member_service_frame/model"
	"github.com/iamsad5566/member_service_frame/repo"

	"github.com/gin-gonic/gin"
)

func TokenChecker(loginTimeRepo repo.LoginTimeInterface) gin.HandlerFunc {
	var context context.Context = context.Background()
	return func(ctx *gin.Context) {
		var token string = ctx.GetHeader("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		var res, refreshToken = model.CertifyToken(loginTimeRepo, context, strings.Replace(token, "Bearer ", "", -1))
		if res == model.TokenExpired {
			ctx.Header("Authorization", "Bearer "+refreshToken)
		} else if res == model.WrongToken {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "wrong token"})
			return
		} else if res == model.LoginExpired {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "login expired"})
			ctx.Header("message", "login expired")
			return
		} else {
			ctx.Next()
		}
	}
}
