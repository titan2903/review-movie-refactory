package repository

import (
	"review_movie/model"

	"gorm.io/gorm"
)

type RepositoryUser interface {
	CreateUser(user model.User) (model.User, error)
	GetDataUser() ([]model.User, error)
	UpdateUser(user model.User) (model.User, error)
	FindByID(ID int) (model.User, error)
	FindByEmail(email string) (model.User, error)
}

type repositoryuser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repositoryuser {
	return &repositoryuser{db}
}

func(r *repositoryuser) FindByEmail(email string) (model.User, error){
	var user model.User
	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func(r *repositoryuser) FindByID(ID int) (model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", ID).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func(r *repositoryuser) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func(r *repositoryuser) GetDataUser() ([]model.User, error) {
	var users []model.User

	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func(r *repositoryuser) UpdateUser(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

