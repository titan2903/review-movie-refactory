package model

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	ID 				int
	Ratings 		int
	Title 			string
	Year 			int
}