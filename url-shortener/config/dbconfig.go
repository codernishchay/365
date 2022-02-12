package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "url-shortener"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func DBConnect() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password=password dbname=url-shortener port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	CheckError(err)

	fmt.Println("Connected!")
	return db, err
}
