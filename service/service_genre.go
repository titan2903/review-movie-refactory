package service

import (
	"review_movie/entities"
	"review_movie/input"
	"review_movie/repository"
)

type ServiceGenre interface {
	GetGenres() ([]entities.Genre, error)
	CreateGenre(input input.CreateGenreInput) (entities.Genre, error)
}

type servicegenre struct {
	repository repository.RepositoryGenre
}

func NewServiceGenre(repository repository.RepositoryGenre) *servicegenre {
	return &servicegenre{repository}
}

func(s *servicegenre) GetGenres() ([]entities.Genre, error) {
	genres, err := s.repository.GetGenreList()
	if err != nil {
		return genres, err
	}
	return genres, nil
}

func(s *servicegenre) CreateGenre(input input.CreateGenreInput) (entities.Genre, error) {
	genre := entities.Genre{}
	genre.Name = input.Name

	//! Proses pembuatan slug secara otomatis

	newGenre, err := s.repository.CreateGenre(genre)
	if err != nil {
		return newGenre, err
	}

	return newGenre, nil
}