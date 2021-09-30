package models

import "github.com/jinzhu/gorm"

type User struct {
    gorm.Model
    Name  string `json:"name"`
    LastName  string `json:"lastName"`
    Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}