package service

import (
	"review_movie/input"
	"review_movie/model"
	"review_movie/repository"
)


type ServiceReview interface {
	CreateReview(input input.CreateReviewInput) (model.Review, error)
	GetReviewMovie() ([]model.Review, error)
}


type servicereview struct {
	repository repository.RepositoryReview
}

func NewServiceReview(repository repository.RepositoryReview) *servicereview {
	return &servicereview{repository}
}

func(s *servicereview) CreateReview(input input.CreateReviewInput) (model.Review, error) {
	review := model.Review{}
	review.UserID = input.UserID
	review.MovieID = input.MovieID
	review.Rate = input.Rate
	review.Review = input.Review

	newGenre, err := s.repository.CreateReview(review)
	if err != nil {
		return newGenre, err
	}

	return newGenre, nil
}

func(s *servicereview) GetReviewMovie() ([]model.Review, error) {
	var review []model.Review
	return review, nil
}