package config_test

import (
	"member_service_frame/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCorsMiddleware(t *testing.T) {
	router := gin.Default()
	router.Use(config.Cors())

	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Test")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "Content-Type, AccessToken,X-CSRF-Token, Authorization, Token", w.Header().Get("Access-Control-Allow-Headers"))
	assert.Equal(t, "POST, GET, PUT, DELETE, PATCH", w.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type", w.Header().Get("Access-Control-Expose-Headers"))
	assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Test", w.Body.String())
}

func TestCorsMiddlewareOptionsRequest(t *testing.T) {
	router := gin.Default()
	router.Use(config.Cors())

	router.OPTIONS("/test", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "Content-Type, AccessToken,X-CSRF-Token, Authorization, Token", w.Header().Get("Access-Control-Allow-Headers"))
	assert.Equal(t, "POST, GET, PUT, DELETE, PATCH", w.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type", w.Header().Get("Access-Control-Expose-Headers"))
	assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
	assert.Equal(t, http.StatusNoContent, w.Code)
}
