package controller

import (
	"member_service_frame/handler"
	"member_service_frame/model"
	"member_service_frame/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTable(server *gin.Engine, userRepo repo.UserRepoInterface, loginTimeRepo repo.LoginTimeInterface) {
	server.POST("/create_table", handler.TokenChecker(loginTimeRepo), createTableHandler(userRepo))
}

// createTableHandler creates a table in the database.
// @Summary Create table
// @Description Creates a new table in the database.
// @Tags table
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "message: Success, content: true if the table was successfully created"
// @Failure 500 {object} map[string]interface{} "message: Internal server error, content: error description"
// @Router /create_table [post]
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
