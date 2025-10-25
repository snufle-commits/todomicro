package model

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	Name        string
	Description string
	Completed   bool
	UserID      uint
}
