package config

import (
	"assignment2/pkg/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	host   string = "localhost"
	port   int    = 5432
	user   string = "postgres"
	pass   string = "root"
	dbname string = "assignment2"
)

func StartDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	if !db.HasTable(&models.Order{}) {
		db.AutoMigrate(&models.Order{})
	}

	if !db.HasTable(&models.Item{}) {
		db.AutoMigrate(&models.Item{})
	}

	log.Default().Println("Connection to Database is Successfull")

	return db, nil
}