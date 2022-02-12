package services

import (
	"github.com/jinzhu/gorm"
)

type DB struct {
	dbConnection *gorm.DB
}

func (db *DB) create(body interface{}, url string) {

}

func (db *DB) delete(id string, url string) {

}

func (db *DB) update(id string, body interface{}, model string) {

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
