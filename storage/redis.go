package storage

import (
	"context"
	"github.com/redis/go-redis/v9"
	"google.golang.org/appengine/log"
	"time"
)

type CacheBackend interface {
	Store(string, string)

	// Lookup a cached response
	Lookup(string) (string, error)

	// Return the number of items in the cache
	Size() int

	// Flush all records in the store
	Flush()

	Close() error

	Ping(context.Context) error
}

var _ CacheBackend = (*RedisBackend)(nil)

type RedisBackend struct {
	Client redis.UniversalClient
	opt    redis.UniversalOptions
}

type RedisBackendOptions struct {
	RedisOptions redis.Options
	KeyPrefix    string
}

func NewRedisBackend(opt redis.UniversalOptions) *RedisBackend {
	b := &RedisBackend{
		Client: redis.NewUniversalClient(&opt),
		opt:    opt,
	}

	_, err := b.Client.FlushAll(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return b
}

// ctx, cancel := context.WithTimeout(ctx, 90*time.Second)
func (r RedisBackend) Ping(ctx context.Context) error {
	return r.Client.Ping(ctx).Err()

}
func (r RedisBackend) Store(key string, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	if err := r.Client.Set(ctx, key, value, time.Duration(-1)).Err(); err != nil {
		log.Errorf(ctx, "redis error - %s", err.Error())
	}
}

func (r RedisBackend) Lookup(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	val, err := r.Client.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}

	return val, nil
}

func (r RedisBackend) Size() int {
	//TODO implement me
	panic("implement me")
}

func (r RedisBackend) Flush() {
	//TODO implement me
	panic("implement me")
}

func (r RedisBackend) Close() error {
	//TODO implement me
	panic("implement me")
}
