package redis

import (
	"context"
	"time"
)

const keySet string = "logincheck:"

// SetLoginTime sets the login time for a given account in Redis.
// It takes a context and the account as input parameters.
// It returns a boolean indicating whether the operation was successful and an error, if any.
func (r *RedisLoginCheckRepository) SetLoginTime(ctx context.Context, account string) (bool, error) {
	err := r.client.Set(ctx, keySet+account, time.Now().Format(time.RFC3339), 0).Err()
	if err != nil {
		return false, err
	}
	return err == nil, err
}

// GetLoginTime retrieves the login time for a given account from Redis.
// It takes a context and the account as input parameters.
// It returns the login time as a time.Time value and an error, if any.
func (r *RedisLoginCheckRepository) GetLoginTime(ctx context.Context, account string) (time.Time, error) {
	result, err := r.client.Get(ctx, keySet+account).Result()
	if err != nil {
		return time.Time{}, err
	}
	return time.Parse(time.RFC3339, result)
}
