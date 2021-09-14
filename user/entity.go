package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             	int
	FullName        string
	Email          	string
	Password   		string
	Role           	string
	CreatedAt      	time.Time
	UpdatedAt	   	time.Time
	DeletedAt		time.Time
}