package postgresclient

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConfig struct {
	Addr     string
	DB       string
	Username string
	Password string
}

type PostgresClient struct {
	client *pgxpool.Pool
}

func NewPostgresClient(cfg PostgresConfig) (*PostgresClient, error) {
	url := "postgres://" + cfg.Username + ":" + cfg.Password + "@" + cfg.Addr + "/" + cfg.DB
	conn, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return &PostgresClient{}, err
	}

	return &PostgresClient{client: conn}, nil
}

func GenerateId() uuid.UUID {
	id, _ := uuid.NewV7()
	return id
}
