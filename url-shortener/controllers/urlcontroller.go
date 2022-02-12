package controllers

import (
	"net/http"

	"gorm.io/gorm"
)

type URLService struct {
	db *gorm.DB
}

func (URLService *URLService) shortenUrl(w http.ResponseWriter, r *http.Request) {

}

func (URLService *URLService) shortenMultipleUrl(w http.ResponseWriter, r *http.Request) {

}

func (URLService *URLService) deleteUrl(w http.ResponseWriter, r *http.Request) {

}
