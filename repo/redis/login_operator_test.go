package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"

	r "github.com/iamsad5566/member_service_frame/repo/redis"
)

// TestRedisLoginCheckRepository is a unit test function that tests the RedisLoginCheckRepository.
// It sets the login time for an account, retrieves the login time, and checks if the value is correctly set in miniredis.
func TestRedisLoginCheckRepository(t *testing.T) {
	// Create a new miniredis instance.
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when running miniredis", err)
	}
	defer mr.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	repo := r.NewLoginCheckRepository(rdb)

	// Set the login time for an account.
	ctx := context.Background()
	account := "user1"
	success, err := repo.SetLoginTime(ctx, account)
	// Check if the login time is set correctly.
	assert.True(t, success)
	assert.NoError(t, err)

	// Retrieve the login time for an account.
	retrievedTime, err := repo.GetLoginTime(ctx, account)
	// Check if the login time is retrieved correctly.
	assert.NoError(t, err)
	assert.WithinDuration(t, time.Now(), retrievedTime, time.Second)

	// Check if the value is correctly set in miniredis.
	storedTime, err := mr.Get("logincheck:" + account)
	assert.NoError(t, err)
	parsedTime, _ := time.Parse(time.RFC3339, storedTime)
	assert.WithinDuration(t, retrievedTime, parsedTime, time.Second)
}
