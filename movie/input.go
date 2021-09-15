package movie


type CreateMovieInput struct {
	Title string `json:"title" form:"title"`
	Year int `json:"year" form:"year"`
	Ratings int `json:"ratings" form:"ratings"`
}