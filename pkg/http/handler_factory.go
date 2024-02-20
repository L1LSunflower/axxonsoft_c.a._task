package http

import (
	"context"
	"encoding/json"
	"github.com/L1LSunflower/axxonsoft_c.a._task/pkg/logger"
	"io"
	"net/http"
	"time"

	"github.com/L1LSunflower/axxonsoft_c.a._task/config"
	"github.com/L1LSunflower/axxonsoft_c.a._task/internal/entities"

	redisRepo "github.com/L1LSunflower/axxonsoft_c.a._task/internal/redis_repository"
	"github.com/L1LSunflower/axxonsoft_c.a._task/internal/services"

	"github.com/L1LSunflower/axxonsoft_c.a._task/pkg/redis"
)

const (
	contentTypeKey = "Content-Type"
	contentTypeVal = "application/json"
)

const (
	processTime = 15 * time.Second
)

func ErrorResponse(resp http.ResponseWriter, errorMessage string, statusCode int) {
	b, err := entities.NewErrorResponse(errorMessage).ToBytes()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp.Header().Add(contentTypeKey, contentTypeVal)
	resp.Write(b)
	resp.WriteHeader(statusCode)
}

func SuccessResponse(resp http.ResponseWriter, data any, requestId string) {
	b, err := json.Marshal(data)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	AfterAttempt(logger.Instance(), requestId, b)
	resp.Header().Add(contentTypeKey, contentTypeVal)
	resp.Write(b)
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
	requestId := req.Header.Get(Id)
	if requestId == "" {
		ErrorResponse(resp, "failed to generate id", http.StatusInternalServerError)
		return
	}
	newTask := new(entities.RegisterTask)
	reqBody, err := io.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		ErrorResponse(resp, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(reqBody, &newTask); err != nil {
		ErrorResponse(resp, err.Error(), http.StatusBadRequest)
		return
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), processTime)
	defer cancelFunc()
	cfg := config.GetConfig()
	redisInstance, err := redis.Instance(cfg.Redis.Host, cfg.Redis.Password, cfg.Redis.Db, cfg.Redis.Tls)
	if err != nil {
		ErrorResponse(resp, "failed connection to redis", http.StatusFailedDependency)
		return
	}
	cache := redisRepo.NewRepository(redisInstance.Client)
	if err = services.RegisterTask(ctx, cache, newTask.ToTask(requestId), cfg.TaskLifetime); err != nil {
		ErrorResponse(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	SuccessResponse(resp, struct {
		Id string `json:"id"`
	}{Id: requestId}, requestId)
}

func Task(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		ErrorResponse(resp, "that method not allowed", http.StatusMethodNotAllowed)
		return
	}
	requestId := req.Header.Get(Id)
	if requestId == "" {
		ErrorResponse(resp, "task is missing in parameter", http.StatusBadRequest)
		return
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), processTime)
	defer cancelFunc()
	cfg := config.GetConfig()
	redisInstance, err := redis.Instance(cfg.Redis.Host, cfg.Redis.Password, cfg.Redis.Db, cfg.Redis.Tls)
	if err != nil {
		ErrorResponse(resp, "failed connection to redis", http.StatusFailedDependency)
		return
	}
	cache := redisRepo.NewRepository(redisInstance.Client)
	task, err := services.TaskStatus(ctx, cache, requestId)
	if err != nil {
		ErrorResponse(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	SuccessResponse(resp, task, requestId)
}
