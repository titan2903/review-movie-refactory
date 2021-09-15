package moviegenre

type MovieGenreInput struct {
	MovieID int `json:"moviesID" form:"moviesID"`
	GenreID int `json:"genreID" form:"genreID"`
}