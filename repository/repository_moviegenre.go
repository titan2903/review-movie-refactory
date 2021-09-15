package repository

import (
	"review_movie/model"

	"gorm.io/gorm"
)


type RepositoryMovieGenre interface {
	CreateGenreMovie(movie_genre model.MovieGenre) (model.MovieGenre, error)
}

type repositorymoviegenre struct {
	db *gorm.DB
}

func NewRepositoryMovieGenre(db *gorm.DB) *repositorymoviegenre {
	return &repositorymoviegenre{db}
}

func(r *repositorymoviegenre) CreateGenreMovie(movie_genre model.MovieGenre) (model.MovieGenre, error) {
	err := r.db.Create(&movie_genre).Error
	if err != nil {
		return movie_genre, err
	}

	return movie_genre, nil
}