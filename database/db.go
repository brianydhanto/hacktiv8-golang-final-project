package database

import (
	"fmt"
	"hacktiv8-golang-final-project/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "root"
	dbPort   = "5432"
	dbname   = "golangfp"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("errror connecting to database :", err)
	}

	fmt.Println("sukses koneksi ke database")
	db.AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})

}

func GetDB() *gorm.DB {
	return db
}
