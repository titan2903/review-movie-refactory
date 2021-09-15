package genre

import "gorm.io/gorm"

type Repository interface {
	GetGenreList() ([]Genre, error)
	CreateGenre(genre Genre) (Genre, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func(r *repository) GetGenreList() ([]Genre, error) {
	var genres []Genre

	err := r.db.Find(&genres).Error
	if err != nil {
		return genres, err
	}

	return genres, nil
}

func(r *repository) CreateGenre(genre Genre) (Genre, error) {
	err := r.db.Create(&genre).Error
	if err != nil {
		return genre, err
	}

	return genre, nil
}