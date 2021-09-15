package input

import (
	"review_movie/entities"
)

type CreateReviewInput struct {
	UserID int `json:"user_id" form:"user_id"`
	MovieID int `json:"movie_id" form:"movie_id" binding:"required"`
	Review string `json:"review" form:"review"`
	Rate int `json:"rate" form:"rate"`
}

type GetReviewInputByID struct {
	ID   int `uri:"id" binding:"required"`
	User entities.User
	Movie entities.Movie
}