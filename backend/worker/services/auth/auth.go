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
		return "", err
	}
	if userExists {
		return "", errors.New("user with that email already exists")
	}

	err = hashPassword(&input.Password)
	if err != nil {
		return "", err
	}

	userId, err := s.postgresService.CreateUser(input)
	if err != nil {
		return "", err
	}

	return userId, nil
}

func (s *Service) LoginUser(input entities.LoginUserInput) (entities.UserProfile, error) {
	userExists, err := s.postgresService.DoesUserExist(input.Email)
	if err != nil {
		return entities.UserProfile{}, err
	}
	if !userExists {
		return entities.UserProfile{}, errors.New("user with that email does not exist")
	}

	userPassword, err := s.postgresService.GetUserPassword(input.Email)
	if err != nil {
		return entities.UserProfile{}, err
	}

	isMatch, err := doPasswordsMatch(userPassword, input.Password)
	if err != nil {
		return entities.UserProfile{}, err
	}
	if !isMatch {
		return entities.UserProfile{}, errors.New("password does not match")
	}

	return s.postgresService.GetUserProfile(input.Email)
}
