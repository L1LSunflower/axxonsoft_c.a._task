package services

import (
	"context"
	"errors"

	"github.com/L1LSunflower/axxonsoft_c.a._task/internal/entities"

	redisRepo "github.com/L1LSunflower/axxonsoft_c.a._task/internal/redis_repository"
)

var taskNotExist = errors.New("task by id does not exist")

func RegisterTask() {

}

func TaskStatus(ctx context.Context, cache redisRepo.CacheInterface, uuid string) (*entities.Task, error) {
	for _, status := range entities.StatusesOrder {
		task, err := cache.Get(ctx, status, uuid)
		if err != nil {
			return nil, err
		}
		if task.Id != "" {
			return task, nil
		}
	}
	return nil, taskNotExist
}
