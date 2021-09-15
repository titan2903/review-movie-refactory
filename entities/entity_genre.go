package entities

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	ID 					int
	Name 				string
}