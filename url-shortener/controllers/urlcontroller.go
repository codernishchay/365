package controllers

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener/models"

	"gorm.io/gorm"
)

type URLService struct {
	db *gorm.DB
}

func (URLService *URLService) shortenUrl(w http.ResponseWriter, r *http.Request) {
	var url models.URL
	json := json.NewDecoder(r.Body).Decode(&url)
	fmt.Println(json)
	fmt.Println(url.URL)
	s := sha1.New()
	s.Write([]byte(url.URL))
	h := s.Sum(nil)
	fmt.Println(h)

	// pass this object to database and check if this is url is unique
}

func (URLService *URLService) shortenMultipleUrl(w http.ResponseWriter, r *http.Request) {

}

func (URLService *URLService) deleteUrl(w http.ResponseWriter, r *http.Request) {

}
