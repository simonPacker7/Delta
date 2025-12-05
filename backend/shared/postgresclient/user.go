package postgresclient

import (
	"context"
	"time"

	"github.com/simonPacker7/Delta/backend/shared/entities"
)

func (p *PostgresClient) GetUserPassword(email string) (string, error) {
	var hashed_password string
	err := p.client.QueryRow(context.Background(), "select hashed_password from users where email=$1", email).Scan(&hashed_password)
	return hashed_password, err
}

func (p *PostgresClient) CreateUser(input entities.RegisterUserInput) (string, error) {
	timeCreated := time.Now()
	id := GenerateId()

	queryString := `insert into users (
			id,
			username,
			email,
			hashed_password,
			created_at
		)
		values (
			$1,
			$2,
			$3,
			$4,
			$5
		)
		`
	_, err := p.client.Exec(context.Background(), queryString,
		id, input.Name, input.Email, input.Password, timeCreated)
	return id.String(), err
}

func (p *PostgresClient) GetUserProfile(email string) (entities.UserProfile, error) {
	var profile entities.UserProfile
	queryString := `select id, username, email from users where email=$1`
	err := p.client.QueryRow(context.Background(), queryString, email).Scan(&profile.ID, &profile.Name, &profile.Email)
	return profile, err
}

func (p *PostgresClient) DoesUserExist(email string) (bool, error) {
	var exists bool
	queryString := "select exists(select 1 from users where email=$1)"
	err := p.client.QueryRow(context.Background(), queryString, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
