package genre

import (
	"time"

	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	ID 					int
	Name 				string
	CreatedAt      		time.Time
	UpdatedAt	   		time.Time
	DeletedAt			time.Time
}