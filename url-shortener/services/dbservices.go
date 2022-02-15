package services

import (
	"github.com/jinzhu/gorm"
)

type DB struct {
	dbConnection *gorm.DB
}

func (db *DB) Create(body interface{}) {
	db.dbConnection.Create(&body)
}

func (db *DB) Delete(body interface{}, id string, url string) {
	db.dbConnection.Delete(&body, 1)
}

func (db *DB) Update(id string, body interface{}, model string) {
	db.dbConnection.Model(&body).Update(body)
}

func (db *DB) getbyID(id string) {

}

func (db *DB) getbyUrl(url string) {

}

func (db *DB) getbyParameter() {

}

func (db *DB) deleteALl() {

}

func (db *DB) AddMany() {

}
