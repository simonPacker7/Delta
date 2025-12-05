package redisclient

import (
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/simonPacker7/Delta/backend/shared/entities"
)

func (r *RedisClient) GetUserPassword(email string) (string, error) {
	recordName := userRecordPrefix + email

	return r.client.HGet(ctx, recordName, "password").Result()
}

func (r *RedisClient) CreateUser(input entities.RegisterUserInput) (string, error) {
	recordName := userRecordPrefix + input.Email
	timeCreated := time.Now().UnixMilli()
	userId := GenerateId()

	if _, err := r.client.Pipelined(ctx, func(rClient redis.Pipeliner) error {
		rClient.HSet(ctx, recordName, input)
		rClient.HSet(ctx, recordName, "timeCreated", timeCreated)
		rClient.HSet(ctx, recordName, "id", userId)
		return nil
	}); err != nil {
		return "", err
	}

	return userId, nil
}

func (r *RedisClient) DoesUserExist(email string) (bool, error) {
	recordName := userRecordPrefix + email
	return r.DoesRecordExist(recordName)
}

func (r *RedisClient) GetUserProfile(email string) (entities.UserProfile, error) {
	recordName := userRecordPrefix + email

	var profileData entities.UserProfile

	err := r.client.HMGet(ctx, recordName, "email", "name", "id").Scan(&profileData)

	if err != nil {
		return entities.UserProfile{}, err
	}

	return profileData, nil
}
