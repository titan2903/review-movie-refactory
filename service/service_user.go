package service

import (
	"errors"
	"review_movie/entities"
	"review_movie/input"
	"review_movie/repository"

	"golang.org/x/crypto/bcrypt"
)

type ServiceUser interface {
	RegisterUser(input input.RegisterUserInput) (entities.User, error)
	Login(input input.LoginInput) (entities.User, error)
	GetUserByID(ID int) (entities.User, error)
	GetAllUsers() ([]entities.User, error)
	UpdateUser(input input.UpdateUserInput, inputEmail input.FindByEmailInput) (entities.User, error)
	FindUserByEmail(email input.FindByEmailInput) (entities.User, error)
}

type serviceuser struct {
	repository repository.RepositoryUser
}

func NewServiceUser(repository repository.RepositoryUser) *serviceuser {
	return &serviceuser{repository}
}

func(s *serviceuser) RegisterUser(input input.RegisterUserInput) (entities.User, error) {
	user := entities.User{}
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

func(s *serviceuser) Login(input input.LoginInput) (entities.User, error) {
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

func(s *serviceuser) GetUserByID(ID int) (entities.User, error) {
	user, err := s.repository.FindByID(ID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found with that ID")
	}

	return user, nil
}


func(s *serviceuser) GetAllUsers() ([]entities.User, error) {
	users, err := s.repository.GetDataUser()
	if err != nil {
		return users, err
	}

	return users, nil
}

func(s *serviceuser) UpdateUser(input input.UpdateUserInput, inputEmail input.FindByEmailInput) (entities.User, error) {
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

func(s *serviceuser) FindUserByEmail(input input.FindByEmailInput) (entities.User, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if input.Email != user.Email{
		return user, errors.New("Email not match")
	}

	return user, nil
}