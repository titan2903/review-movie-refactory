package repository

import (
	"review_movie/entities"

	"gorm.io/gorm"
)

type RepositoryUser interface {
	CreateUser(user entities.User) (entities.User, error)
	GetDataUser() ([]entities.User, error)
	UpdateUser(user entities.User) (entities.User, error)
	FindByID(ID int) (entities.User, error)
	FindByEmail(email string) (entities.User, error)
}

type repositoryuser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repositoryuser {
	return &repositoryuser{db}
}

func(r *repositoryuser) FindByEmail(email string) (entities.User, error){
	var user entities.User
	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func(r *repositoryuser) FindByID(ID int) (entities.User, error) {
	var user entities.User
	err := r.db.Where("id = ?", ID).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func(r *repositoryuser) CreateUser(user entities.User) (entities.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func(r *repositoryuser) GetDataUser() ([]entities.User, error) {
	var users []entities.User

	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func(r *repositoryuser) UpdateUser(user entities.User) (entities.User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

