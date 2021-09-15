package repository

import (
	"review_movie/model"

	"gorm.io/gorm"
)

type RepositoryMovie interface {
	CreateMovie(movie model.Movie) (model.Movie, error)
	GetAllMovies() ([]model.Movie, error)
}

type repositorymovie struct {
	db *gorm.DB
}

func NewRepositoryMovie(db *gorm.DB) *repositorymovie {
	return &repositorymovie{db}
}


func(r *repositorymovie) CreateMovie(movie model.Movie) (model.Movie, error) {
	err := r.db.Create(&movie).Error
	if err != nil {
		return movie, err
	}

	return movie, nil
}

func(r *repositorymovie) GetAllMovies() ([]model.Movie, error) {
	var movies []model.Movie
	return movies, nil
}