package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisStore struct {
	client *redis.Client
}

func newRedisClient(network, addr, password string, db int) *RedisStore {
	cli := &RedisStore{
		client: redis.NewClient(&redis.Options{Network: network, Addr: addr, Password: password, DB: db}),
	}
	if err := cli.client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	return cli
}

func (store *RedisStore) Set(ctx context.Context, key string, value interface{}, ttl int) error {
	return store.client.Set(ctx, key, value, time.Duration(ttl)).Err()
}

func (store *RedisStore) Get(ctx context.Context, key string) (data any, ok bool) {
	data = store.client.Get(ctx, key).String()
	if data == "" {
		return
	}
	return data, true
}

func (store *RedisStore) Delete(ctx context.Context, keys []string) error {
	return store.client.Del(ctx, keys...).Err()
}
