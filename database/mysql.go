package database

import (
	"fmt"

	"github.com/onainadapdap1/golang-crud-redis/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionMySQLDB(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUsername, config.DBPassword, config.DBName)
	fmt.Println("dsn : ", dsn)
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	fmt.Println("connected successfully to the database")
	return db
}	
