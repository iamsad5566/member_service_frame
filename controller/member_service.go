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

// registerHandler registers a new user.
// @Summary Register user
// @Description Registers a new user if the account does not already exist.
// @Tags user
// @Accept json
// @Produce json
// @Param user body object.User true "User to register"
// @Success 200 {object} map[string]interface{} "message: User registered successfully"
// @Failure 400 {object} map[string]interface{} "message: Bad request, content: Account already exists"
// @Failure 500 {object} map[string]interface{} "message: Internal server error, content: Error description"
// @Router /member/register [post]
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

// loginHandler logs in a user.
// @Summary User login
// @Description Logs in a user and returns a token if the authentication is successful.
// @Tags user
// @Accept json
// @Produce json
// @Param user body object.User true "User credentials"
// @Success 200 {object} map[string]interface{} "message: Success, content: Token"
// @Failure 401 {object} map[string]interface{} "message: Unauthorized, content: Error description"
// @Router /member/login [post]
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

// updateHandler updates a user's password.
// @Summary Update user password
// @Description Updates the password for a user.
// @Tags user
// @Accept json
// @Produce json
// @Param updateUserPassword body request.UpdateUserPassword true "User ID and new password"
// @Success 200 {object} map[string]interface{} "message: Success, content: true if the password was successfully updated"
// @Failure 401 {object} map[string]interface{} "message: Unauthorized, content: Password incorrect"
// @Router /member/update_password [post]
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

// checkExistsIDHandler checks if a user ID already exists.
// @Summary Check user ID
// @Description Checks if a user ID already exists in the database.
// @Tags user
// @Accept json
// @Produce json
// @Param user body object.User true "User ID to check"
// @Success 200 {object} map[string]interface{} "message: Success, content: true if the user ID does not exist"
// @Failure 401 {object} map[string]interface{} "message: Account exists, content: false if the user ID already exists"
// @Failure 500 {object} map[string]interface{} "message: Internal server error, content: Error description"
// @Router /member/check_exists_id [post]
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
