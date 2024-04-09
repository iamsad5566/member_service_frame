package model_test

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"

	"github.com/iamsad5566/member_service_frame/model"
	"github.com/iamsad5566/member_service_frame/object"
	r "github.com/iamsad5566/member_service_frame/repo/redis"
	"github.com/iamsad5566/member_service_frame/util"
)

func TestCertifyToken(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when running miniredis", err)
	}
	defer mr.Close()

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	ctx := context.Background()
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIiLCJzdWIiOiJ0ZXN0MTIzIiwiZXhwIjoxNzEyMzk3Nzg5fQ.OLica5O99b85zew8poeTH7JV_46Ly-8dsIrzTD0wrOc"
	client.Set(ctx, "logincheck:test123", time.Now().Format(time.RFC3339), 10*time.Minute)
	loginTimeRepo := r.NewLoginCheckRepository(client)
	var status, _ = model.CertifyToken(loginTimeRepo, ctx, token)
	assert.Equal(t, (model.Pass | model.TokenExpired), status)

	token = "45566"
	status, _ = model.CertifyToken(loginTimeRepo, ctx, token)
	assert.Equal(t, model.WrongToken, status)

	token = util.GenerateToken(&object.User{UserID: "d878392c-0b31-4f8b-a09a-1a6bdd0d1cc9",
		Account: "test123"})
	status, _ = model.CertifyToken(loginTimeRepo, ctx, token)
	assert.Equal(t, model.Pass, status)

	client.Set(ctx, "logincheck:test123", time.Now().Add(-20*24*time.Hour).Format(time.RFC3339), 10*time.Minute)
	status, _ = model.CertifyToken(loginTimeRepo, ctx, token)
	assert.Equal(t, model.LoginExpired, status)
}

func TestCertifyOauthAccount(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when running miniredis", err)
	}
	defer mr.Close()

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	var testAccount string = "abc"

	ctx := context.Background()
	client.Set(ctx, "logincheck:"+testAccount, time.Now().Format(time.RFC3339), 10*time.Minute)
	loginTimeRepo := r.NewLoginCheckRepository(client)
	var status = model.CertifyOauthAccount(loginTimeRepo, ctx, testAccount)
	assert.Equal(t, model.Pass, status)

	client.Set(ctx, "logincheck:"+testAccount, time.Now().Add(-2*24*time.Hour).Format(time.RFC3339), 10*time.Minute)
	status = model.CertifyOauthAccount(loginTimeRepo, ctx, testAccount)
	assert.Equal(t, model.LoginExpired, status)

	status = model.CertifyOauthAccount(loginTimeRepo, ctx, "test")
	assert.Equal(t, model.WrongToken, status)
}
