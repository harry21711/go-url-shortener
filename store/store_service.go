package store

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Define the struct wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

// Top level declarations for the storeService and Redis context
var (
	storeService = &StorageService{}
)

// Initializing the store service and return a store pointer
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "qGFmeU0Zgp/KASrv5bAcsQqaVh2kiQg4Cmbs8TThwnzst8fXjGCOTy+V4qYqg6UfMzMfWuutfo/2ib8B",
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

// Storig this for 6 hours as of now, but it dependes on requirement
const CacheDuration = 6 * time.Hour

/* save the mapping between the originalUrl and the generated shortUrl url */
func SaveUrlMapping(shortUrl string, originalUrl string) error {
	err := storeService.redisClient.Set(shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
	return err
}
