package initializers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DATABASE *gorm.DB

func ConnectionDb() {
	var err error
	dsn := os.Getenv("DB_URL")
	DATABASE, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("impossible de se connecter a la base de donnée")
	}
}
