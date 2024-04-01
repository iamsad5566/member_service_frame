package controller

import (
	"member_service_frame/model"
	"member_service_frame/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTable(server *gin.Engine, userRepo repo.UserRepoInterface) {
	server.POST("/create_table", createTableHandler(userRepo))
}

// createTableHandler is a handler function that creates a table using the provided user repository.
// It returns a gin.HandlerFunc that can be used as a route handler.
// The function takes a repo.UserRepoInterface as a parameter, which represents the user repository.
// The handler function expects a gin.Context as its parameter, which represents the HTTP request context.
// It returns a JSON response with a success message and content if the table creation is successful,
// or an error message and content if there is an internal server error.
func createTableHandler(repo repo.UserRepoInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var success, err = model.CreateTable(repo)
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
