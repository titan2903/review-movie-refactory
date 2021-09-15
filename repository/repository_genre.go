package repository

import (
	"review_movie/model"

	"gorm.io/gorm"
)

type RepositoryGenre interface {
	GetGenreList() ([]model.Genre, error)
	CreateGenre(genre model.Genre) (model.Genre, error)
}

type repositorygenre struct {
	db *gorm.DB
}

func NewRepositoryGenre(db *gorm.DB) *repositorygenre {
	return &repositorygenre{db}
}

func(r *repositorygenre) GetGenreList() ([]model.Genre, error) {
	var genres []model.Genre

	err := r.db.Find(&genres).Error
	if err != nil {
		return genres, err
	}

	return genres, nil
}

func(r *repositorygenre) CreateGenre(genre model.Genre) (model.Genre, error) {
	err := r.db.Create(&genre).Error
	if err != nil {
		return genre, err
	}

	return genre, nil
}