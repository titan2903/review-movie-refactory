package review

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ID 					int
	Review 				string
	Rate 				int
	UserID 				int
	MovieID 			int
	CreatedAt      		time.Time
	UpdatedAt	   		time.Time
	DeletedAt			time.Time
}