package review


type Service interface {
	CreateReview(input CreateReviewInput) (Review, error)
	GetReviewMovie() ([]Review, error)
}


type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func(s *service) CreateReview(input CreateReviewInput) (Review, error) {
	review := Review{}
	review.UserID = input.UserID
	review.MovieID = input.MovieID
	review.Rate = input.Rate
	review.Review = input.Review

	newGenre, err := s.repository.CreateReview(review)
	if err != nil {
		return newGenre, err
	}

	return newGenre, nil
}

func(s *service) GetReviewMovie() ([]Review, error) {
	var review []Review
	return review, nil
}