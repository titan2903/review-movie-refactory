package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	GetUserByID(ID int) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(input UpdateUserInput, inputEmail FindByEmailInput) (User, error)
	FindUserByEmail(email FindByEmailInput) (User, error)
}

type service struct { //! memanggil repository
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func(s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.FullName = input.FullName
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	if input.Role == "admin" || input.Role == "user" {
		user.Role = input.Role
	}

	newUser, err := s.repository.CreateUser(user)

	if err != nil {
		return newUser ,err
	}

	return newUser, nil
}

func(s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func(s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found with that ID")
	}

	return user, nil
}


func(s *service) GetAllUsers() ([]User, error) {
	users, err := s.repository.GetDataUser()
	if err != nil {
		return users, err
	}

	return users, nil
}

func(s *service) UpdateUser(input UpdateUserInput, inputEmail FindByEmailInput) (User, error) {
	user, err := s.repository.FindByEmail(inputEmail.Email)
	if err != nil {
		return user, err
	}

	if input.Email != "" {
		user.FullName = input.FullName
	} else if input.FullName != "" {
		user.Email = input.Email
	}

	updatedUser, err := s.repository.UpdateUser(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func(s *service) FindUserByEmail(input FindByEmailInput) (User, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if input.Email != user.Email{
		return user, errors.New("Email not match")
	}

	return user, nil
}