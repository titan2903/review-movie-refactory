package review

import (
	"review_movie/movie"
	"review_movie/user"
)

type CreateReviewInput struct {
	UserID int `json:"user_id" binding:"required"`
	MovieID int `json:"movie_id" binding:"required"`
	Review string `json:"review" form:"review"`
	Rate int `json:"rate" form:"rate"`
}

type GetReviewInputByID struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
	Movie movie.Movie
}