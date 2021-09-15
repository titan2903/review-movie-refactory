package review

import (
	"review_movie/movie"
	"review_movie/user"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ID 					int
	Review 				string
	Rate 				int
	UserID 				int
	MovieID 			int
	User				user.User
	Movie				movie.Movie
}