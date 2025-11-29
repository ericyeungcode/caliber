package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

var ErrReadCacheData = errors.New("internal cache read error")

type TComputeFunc[T any] func() (*T, error)

func DefaultCacheCtx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	return ctx
}

func StoreRedis(ctx context.Context, redisCli *redis.Client,
	redisKey string, val any, ttl time.Duration) error {

	jsonBuf, err := json.Marshal(val)
	if err != nil {
		log.Errorf("marshal error: %v, key=%v", err, redisKey)
		return ErrReadCacheData
	}

	if err := redisCli.Set(ctx, redisKey, jsonBuf, ttl).Err(); err != nil {
		log.Errorf("redis Set error: %v, key=%v", err, redisKey)
		return ErrReadCacheData
	}

	return nil
}

func GetRedisOrCompute[T any](ctx context.Context, redisCli *redis.Client, redisKey string, ttl time.Duration, computeFunc TComputeFunc[T]) (*T, bool /*fromCache or not */, error) {
	// Try cache first
	buf, err := redisCli.Get(ctx, redisKey).Bytes()
	if err == nil {
		var value T
		if uerr := json.Unmarshal(buf, &value); uerr != nil {
			log.Errorf("cache exists but unmarshal error: %v, key=%v, buf=%v",
				uerr, redisKey, string(buf))
			return nil, false, ErrReadCacheData
		}
		return &value, true, nil
	}

	// Unexpected redis errors: do NOT silently fall back to getter
	if err != redis.Nil {
		log.Errorf("redis Get error: %v, key=%v", err, redisKey)
		return nil, false, ErrReadCacheData
	}

	// Cache miss load from DB or external source
	ptr, err := computeFunc()
	if err != nil {
		log.Errorf("objectGetter error: %v", err)
		return nil, false, ErrReadCacheData
	}

	// Populate cache
	if err := StoreRedis(ctx, redisCli, redisKey, ptr, ttl); err != nil {
		return nil, false, ErrReadCacheData
	}

	return ptr, false, nil
}
