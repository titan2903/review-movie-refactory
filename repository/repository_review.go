package repository

import (
	"review_movie/model"

	"gorm.io/gorm"
)

type RepositoryReview interface {
	GetReviewByMovieID(movieID int) ([]model.Review, error)
	CreateReview(review model.Review) (model.Review, error)
}

type repositoryreview struct {
	db *gorm.DB
}

func NewRepositoryReview(db *gorm.DB) *repositoryreview {
	return &repositoryreview{db}
}

func(r *repositoryreview) CreateReview(review model.Review) (model.Review, error) {
	err := r.db.Create(&review).Error
	if err != nil {
		return review, err
	}

	return review, nil
}

func(r *repositoryreview) GetReviewByMovieID(movieID int) ([]model.Review, error) {
	var reviews []model.Review

	err := r.db.Preload("Movie").Where("movie_id = ?", movieID).Order("id desc").Find(&reviews).Error
	if err != nil {
		return reviews, err
	}
	return reviews, nil
}
