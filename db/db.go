package db

import (
	"dashboardapi/config"
	"dashboardapi/db/models"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func RunDb() *gorm.DB {

	user := config.Conf.DB_USERNAME
	host := config.Conf.DB_HOST
	pass := config.Conf.DB_PASSWORD
	dbName := config.Conf.DB_DATABASE
	port := config.Conf.DB_PORT

	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbName)

	sqlDB, err := sql.Open("mysql", DSN)

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute)
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	if err != nil {
		panic("Failed to connect to the database sql.Open")
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database gorm.Open")
	}

	if config.Conf.NODE_ENV == "development" {

		err = db.AutoMigrate(models.User{})
		fmt.Print("migrate")
		if err != nil {
			panic("Failed to migrate tables")
		}
	}

	DB = db
	return db
}
