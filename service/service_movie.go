package service

import (
	"review_movie/input"
	"review_movie/model"
	"review_movie/repository"
)

type ServiceMovie interface {
	CreateMovie(input input.CreateMovieInput) (model.Movie, error)
	GetAllMovies() ([]model.Movie, error)
}


type servicemovie struct {
	repository repository.RepositoryMovie
}

func NewServiceMovie(repository repository.RepositoryMovie) *servicemovie {
	return &servicemovie{repository}
}

func(s *servicemovie) CreateMovie(input input.CreateMovieInput) (model.Movie, error) {
	movie := model.Movie{}
	movie.Title = input.Title
	movie.Ratings = input.Ratings
	movie.Year = input.Year

	newGenre, err := s.repository.CreateMovie(movie)
	if err != nil {
		return newGenre, err
	}

	return newGenre, nil
}

func(s *servicemovie) GetAllMovies() ([]model.Movie, error) {
	var movies []model.Movie
	return movies, nil
}