package contracts

import (
	"context"
	"time"
)

type ICache interface {
	SetCache(ctx context.Context, key string, value map[string]interface{}, expiration *time.Duration) error
	GetCache(ctx context.Context, key string) (map[string]interface{}, error)
	DeleteCache(ctx context.Context, key string) error
}
