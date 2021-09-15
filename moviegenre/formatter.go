package moviegenre

import (
	"time"
)

type MovieGenreFormatter struct {
	ID int `json:"ID"`
	MovieID int `json:"movie_id"`
	GenreID int `json:"genre_id"`
	Genre string `json:"genre"`
	CreatedAt      	time.Time `json:"CreatedAt"`
	UpdatedAt	   	time.Time `json:"UpdatedAt"`
	DeletedAt		*time.Time `json:"DeletedAt"`
}

func FormatCreateMovieGenreResponse(movie_genre MovieGenre) MovieGenreFormatter {
	formatter := MovieGenreFormatter{
		ID: movie_genre.ID,
		MovieID:       movie_genre.MovieID,
		GenreID: movie_genre.GenreID,
		Genre: movie_genre.Genre,
		CreatedAt: movie_genre.CreatedAt,
		UpdatedAt: movie_genre.UpdatedAt,
		DeletedAt: nil,
	}

	return formatter
}