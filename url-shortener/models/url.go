package models

import "github.com/jinzhu/gorm"

type URL struct {
	gorm.Model
	URL            string `json: "url"`
	ShortURL       string `json:"shorturl"`
	Requests       string `json: "request" `
	TimeOfCreation string `json : "timeofcreation"`
}
