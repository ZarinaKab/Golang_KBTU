package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GETDB() (*gorm.DB, error) {
	dsn := "postgres://ZarinaKab:DbxWZqF1UI0L@ep-wandering-haze-525394.ap-southeast-1.aws.neon.tech/neondb"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("I can't connect with database")
	}
	db.AutoMigrate(&User{}, &Books{})

	return db, nil
}
