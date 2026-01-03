package authService

import (
	"errors"

	"github.com/simonPacker7/Delta/backend/shared/entities"
	"github.com/simonPacker7/Delta/backend/shared/postgresclient"
)

type Service struct {
	postgresService *postgresclient.PostgresClient
}

func NewService(p *postgresclient.PostgresClient) *Service {
	return &Service{
		postgresService: p,
	}
}

func (s *Service) RegisterUser(input entities.RegisterUserInput) (string, error) {
	userExists, err := s.postgresService.DoesUserExist(input.Email)
	if err != nil {
		return "", errors.New("A user with that email already exists")
	}
	if userExists {
		return "", errors.New("A user with that email already exists")
	}

	err = hashPassword(&input.Password)
	if err != nil {
		return "", errors.New("Error registering user, please try again")
	}

	userId, err := s.postgresService.CreateUser(input)
	if err != nil {
		return "", errors.New("Error registering user, please try again")
	}

	return userId, nil
}

func (s *Service) LoginUser(input entities.LoginUserInput) (entities.UserProfile, error) {
	userExists, err := s.postgresService.DoesUserExist(input.Email)
	if err != nil {
		return entities.UserProfile{}, errors.New("Invalid email or password")
	}
	if !userExists {
		return entities.UserProfile{}, errors.New("Invalid email or password")
	}

	userPassword, err := s.postgresService.GetUserPassword(input.Email)
	if err != nil {
		return entities.UserProfile{}, errors.New("Invalid email or password")
	}

	isMatch, err := doPasswordsMatch(userPassword, input.Password)
	if err != nil {
		return entities.UserProfile{}, errors.New("Invalid email or password")
	}
	if !isMatch {
		return entities.UserProfile{}, errors.New("Invalid email or password")
	}

	return s.postgresService.GetUserProfile(input.Email)
}
