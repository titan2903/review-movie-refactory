package movie

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	ID 				int
	Ratings 		int
	Title 			string
	Year 			int
	CreatedAt      	time.Time
	UpdatedAt	   	time.Time
	DeletedAt		time.Time
}