package redis_repository

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/L1LSunflower/axxonsoft_c.a._task/internal/entities"
)

var (
	resultIsNil = errors.New("redis: nil")
)

type CacheRepository struct {
	cli *redis.Client
}

func (c *CacheRepository) Set(ctx context.Context, status entities.Status, task *entities.Task, lifetime int) error {
	bytes, err := task.ToBytes()
	if err != nil {
		return err
	}
	if err = c.cli.Set(ctx, string(status)+":"+task.Id, bytes, time.Duration(lifetime)*time.Minute).Err(); err != nil {
		return err
	}
	return nil
}

func (c *CacheRepository) Get(ctx context.Context, status entities.Status, uuid string) (*entities.Task, error) {
	result, err := c.cli.Get(ctx, string(status)+":"+uuid).Result()
	if err != nil && !errors.Is(err, resultIsNil) {
		return nil, err
	}
	task := new(entities.Task)
	if err = json.Unmarshal([]byte(result), &task); err != nil {
		return nil, err
	}
	return task, nil
}
