package cache

import (
	"context"

	"github.com/YogeLiu/CloudDisk/pkg/conf"
)

var Store Driver

func Init() {
	if conf.RedisConfig.Password != "" && conf.RedisConfig.Addr != "" {
		Store = newRedisClient(conf.RedisConfig.Network, conf.RedisConfig.Addr, conf.RedisConfig.Password, conf.RedisConfig.DB)
	} else {
		Store = NewMemoryStore()
	}
}

// Driver 键值缓存存储容器
type Driver interface {
	Set(ctx context.Context, key string, value any, ttl int) error
	Get(ctx context.Context, key string) (any, bool)
	// Gets(ctx context.Context, keys []string) (map[string]interface{}, []string)
	// Sets(ctx context.Context, values map[string]interface{}, prefix string) error
	Delete(ctx context.Context, keys []string) error
}
