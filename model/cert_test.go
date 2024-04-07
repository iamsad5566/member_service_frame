package model_test

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"

	"member_service_frame/model"
	r "member_service_frame/repo/redis"
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
}
