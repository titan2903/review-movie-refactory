package review

import "gorm.io/gorm"

type Repository interface {
	GetReviewByMovieID(movieID int) ([]Review, error)
	CreateReview(review Review) (Review, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func(r *repository) CreateReview(review Review) (Review, error) {
	err := r.db.Create(&review).Error
	if err != nil {
		return review, err
	}

	return review, nil
}

func(r *repository) GetReviewByMovieID(movieID int) ([]Review, error) {
	var reviews []Review

	err := r.db.Preload("Movie").Where("movie_id = ?", movieID).Order("id desc").Find(&reviews).Error
	if err != nil {
		return reviews, err
	}
	return reviews, nil
}
