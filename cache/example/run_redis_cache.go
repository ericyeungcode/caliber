package main

import (
	"time"

	"github.com/ericyeungcode/caliber/cache"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func RunRedisCacheDemo() {
	ctx := cache.DefaultCacheCtx()

	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:63790",
		DB:   0,
	})

	redisKey := "test:cache_get_or_fill"

	checkVal, err := rdb.Get(ctx, redisKey).Bytes()
	log.Infof("init check value %v, err:%v", string(checkVal), err)

	// delete only this key
	delErr := rdb.Del(ctx, redisKey).Err()
	log.Infof("del err:%v", delErr)

	type MyType struct {
		Value string `json:"value"`
	}

	ttl := 5 * time.Minute

	calls := 0
	getter := func() (*MyType, error) {
		calls++
		return &MyType{Value: "hello world"}, nil
	}

	// 1. First call → cache miss → getter runs
	val1, fromCache1, err := cache.GetRedisOrCompute(ctx, rdb, redisKey, ttl, getter)
	log.Infof("first call, val1:%v, fromCache1:%v, err:%v", val1, fromCache1, err)

	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	if fromCache1 {
		log.Fatalf("expected first call to be from getter, got from cache")
	}
	if val1.Value != "hello world" {
		log.Fatalf("unexpected value: %v", val1.Value)
	}
	if calls != 1 {
		log.Fatalf("getter should have been called once, got %d", calls)
	}

	// 2. Second call → should be cached → getter not executed
	val2, fromCache2, err := cache.GetRedisOrCompute(ctx, rdb, redisKey, ttl, getter)
	log.Infof("second call, val2:%v, fromCache2:%v, err:%v", val2, fromCache2, err)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	if !fromCache2 {
		log.Fatalf("expected second call to be from cache")
	}
	if val2.Value != "hello world" {
		log.Fatalf("unexpected value: %v", val2.Value)
	}
	if calls != 1 {
		log.Fatalf("getter should NOT be called again, got calls=%d", calls)
	}

	// 3rd
	val3, fromCache3, err := cache.GetRedisOrCompute(ctx, rdb, redisKey, ttl, getter)
	log.Infof("third call, val3:%v, fromCache3:%v, err:%v, calls:%v", val3, fromCache3, err, calls)

}
