package config

import (
	"log"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(conf *Config) *gorm.DB {

	log.Printf("Connecting to mysql database")
	db, err := gorm.Open(mysql.Open(conf.DatabaseSourceName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Error, While connection mysql database: %v\n", err)
		return nil
	}

	return db

}
