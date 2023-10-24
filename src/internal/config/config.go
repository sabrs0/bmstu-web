package config

import (
	"log"
	"os"
	"time"

	cleanenv "github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"` // env-required не запускает, если не задали это поле
	HTTPServer `yaml:"http_server"`
}
type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config { //приставка must - я так помечаю функции, которые не возвращают ошибки, а фаталятся
	cfgPath := os.Getenv("FND_CFG_PATH")
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatalf("Cant load config file %s: %s", cfgPath, err)
	}
	var cfg Config
	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		log.Fatalf("Cant load config file: %s", err)
	}
	return &cfg
}
