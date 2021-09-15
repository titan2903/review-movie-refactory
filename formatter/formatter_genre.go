package formatter

import (
	"review_movie/entities"
	"time"
)

type GenreFormatter struct {
	ID int `json:"ID"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt"`
}

type GenreFormatterListGenre struct {
	ID int `json:"ID"`
	Name string `json:"name"`
}

func FormatGenre(genre entities.Genre) GenreFormatter {
	formatter := GenreFormatter{}
	formatter.ID = genre.ID
	formatter.Name = genre.Name
	formatter.CreatedAt = genre.CreatedAt
	formatter.UpdatedAt = genre.UpdatedAt
	formatter.DeletedAt = nil

	return formatter
}

func FormatGenreGetResponse(genre entities.Genre) GenreFormatterListGenre {
	formatter := GenreFormatterListGenre{}
	formatter.ID = genre.ID
	formatter.Name = genre.Name

	return formatter
}

func FormatGenres(genres []entities.Genre) []GenreFormatterListGenre {
	if len(genres) == 0 {
		return []GenreFormatterListGenre{}
	}

	var genreFormatter []GenreFormatterListGenre

	for _, genre := range genres {
		formatter := FormatGenreGetResponse(genre)
		genreFormatter = append(genreFormatter, formatter)
	}

	return genreFormatter
}