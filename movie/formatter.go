package movie

type MovieFormatter struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Ratings int `json:"ratings"`
	Year int `json:"year"`
}

func FormatCreateMovieResponse(movie Movie) MovieFormatter {
	formatter := MovieFormatter{
		ID: movie.ID,
		Ratings:       movie.Ratings,
		Title: movie.Title,
		Year: movie.Year,
	}

	return formatter
}