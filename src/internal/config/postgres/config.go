package postgres

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db_name" `
	SslMode  string `yaml:"ssl_mode" env-default:"disable"`
	Port     string `yaml:"port" env-default:"5432"`
	Host     string `yaml:"host"`
}

func (c *Config) String() string {
	return "host=" + c.Host + "user=" + c.User + "password=" + c.Password + "dbname=" + c.DbName + "port=" + c.Port + "sslmode=" + c.SslMode
}
func MustLoad() *Config {
	cfgPath := os.Getenv("FND_DB_CFG_PATH")
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatalf("Cant load db config file %s: %s", cfgPath, err)
	}
	var cfg Config
	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		log.Fatalf("Cant load db config file: %s", err)
	}
	return &cfg
}
