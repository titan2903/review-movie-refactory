package service

import (
	"review_movie/entities"
	"review_movie/input"
	"review_movie/repository"
)

type ServiceMovie interface {
	CreateMovie(input input.CreateMovieInput) (entities.Movie, error)
	GetAllMovies() ([]entities.Movie, error)
}


type servicemovie struct {
	repository repository.RepositoryMovie
}

func NewServiceMovie(repository repository.RepositoryMovie) *servicemovie {
	return &servicemovie{repository}
}

func(s *servicemovie) CreateMovie(input input.CreateMovieInput) (entities.Movie, error) {
	movie := entities.Movie{}
	movie.Title = input.Title
	movie.Ratings = input.Ratings
	movie.Year = input.Year

	newGenre, err := s.repository.CreateMovie(movie)
	if err != nil {
		return newGenre, err
	}

	return newGenre, nil
}

func(s *servicemovie) GetAllMovies() ([]entities.Movie, error) {
	var movies []entities.Movie
	return movies, nil
}