package moviegenre

import (
	"gorm.io/gorm"
)

type MovieGenre struct {
	gorm.Model
	ID 				int
	Movie 			string
	Genre 			string
	MovieID 		int
	GenreID 		int
}