package formatter

import (
	"review_movie/entities"
	"time"
)

type ReviewFormatter struct {
	ID int `json:"id"`
	Review string `json:"review"`
	Rate int `json:"rate"`
	UserID int `json:"user_id"`
	MovieID int `json:"movies_id"`
	CreatedAt      	time.Time `json:"CreatedAt"`
	UpdatedAt	   	time.Time `json:"UpdatedAt"`
	DeletedAt		*time.Time `json:"DeletedAt"`
}

func FormatCreateReviewResponse(review entities.Review) ReviewFormatter {
	formatter := ReviewFormatter{
		ID: review.ID,
		UserID:         review.UserID,
		MovieID:       review.MovieID,
		Review: review.Review,
		Rate: review.Rate,
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
		DeletedAt: nil,
	}

	return formatter
}