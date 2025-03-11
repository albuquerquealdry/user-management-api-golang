package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Email       string `gorm:"not null;unique"`
	Password    string `gorm:"not null"`
	Birthday    string `gorm:"not null"`
	Address     string `gorm:"not null"`
	Postalcode  int    `gorm:"not null"`
	CPF         int    `gorm:"not null"`
	Nationality string `gorm:"not null"`
	Score       int    `gorm:"not null"`
	Status      string `gorm:"not null"`
	MotherName  string `gorm:"not null"`
}
