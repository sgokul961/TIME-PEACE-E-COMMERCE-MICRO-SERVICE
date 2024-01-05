package db

import (
	"log"

	"github.com/sgokul961/time-peace-admin-service/pkg/config"
	"github.com/sgokul961/time-peace-admin-service/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase(c config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(c.DBUrl), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&domain.Users{})

	return db, err
}
