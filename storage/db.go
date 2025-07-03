package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {
	dsn := "host=localhost user=postgres password=postgres dbname=slips port=5432"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("cannot connect to db")
	}

}

func DB() *gorm.DB {
	return db
}
