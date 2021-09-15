package movie

type Service interface {
	CreateMovie(input CreateMovieInput) (Movie, error)
	GetAllMovies() ([]Movie, error)
}


type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func(s *service) CreateMovie(input CreateMovieInput) (Movie, error) {
	movie := Movie{}
	movie.Title = input.Title
	movie.Ratings = input.Ratings
	movie.Year = input.Year

	newGenre, err := s.repository.CreateMovie(movie)
	if err != nil {
		return newGenre, err
	}

	return newGenre, nil
}

func(s *service) GetAllMovies() ([]Movie, error) {
	var movies []Movie
	return movies, nil
}