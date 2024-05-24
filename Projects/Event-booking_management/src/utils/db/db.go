package db

import (
	"database/sql"
	"fmt"
	"os"

	"event-booking/config"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	sqlDB, err := sql.Open("postgres", config.Conf.DatabaseURL)
	if err != nil {
		fmt.Println("Unable to open postgres connection.Err:", err)
		os.Exit(1)
	}

	sqlDB.SetConnMaxIdleTime(5)
	sqlDB.SetMaxIdleConns(config.Conf.MaxDBConn)
	sqlDB.SetConnMaxLifetime(10)

	db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("Unable to create gorm Connection.Err:", err)
		os.Exit(1)
	}
}

type DB struct {
	*gorm.DB
}

func New() *DB {
	return &DB{
		DB: db,
	}
}
