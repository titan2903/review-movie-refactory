package service

import (
	"review_movie/entities"
	"review_movie/input"
	"review_movie/repository"
)

type ServiceMovieGenre interface {
	CreateMovieGenre(input input.MovieGenreInput) (entities.MovieGenre, error)
}

type servicemoviegenre struct {
	repository repository.RepositoryMovieGenre
}

func NewServiceMovieGenre(repository repository.RepositoryMovieGenre) *servicemoviegenre {
	return &servicemoviegenre{repository}
}

func(s *servicemoviegenre) CreateMovieGenre(input input.MovieGenreInput) (entities.MovieGenre, error) {
	movie_genre := entities.MovieGenre{}
	movie_genre.MovieID = input.MovieID
	movie_genre.GenreID = input.GenreID

	newGenre, err := s.repository.CreateGenreMovie(movie_genre)
	if err != nil {
		return newGenre, err
	}

	return newGenre, nil
}