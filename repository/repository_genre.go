package repository

import (
	"review_movie/entities"

	"gorm.io/gorm"
)

type RepositoryGenre interface {
	GetGenreList() ([]entities.Genre, error)
	CreateGenre(genre entities.Genre) (entities.Genre, error)
}

type repositorygenre struct {
	db *gorm.DB
}

func NewRepositoryGenre(db *gorm.DB) *repositorygenre {
	return &repositorygenre{db}
}

func(r *repositorygenre) GetGenreList() ([]entities.Genre, error) {
	var genres []entities.Genre

	err := r.db.Find(&genres).Error
	if err != nil {
		return genres, err
	}

	return genres, nil
}

func(r *repositorygenre) CreateGenre(genre entities.Genre) (entities.Genre, error) {
	err := r.db.Create(&genre).Error
	if err != nil {
		return genre, err
	}

	return genre, nil
}