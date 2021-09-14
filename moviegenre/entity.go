package moviegenre

import (
	"time"

	"gorm.io/gorm"
)

type MovieGenre struct {
	gorm.Model
	ID 				int
	Movie 			string
	Genre 			string
	MovieID 		int
	GenreID 		int
	CreatedAt      	time.Time
	UpdatedAt	   	time.Time
	DeletedAt		time.Time
}