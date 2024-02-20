package services

import (
	"context"
	"time"

	"github.com/L1LSunflower/axxonsoft_c.a._task/internal/cerrors"
	"github.com/L1LSunflower/axxonsoft_c.a._task/internal/entities"

	redisRepo "github.com/L1LSunflower/axxonsoft_c.a._task/internal/redis_repository"

	"github.com/L1LSunflower/axxonsoft_c.a._task/pkg/logger"
)

func RegisterTask(ctx context.Context, cache redisRepo.CacheInterface, task *entities.Task, lifetime int) error {
	if err := cache.Set(ctx, entities.NewStatus, task, lifetime); err != nil {
		logger.Instance().Error(&logger.Message{
			Message:     cerrors.ErrCreateTask.Error(),
			FullMessage: cerrors.ErrCreateTask.Error() + " with error: " + err.Error(),
			Datetime:    time.Now().Unix(),
			RequestId:   task.Id,
		})
		return cerrors.ErrCreateTask
	}
	return nil
}

func TaskStatus(ctx context.Context, cache redisRepo.CacheInterface, uuid string) (*entities.Task, error) {
	for _, status := range entities.StatusesOrder {
		task, err := cache.Get(ctx, status, uuid)
		if err != nil {
			logger.Instance().Error(&logger.Message{
				Message:     cerrors.ErrGetTask.Error(),
				FullMessage: cerrors.ErrGetTask.Error() + " with error: " + err.Error(),
				Datetime:    time.Now().Unix(),
				RequestId:   uuid,
			})
			return nil, err
		}
		if task.Id != "" {
			return task, nil
		}
	}
	return nil, cerrors.TaskNotExist
}
