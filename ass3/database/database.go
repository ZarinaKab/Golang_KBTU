package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GETDB() (*gorm.DB, error) {
	dsn := "url"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("I can't connect with database")
	}
	db.AutoMigrate(&User{}, &Books{})

	return db, nil
}
