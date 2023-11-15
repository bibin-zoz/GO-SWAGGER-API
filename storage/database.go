package storage

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=userdb sslmode=disable password=8596")
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return nil, err
	}

	DB = db
	return DB, nil

}
