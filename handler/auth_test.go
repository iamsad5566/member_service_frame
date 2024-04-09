package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/iamsad5566/member_service_frame/handler"
	"github.com/iamsad5566/member_service_frame/object"
	"github.com/iamsad5566/member_service_frame/util"

	"github.com/alicebob/miniredis/v2"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"

	r "github.com/iamsad5566/member_service_frame/repo/redis"
)

func TestTokenCheckerWithRedis(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when running miniredis", err)
	}
	defer mr.Close()

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	loginTimeRepo := r.NewLoginCheckRepository(client)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(handler.TokenChecker(loginTimeRepo))

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	// Simulate a valid token stored in Redis
	user := &object.User{UserID: "d878392c-0b31-4f8b-a09a-1a6bdd0d1cc9", Account: "test123"}
	token := util.GenerateToken(user)

	t.Run("Valid token", func(t *testing.T) {
		client.Set(context.Background(), "logincheck:test123", time.Now().AddDate(-1, 0, 0).Format(time.RFC3339), 10*time.Minute)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)
		if req.Header.Get("message") == "login expired" {
			assert.Equal(t, http.StatusUnauthorized, w.Code)
		} else {
			assert.Equal(t, http.StatusOK, w.Code)
		}
	})

	t.Run("Login Expire", func(t *testing.T) {
		client.Set(context.Background(), "logincheck:test123", time.Now().Format(time.RFC3339), 10*time.Minute)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)
		if req.Header.Get("message") == "login expired" {
			assert.Equal(t, http.StatusUnauthorized, w.Code)
		} else {
			assert.Equal(t, http.StatusOK, w.Code)
		}
	})

	t.Run("Invalid token", func(t *testing.T) {
		// Simulate an invalid token not present in Redis
		token := "invalid_token"

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("No token", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
