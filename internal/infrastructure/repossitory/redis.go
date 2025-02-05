package repository
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	"github.com/go-redis/redis/v8"
)

var _ contracts.CacheInterface = (*RedisService)(nil)

type RedisService struct {
	client *redis.Client
}

func NewRedisService(addr, password string, db int) (*RedisService, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("error connect Redis: %v", err)
	}

	log.Println("Successful connection")
	return &RedisService{client: client}, nil
}

func (r *RedisService) SetCache(ctx context.Context, key string, value map[string]interface{}, expiration *time.Duration) error {
	var err error
	var expiration_value time.Duration

	detailsJSON, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("error serializer value to JSON: %v", err)
	}
	
	if expiration != nil {
		expiration_value = *expiration
	}else {expiration_value = 0}

	err = r.client.Set(ctx, key, detailsJSON, expiration_value).Err()

	if err != nil {
		return fmt.Errorf("error set value cache: %v", err)
	}

	log.Printf("Value %s set witch key %s Redis.\n", value, key)
	return nil
}

func (r *RedisService) DeleteCache(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("error delete key %s cache: %v", key, err)
	}
	log.Printf("Key %s deleted in Redis.\n", key)
	return nil
}

func (r *RedisService) GetCache(ctx context.Context, key string) (map[string]interface{}, error) {
	data, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("key %s not found in cache", key)
	} else if err != nil {
		return nil, fmt.Errorf("error  search key %s cache: %v", key, err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return nil, fmt.Errorf("error deserializer value in cache: %v", err)
	}

	return result, nil
}
