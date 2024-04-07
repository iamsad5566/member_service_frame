package redis

import (
	"context"
	"time"

	"github.com/iamsad5566/member_service_frame/repo"

	"github.com/redis/go-redis/v9"
)

// RedisLoginCheckRepository is a repository implementation that uses Redis for login time checks.
type RedisLoginCheckRepository struct {
	client RedisClientInterface
}

// RedisClientInterface defines the interface for interacting with Redis.
type RedisClientInterface interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

// NewLoginCheckRepository creates a new instance of RedisLoginCheckRepository.
func NewLoginCheckRepository(client RedisClientInterface) *RedisLoginCheckRepository {
	return &RedisLoginCheckRepository{client: client}
}

// RedisLoginCheckRepository implements the LoginTimeInterface from the repo package.
var _ repo.LoginTimeInterface = (*RedisLoginCheckRepository)(nil)
