package sessionService

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis/v3"
	"github.com/simonPacker7/Delta/backend/shared/entities"
	"github.com/simonPacker7/Delta/backend/shared/redisclient"
)

type Service struct {
	Store        *session.Store
	redisService *redisclient.RedisClient
}

func NewService(r *redisclient.RedisClient, cfg *redisclient.RedisConfig) *Service {
	Storage := redis.New(redis.Config{
		Addrs:    []string{cfg.Addr},
		Password: cfg.Password,
		Database: cfg.DB,
	})
	Store := session.New(session.Config{
		Storage:    Storage,
		Expiration: 30 * time.Minute,
	})
	return &Service{
		Store:        Store,
		redisService: r,
	}
}

func (s *Service) SetSession(c *fiber.Ctx, userId string, name string, email string) error {
	sess, err := s.Store.Get(c)
	if err != nil {
		return err
	}

	sess.Set("id", userId)
	sess.Set("name", name)
	sess.Set("email", email)

	return sess.Save()
}

func (s *Service) GetSession(c *fiber.Ctx) (entities.SessionContext, error) {
	sess, err := s.Store.Get(c)
	if err != nil {
		return entities.SessionContext{}, err
	}

	emailValue := sess.Get("email")
	if emailValue == nil {
		return entities.SessionContext{}, errors.New("empty session")
	}
	email := fmt.Sprintf("%v", emailValue)

	nameValue := sess.Get("name")
	if nameValue == nil {
		return entities.SessionContext{}, errors.New("empty session")
	}
	name := fmt.Sprintf("%v", nameValue)

	idValue := sess.Get("id")
	if idValue == nil {
		return entities.SessionContext{}, errors.New("empty session")
	}
	id := fmt.Sprintf("%v", idValue)

	return entities.SessionContext{
		ID:    id,
		Name:  name,
		Email: email,
	}, nil
}

func (s *Service) DestorySession(c *fiber.Ctx) {
	sess, err := s.Store.Get(c)
	if err != nil {
		return
	}

	sess.Destroy()
}
