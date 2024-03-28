package controller

import (
	"member_service_frame/handler"
	"member_service_frame/model"
	"member_service_frame/object"
	"member_service_frame/repo"
	"member_service_frame/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MemberServiceGroup(server *gin.Engine, userRepo repo.UserRepoInterface, loginTimeRepo repo.LoginTimeInterface) {
	group := server.Group("/member")
	{
		group.POST("/register", handler.UserChecker, registerHandler(userRepo))
		group.POST("/login", handler.UserChecker, loginHandler(userRepo, loginTimeRepo))
	}
}

func registerHandler(repo repo.UserRepoInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var rawRequest, _ = ctx.Get(handler.REQUEST_BODY)
		var user *object.User = rawRequest.(*object.User)

		var success, err = model.AccountRegister(repo, user)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
				"content": err,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Success",
				"content": success,
			})
		}
	}
}

func loginHandler(userRepo repo.UserRepoInterface, loginTimeRepo repo.LoginTimeInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var rawRequest, _ = ctx.Get(handler.REQUEST_BODY)
		var user *object.User = rawRequest.(*object.User)

		_, err := model.AuthenticateUser(userRepo, loginTimeRepo, user)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"content": err,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Success",
				"content": util.GenerateToken(user),
			})
		}
	}
}
