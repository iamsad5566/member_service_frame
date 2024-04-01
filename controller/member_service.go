package controller

import (
	"member_service_frame/handler"
	"member_service_frame/model"
	"member_service_frame/object"
	"member_service_frame/object/request"
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
		group.POST("/update_password", handler.TokenChecker(loginTimeRepo), handler.UpdateUserPasswordChecker, updateHandler(userRepo))
		group.POST("/check_exists_id", handler.UserChecker, checkExistsIDHandler(userRepo))
	}
}

func registerHandler(repo repo.UserRepoInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var rawRequest, _ = ctx.Get(handler.REQUEST_BODY)
		var user *object.User = rawRequest.(*object.User)

		var success, err = model.CheckExistsID(repo, user)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
				"content": err.Error(),
			})
		} else {
			if success {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "Bad request",
					"content": "Account already exists",
				})
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "",
					"content": "",
				})
			}
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
				"content": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Success",
				"content": util.GenerateToken(user),
			})
		}
	}
}

func updateHandler(userRepo repo.UserRepoInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var rawRequest, _ = ctx.Get(handler.REQUEST_BODY)
		var updateUserPassword *request.UpdateUserPassword = rawRequest.(*request.UpdateUserPassword)
		var success, err = model.UpdatePassword(userRepo, updateUserPassword)
		if err != nil && err.Error() == "password incorrect" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "",
				"content": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Success",
				"content": success,
			})
		}
	}
}

func checkExistsIDHandler(userRepo repo.UserRepoInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var rawRequest, _ = ctx.Get(handler.REQUEST_BODY)
		var user *object.User = rawRequest.(*object.User)

		var exists, err = model.CheckExistsID(userRepo, user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
				"content": err.Error(),
			})
		} else {
			if exists {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "Account exists",
					"content": false,
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Success",
					"content": true,
				})
			}
		}
	}
}
