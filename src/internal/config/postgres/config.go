package postgres

import (
	"os"
)

type Config struct {
	User     string
	Password string
	DbName   string
	SslMode  string
	Port     string
	Host     string
}

func (c *Config) String() string {
	return "host=" + c.Host + " user=" + c.User + " password=" + c.Password + " dbname=" + c.DbName + " port=" + c.Port + " sslmode=" + c.SslMode
}
func MustLoad() *Config {

	return &Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		Host:     os.Getenv("DB_HOST"),
		SslMode:  os.Getenv("DB_SSL_MODE"),
	}
}
