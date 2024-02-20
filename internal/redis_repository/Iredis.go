package redis_repository

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/L1LSunflower/axxonsoft_c.a._task/internal/entities"
)

type CacheInterface interface {
	Set(ctx context.Context, status entities.Status, task *entities.Task, lifetime int) error
	Get(ctx context.Context, status entities.Status, uuid string) (*entities.Task, error)
}

func NewRepository(cli *redis.Client) CacheInterface {
	return &CacheRepository{cli: cli}
}
