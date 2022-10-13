package database

import (
	"assignment-2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "127.0.0.1"
	port     = 5433
	user     = "postgres"
	password = "12345"
	dbname   = "assignment-2"
)

var (
	DB  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host = %s port = %d user = %s "+
		"password = %s dbname= %s sslmode= disable ",
		host, port, user, password, dbname)
	DB, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connection to db:", err)
	}

	DB.Debug().AutoMigrate(models.Order{}, models.Item{})
}

func GetDB() *gorm.DB {
	return DB
}
