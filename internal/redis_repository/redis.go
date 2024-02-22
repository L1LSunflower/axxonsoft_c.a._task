package redis_repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/L1LSunflower/axxonsoft_c.a._task/internal/entities"
	"github.com/redis/go-redis/v9"
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
	if err = c.cli.HSet(ctx, string(status), task.Id, bytes).Err(); err != nil {
		return err
	}
	return nil
}

func (c *CacheRepository) Get(ctx context.Context, status entities.Status, uuid string) (*entities.Task, error) {
	result, err := c.cli.HGet(ctx, string(status), uuid).Result()
	if err != nil && !errors.Is(err, resultIsNil) {
		return nil, err
	}
	task := new(entities.Task)
	if err = json.Unmarshal([]byte(result), &task); err != nil {
		return nil, err
	}
	return task, nil
}

func (c *CacheRepository) Delete(ctx context.Context, status entities.Status, uuid string) error {
	if err := c.cli.HDel(ctx, string(status), uuid).Err(); err != nil {
		return err
	}
	return nil
}

func (c *CacheRepository) GetTaskList(ctx context.Context, status entities.Status) ([]*entities.Task, error) {
	result, err := c.cli.Get(ctx, string(status)).Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	tasks := []*entities.Task{}
	return tasks, nil
}
