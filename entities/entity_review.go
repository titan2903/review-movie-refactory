package entities

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ID 					int
	Review 				string
	Rate 				int
	UserID 				int
	MovieID 			int
	User				User
	Movie				Movie
}