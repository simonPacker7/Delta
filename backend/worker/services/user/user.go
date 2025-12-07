package userService

import (
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

func (s *Service) GetUserProfile(email string) (entities.UserProfile, error) {
	return s.postgresService.GetUserProfile(email)
}
