package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func ConnectDatabase() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("error connecting to database")
	}
}
