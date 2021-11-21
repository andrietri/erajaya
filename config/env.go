package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	DbHost    string
	DbPort    string
	DbUser    string
	DbPass    string
	DbName    string
	DbMigrate string
	AppPort   string
}

var ENV Configuration

func LoadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mapENV()
}

func mapENV() {
	ENV.DbHost = os.Getenv("DB_HOST")
	ENV.DbPort = os.Getenv("DB_PORT")
	ENV.DbUser = os.Getenv("DB_USER")
	ENV.DbPass = os.Getenv("DB_PASS")
	ENV.DbName = os.Getenv("DB_NAME")
	ENV.DbMigrate = os.Getenv("DB_AUTO_CREATE")
	ENV.AppPort = os.Getenv("APP_PORT")
}
