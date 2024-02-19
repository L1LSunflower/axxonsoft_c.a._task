package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/L1LSunflower/axxonsoft_c.a._task/config"
	"github.com/L1LSunflower/axxonsoft_c.a._task/internal/entities"
	
	redisRepo "github.com/L1LSunflower/axxonsoft_c.a._task/internal/redis_repository"
	"github.com/L1LSunflower/axxonsoft_c.a._task/internal/services"

	"github.com/L1LSunflower/axxonsoft_c.a._task/pkg/http/middlewares"
	"github.com/L1LSunflower/axxonsoft_c.a._task/pkg/redis"
)

const (
	contentTypeKey = "content-type"
	contentTypeVal = "application/json"
)

func ErrorResponse(resp http.ResponseWriter, errorMessage string, statusCode int) {
	b, err := entities.NewErrorResponse(errorMessage).ToBytes()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp.Write(b)
	resp.Header().Set(contentTypeKey, contentTypeVal)
	resp.WriteHeader(statusCode)
}

func SuccessResponse(resp http.ResponseWriter, data any) {
	b, err := json.Marshal(data)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp.Write(b)
	resp.Header().Set(contentTypeKey, contentTypeVal)
	resp.WriteHeader(http.StatusOK)
}

func RegisterTask(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		ErrorResponse(resp, "that method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if cTypeVal := req.Header.Get(contentTypeKey); cTypeVal != contentTypeVal {
		ErrorResponse(resp, "wrong content type", http.StatusBadRequest)
		return
	}
}

func Task(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		ErrorResponse(resp, "that method not allowed", http.StatusMethodNotAllowed)
		return
	}
	requestId := req.Header.Get(middlewares.Id)
	if requestId == "" {
		ErrorResponse(resp, "task is missing in parameter", http.StatusBadRequest)
		return
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	cfg := config.GetConfig()
	redisInstance, err := redis.Instance(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Db, cfg.Redis.Tls)
	if err != nil {
		ErrorResponse(resp, "failed connection to redis", http.StatusFailedDependency)
	}
	cache := redisRepo.NewRepository(redisInstance.Client)
	task, err := services.TaskStatus(ctx, cache, requestId)
	if err != nil {
		ErrorResponse(resp, err.Error(), http.StatusInternalServerError)
	}
	SuccessResponse(resp, task)
}
