package repository

import (
	"review_movie/entities"

	"gorm.io/gorm"
)

type RepositoryMovie interface {
	CreateMovie(movie entities.Movie) (entities.Movie, error)
	GetAllMovies() ([]entities.Movie, error)
}

type repositorymovie struct {
	db *gorm.DB
}

func NewRepositoryMovie(db *gorm.DB) *repositorymovie {
	return &repositorymovie{db}
}


func(r *repositorymovie) CreateMovie(movie entities.Movie) (entities.Movie, error) {
	err := r.db.Create(&movie).Error
	if err != nil {
		return movie, err
	}

	return movie, nil
}

func(r *repositorymovie) GetAllMovies() ([]entities.Movie, error) {
	var movies []entities.Movie
	return movies, nil
}