package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Digisata/dts-hactiv8-golang-chap3/docs"
	"github.com/Digisata/dts-hactiv8-golang-chap3/models"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DEBUG_MODE = true
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	schemes := []string{"https"}

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	env := os.Getenv("ENV")
	if env == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		port = os.Getenv("PORT")
		host = fmt.Sprintf("%s:%s", os.Getenv("HOST"), port)
		schemes = []string{"http"}

		dbHost = os.Getenv("DB_HOST")
		dbPort = os.Getenv("DB_PORT")
		dbName = os.Getenv("DB_NAME")
		dbUsername = os.Getenv("DB_USERNAME")
		dbPassword = os.Getenv("DB_PASSWORD")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUsername, dbPassword, dbName)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	docs.SwaggerInfo.Version = os.Getenv("DOCS_VERSION")
	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.BasePath = os.Getenv("BASE_PATH")
	docs.SwaggerInfo.Schemes = schemes

	db.Debug().AutoMigrate(&models.User{}, &models.SocialMedia{}, &models.Photo{}, &models.Comment{})
}

func GetDB() *gorm.DB {
	if DEBUG_MODE {
		return db.Debug()
	}

	return db
}
