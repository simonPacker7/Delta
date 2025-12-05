package redisclient

import (
	"context"
	"math/rand"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type RedisClient struct {
	client *redis.Client
}

var ctx = context.Background()

const userRecordPrefix = "user:"
const gameSessionRecordPrefix = "game:session:"
const gamePrivateSessionRecordPrefix = "game:private:session:"
const matchmakingQueuePrefix = "game:open"
const gameExpiredQueuePrefix = "game:expired:queue"

func NewRedisClient(cfg RedisConfig) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &RedisClient{client: rdb}
}

func GenerateId() string {
	id, _ := uuid.NewV7()
	return id.String()
}

// Returns a random alphanumeric string thats 8 characters long
func GenerateGameCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Create a byte slice to hold the random data
	bytes := make([]byte, 8)

	for i := range bytes {
		bytes[i] = charset[rand.Intn(len(charset))]
	}

	return string(bytes)
}

func (r *RedisClient) DoesRecordExist(recordName string) (bool, error) {
	val, err := r.client.Exists(ctx, recordName).Result()

	if err != nil {
		return false, err
	}

	if val == 1 {
		return true, nil
	} else {
		return false, nil
	}
}
