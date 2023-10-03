package dataaccessTest

import (
	conf_pk "github.com/sabrs0/bmstu-web/src/internal/config_pkg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = conf_pk.Connect_string
var mockDB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
