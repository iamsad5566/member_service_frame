package handler

import (
	"context"
	"member_service_frame/model"
	"member_service_frame/repo"
	"net/http"
	"strings"

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

		var res, _ = model.CertifyToken(loginTimeRepo, context, strings.Replace(token, "Bearer ", "", -1))
		if res != model.Pass {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
	}
}
