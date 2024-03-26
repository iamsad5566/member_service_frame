package controller

import (
	"member_service_frame/handler"
	"member_service_frame/object"
	"member_service_frame/repo"

	"github.com/gin-gonic/gin"
)

func MemberServiceGroup(server *gin.Engine, memberRepo repo.UserRepoInterface) {
	group := server.Group("/member")
	{
		group.POST("/register")
	}
}

func registerHandler(repo repo.UserRepoInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var rawRequest, _ = ctx.Get(handler.REQUEST_BODY)
		var user *object.User = rawRequest.(*object.User)

	}
}
