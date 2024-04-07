package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iamsad5566/member_service_frame/handler"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestUserChecker is a unit test function for the UserChecker handler.
// It tests the behavior of the handler when different inputs are provided.
func TestUserChecker(t *testing.T) {
	// Initialize Gin Engine
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/test", handler.UserChecker)

	tests := []struct {
		description  string
		body         string
		expectedCode int
	}{
		{
			description:  "Valid input",
			body:         `{"account": "user1", "password": "pass1"}`,
			expectedCode: http.StatusOK,
		},
		{
			description:  "Invalid input - empty account",
			body:         `{"account": "", "password": "pass1"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			description:  "Invalid input - bad JSON",
			body:         `{"account": "user1", "password": "pass1"`,
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/test", bytes.NewBufferString(test.body))
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, test.expectedCode, resp.Code)
		})
	}
}

func TestUpdateUserPasswordChecker(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/test", handler.UpdateUserPasswordChecker)

	tests := []struct {
		description  string
		body         string
		expectedCode int
	}{
		{
			description:  "Valid input",
			body:         `{"account": "user1", "password": "pass1", "new_password": "pass2"}`,
			expectedCode: http.StatusOK,
		},
		{
			description:  "Invalid input - empty account",
			body:         `{"account": "", "password": "pass1", "new_password": "pass2"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			description:  "Invalid input - empty password",
			body:         `{"account": "user1", "password": "", "new_password": "pass2"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			description:  "Invalid input - empty new password",
			body:         `{"account": "user1", "password": "pass1", "new_password": ""}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			description:  "Invalid input - new password same as old password",
			body:         `{"account": "user1", "password": "pass1", "new_password": "pass1"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			description:  "Invalid input - bad JSON",
			body:         `{"account": "user1", "password": "pass1", "new_password": "pass2"`,
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/test", bytes.NewBufferString(test.body))
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, test.expectedCode, resp.Code)
		})
	}
}
