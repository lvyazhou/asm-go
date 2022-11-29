package cache

import (
	"asm_platform/infrastructure/config"
	"asm_platform/infrastructure/pkg/slog"
	"context"
	"github.com/go-redis/redis"
	"time"
)

var ctx = context.Background()
var redisClient *redis.Client

// Init redis init
func Init() {
	if err := newRedisDriver(); err != nil {
		slog.Panicf("redis error on database initialization: %s\n", err)
		return
	}
}

func Close() error {
	return redisClient.Close()
}

func Client() *redis.Client {
	return redisClient
}

func newRedisDriver() error {
	conf := config.GetConfig()
	rdb := conf.GetInt("redis.db")
	rsize := conf.GetInt("redis.size")
	retry := conf.GetInt("redis.retry")
	addr := conf.GetString("redis.addr")
	pwd := conf.GetString("redis.pwd")

	redisClient = redis.NewClient(&redis.Options{
		Addr:        addr,             // Redis地址
		Password:    pwd,              // no password set
		DB:          rdb,              // Redis库
		PoolSize:    rsize,            // Redis连接池大小
		MaxRetries:  retry,            // 最大重试次数
		IdleTimeout: 10 * time.Second, // 空闲链接超时时间
	})
	pong, err := redisClient.Ping().Result()
	if err == redis.Nil {
		slog.Errorf("Redis链接异常：%v", err.Error())
		return err
	} else if err != nil {
		slog.Errorf("Redis链接失败：%v", err.Error())
		return err
	} else {
		slog.Infof("Successfully connected to redis database. %v", pong)
		return nil
	}
}

// Set 保存
func Set(key string, value interface{}, expiration time.Duration) (ret string, err error) {
	return redisClient.Set(key, value, expiration).Result()
}

// Get 获取
func Get(key string) (ret string, err error) {
	return redisClient.Get(key).Result()
}

// Exists 是否存在
func Exists(key string) bool {
	c, e := redisClient.Exists(key).Result()
	if e != nil {
		slog.Errorf("run cache exists error, key: %v\n", key)
		return false
	}
	return c > 0
}

// Delete 删除
func Delete(key string) (err error) {
	_, err = redisClient.Del(key).Result()
	slog.Infof("[redis del] success, key: %v\n", key)
	return
}

// Expire 过期时间
func Expire(key string, expiration time.Duration) (bool, error) {
	return redisClient.Expire(key, expiration).Result()
}

//----------------------------------------- 发布订阅 -----------------------------------------------

// Subscribe 订阅
func Subscribe(channels ...string) *redis.PubSub {
	return redisClient.Subscribe(channels...)
}

// Publish 将消息发布到频道。
func Publish(channel string, message interface{}) *redis.IntCmd {
	return redisClient.Publish(channel, message)
}
