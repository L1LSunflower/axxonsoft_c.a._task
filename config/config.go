package config

import (
	"os"
	"strconv"
	"sync"
)

type Config struct {
	Port         int
	TaskLifetime int
	Redis        *Redis
	Logger       *Logger
}

type Redis struct {
	Host     string
	Password string
	Db       int
	Tls      bool
}

type Logger struct {
	level int
}

var (
	configInstance     *Config
	configInstanceOnce sync.Once
)

func GetConfig() *Config {
	if configInstance != nil {
		return configInstance
	}
	configInstanceOnce.Do(func() {
		configInstance = &Config{
			Port:         envInt("PORT", 80),
			TaskLifetime: envInt("TASK_LIFETIME", 15), // Task lifetime represents time in minutes
			Redis: &Redis{
				Host:     envString("REDIS_HOST", "localhost"),
				Password: envString("REDIS_PASSWORD", "secret"),
				Db:       envInt("REDIS_DB", 0),
				Tls:      false,
			},
			Logger: &Logger{
				level: envInt("LOGGER_LEVEL", -4), // Default value means debugger value
			},
		}
	})
	return configInstance
}

func envString(name string, defaultVal string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultVal
	}
	return val
}

func envInt(name string, defaultVal int) int {
	val, err := strconv.Atoi(os.Getenv(name))
	if err != nil {
		return defaultVal
	}
	return val
}
