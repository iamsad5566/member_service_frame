package handler

import (
	"member_service_frame/object"
	"net/http"

	"github.com/gin-gonic/gin"
)

const REQUEST_BODY = "requestBody"

// UserChecker is a handler function that checks the validity of the user information provided in the request body.
// It binds the JSON data to the 'user' variable and checks if the 'Account' field is empty.
// If the input is invalid, it aborts the request with a JSON response containing an error message.
// If the input is valid, it sets the 'user' object in the context for further processing.
func UserChecker(ctx *gin.Context) {
	var user *object.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil || user.Account == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	}
	ctx.Set(REQUEST_BODY, user)
}
