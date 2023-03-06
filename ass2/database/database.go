package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GETDB() (*gorm.DB, error) {
	dsn := "user = postgres password=12345 dbname=golang2 host=localhost port=5432 sslmode=disable TimeZone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("I can't connect with database")
	}
	db.AutoMigrate(&User{}, &Product{})

	return db, nil
}
