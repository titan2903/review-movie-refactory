package genre

type Service interface {
	GetGenres() ([]Genre, error)
	CreateGenre(input CreateGenreInput) (Genre, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func(s *service) GetGenres() ([]Genre, error) {
	genres, err := s.repository.GetGenreList()
	if err != nil {
		return genres, err
	}
	return genres, nil
}

func(s *service) CreateGenre(input CreateGenreInput) (Genre, error) {
	genre := Genre{}
	genre.Name = input.Name

	//! Proses pembuatan slug secara otomatis

	newGenre, err := s.repository.CreateGenre(genre)
	if err != nil {
		return newGenre, err
	}

	return newGenre, nil
}