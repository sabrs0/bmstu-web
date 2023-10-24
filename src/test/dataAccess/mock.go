package dataaccessTest

import (
	cfg "github.com/sabrs0/bmstu-web/src/internal/config/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = cfg.MustLoad().String()
var mockDB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
