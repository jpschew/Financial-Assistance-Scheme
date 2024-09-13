package redis

import (
	"FinancialAssistanceScheme/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var rm RedisManager

type IRedis interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
	IncrBy(ctx context.Context, key string, value int64) *redis.IntCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	MGet(ctx context.Context, keys ...string) *redis.SliceCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Exists(ctx context.Context, keys ...string) *redis.IntCmd

	RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	LPop(ctx context.Context, key string) *redis.StringCmd
	LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd
	LSet(ctx context.Context, key string, index int64, value interface{}) *redis.StatusCmd

	Ping(ctx context.Context) *redis.StatusCmd
	Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd
}

type RedisManager struct {
	client IRedis
}

func (rm *RedisManager) Set(key string, value interface{}, expire_ms int64) error {
	_, err := rm.client.Set(context.Background(), key, value, time.Duration(expire_ms)*time.Millisecond).Result()
	//if err != nil {
	//	//logrus.WithField("error_code", ec.EC_REDIS_SETKEY).Errorf("set redis key %s failed %v", key, err)
	//	log.Printf("set redis key %s failed %v\n", key, err)
	//}
	return err
}

func (rm *RedisManager) GetString(key string) (string, error) {
	return rm.client.Get(context.Background(), key).Result()
}

func (rm *RedisManager) Del(keys ...string) error {
	return rm.client.Del(context.Background(), keys...).Err()
}

// InitRedis initializes the Redis connection
func InitRedis(appConfig config.AppConfig) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", appConfig.Redis.Host, appConfig.Redis.Port),
		Password: appConfig.Redis.Password, // No password set
		DB:       appConfig.Redis.DB,       // Use default DB
	})

	// Test the connection
	ctx := context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
		return
	}

	log.Println("Connected to Redis")
	rm.client = redisClient
}

func GetRedisManager() *RedisManager {
	return &rm
}
