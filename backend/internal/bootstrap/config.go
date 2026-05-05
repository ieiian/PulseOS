package bootstrap

import (
	"os"

	"gopkg.in/yaml.v3"
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

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
