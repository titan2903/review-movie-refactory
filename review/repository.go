package review

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetReviewByID(reviewID int) ([]Review, error)
	CreateReview(review Review) (Review, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}


func(r *repository) GetReviewByID(reviewID int) ([]Review, error) {
	var review []Review
	return review, nil
}

func(r *repository) CreateReview(review Review) (Review, error) {
	err := r.db.Create(&review).Error
	if err != nil {
		return review, err
	}

	return review, nil
}
