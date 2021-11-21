package config

import (
	"erajaya/entities/schema"
	"fmt"
	"log"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {
	dbHost := ENV.DbHost
	dbPort := ENV.DbPort
	dbUser := ENV.DbUser
	dbPass := ENV.DbPass
	dbName := ENV.DbName
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	dbConn, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	autoCreate := ENV.DbMigrate
	if autoCreate == "true" {
		fmt.Println("Dropping and recreating all tables")
		dbConn.AutoMigrate(
			&schema.Product{},
		)
		fmt.Println("All tables recreated successfully")
	}

	return dbConn
}
