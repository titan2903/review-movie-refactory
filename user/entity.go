package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             	int
	FullName        string
	Email          	string
	Password   		string
	Role           	string
}