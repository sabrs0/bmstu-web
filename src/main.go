package main

import (
	//"github.com/sabrs0/bmstu-web/src/views"

	conf_pk "github.com/sabrs0/bmstu-web/src/config_pkg"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := conf_pk.Connect_string
	/*db*/ _, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err == nil {
		//views.Gtk_init(db)
	}

}
