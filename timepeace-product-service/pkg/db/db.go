package db

import (
	"log"

	"github.com/sgokul961/timepeace-product-service/pkg/config"
	"github.com/sgokul961/timepeace-product-service/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDataBase(c config.Config) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(c.DBUrl), &gorm.Config{})

	if err != nil {

		log.Fatalln(err)
	}
	db.AutoMigrate(&domain.Products{})
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.Brand{})
	return db, err
}
