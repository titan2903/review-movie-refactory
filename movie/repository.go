package movie

import "gorm.io/gorm"

type Repository interface {
	CreateMovie(movie Movie) (Movie, error)
	GetAllMovies() ([]Movie, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}


func(r *repository) CreateMovie(movie Movie) (Movie, error) {
	err := r.db.Create(&movie).Error
	if err != nil {
		return movie, err
	}

	return movie, nil
}

func(r *repository) GetAllMovies() ([]Movie, error) {
	var movies []Movie
	return movies, nil
}