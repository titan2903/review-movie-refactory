package moviegenre

type Service interface {
	CreateMovieGenre(input MovieGenreInput) (MovieGenre, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func(s *service) CreateMovieGenre(input MovieGenreInput) (MovieGenre, error) {
	movie_genre := MovieGenre{}
	movie_genre.MovieID = input.MovieID
	movie_genre.GenreID = input.GenreID

	newGenre, err := s.repository.CreateGenreMovie(movie_genre)
	if err != nil {
		return newGenre, err
	}

	return newGenre, nil
}