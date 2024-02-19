package redis_repository

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/L1LSunflower/axxonsoft_c.a._task/internal/entities"
)

type CacheRepository struct {
	cli *redis.Client
}

func (c *CacheRepository) Set(ctx context.Context, uuid string, task *entities.Task, lifetime int) error {
	bytes, err := task.ToBytes()
	if err != nil {
		return err
	}
	if err = c.cli.Set(ctx, string(entities.NewStatus)+":"+uuid, bytes, time.Duration(lifetime)*time.Minute).Err(); err != nil {
		return err
	}
	return nil
}

func (c *CacheRepository) Update(ctx context.Context, task *entities.Task, status entities.Status, uuid string) error {
	bytes, err := task.ToBytes()
	if err != nil {
		return err
	}
	if err = c.cli.GetSet(ctx, string(status)+":"+uuid, bytes).Err(); err != nil {
		return err
	}
	return nil
}

func (c *CacheRepository) Get(ctx context.Context, status entities.Status, uuid string) (*entities.Task, error) {
	result, err := c.cli.Get(ctx, string(status)+":"+uuid).Result()
	if err != nil {
		return nil, err
	}
	task := new(entities.Task)
	if err = json.Unmarshal([]byte(result), &task); err != nil && errors.Is(err) {
		return nil, err
	}
	return nil, nil
}
