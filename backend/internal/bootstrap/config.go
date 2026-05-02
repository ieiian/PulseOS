package bootstrap

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	App struct {
		Name string `yaml:"name"`
		Env  string `yaml:"env"`
	} `yaml:"app"`
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Driver   string `yaml:"driver"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`
	Redis struct {
		Addr string `yaml:"addr"`
	} `yaml:"redis"`
	ObjectStorage struct {
		Endpoint  string `yaml:"endpoint"`
		Bucket    string `yaml:"bucket"`
		AccessKey string `yaml:"access_key"`
		SecretKey string `yaml:"secret_key"`
		UseSSL    bool   `yaml:"use_ssl"`
	} `yaml:"object_storage"`
	AI struct {
		Provider string `yaml:"provider"`
		Model    string `yaml:"model"`
	} `yaml:"ai"`
}

func LoadConfig(path string) (Config, error) {
	var cfg Config

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	parseConfig(&cfg, string(data))

	return cfg, nil
}

func parseConfig(cfg *Config, raw string) {
	var section string

	for _, line := range strings.Split(raw, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		if !strings.HasPrefix(line, " ") && strings.HasSuffix(trimmed, ":") {
			section = strings.TrimSuffix(trimmed, ":")
			continue
		}

		parts := strings.SplitN(trimmed, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.Trim(strings.TrimSpace(parts[1]), "\"")

		switch section {
		case "app":
			switch key {
			case "name":
				cfg.App.Name = value
			case "env":
				cfg.App.Env = value
			}
		case "server":
			if key == "port" {
				cfg.Server.Port = atoi(value)
			}
		case "database":
			switch key {
			case "driver":
				cfg.Database.Driver = value
			case "host":
				cfg.Database.Host = value
			case "port":
				cfg.Database.Port = atoi(value)
			case "name":
				cfg.Database.Name = value
			case "user":
				cfg.Database.User = value
			case "password":
				cfg.Database.Password = value
			case "sslmode":
				cfg.Database.SSLMode = value
			}
		case "redis":
			if key == "addr" {
				cfg.Redis.Addr = value
			}
		case "object_storage":
			switch key {
			case "endpoint":
				cfg.ObjectStorage.Endpoint = value
			case "bucket":
				cfg.ObjectStorage.Bucket = value
			case "access_key":
				cfg.ObjectStorage.AccessKey = value
			case "secret_key":
				cfg.ObjectStorage.SecretKey = value
			case "use_ssl":
				cfg.ObjectStorage.UseSSL = atob(value)
			}
		case "ai":
			switch key {
			case "provider":
				cfg.AI.Provider = value
			case "model":
				cfg.AI.Model = value
			}
		}
	}
}

func atoi(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return result
}

func atob(value string) bool {
	return strings.EqualFold(value, "true")
}
