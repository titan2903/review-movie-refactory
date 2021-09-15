package moviegenre

import "gorm.io/gorm"


type Repository interface {
	CreateGenreMovie(movie_genre MovieGenre) (MovieGenre, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func(r *repository) CreateGenreMovie(movie_genre MovieGenre) (MovieGenre, error) {
	err := r.db.Create(&movie_genre).Error
	if err != nil {
		return movie_genre, err
	}

	return movie_genre, nil
}