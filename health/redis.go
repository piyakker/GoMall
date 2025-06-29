package health

import (
	"context"

	config "github.com/piyakker/GoMall/configs"
	"github.com/redis/go-redis/v9"
)

type RedisChecker struct {
	cfg *config.Config
}

func NewRedisChecker(cfg *config.Config) *RedisChecker {
	return &RedisChecker{cfg: cfg}
}

func (r *RedisChecker) Check(ctx context.Context) error {
	opt, err := redis.ParseURL(r.cfg.RedisAddr)
	if err != nil {
		return err
	}
	client := redis.NewClient(opt)
	defer client.Close()
	return client.Ping(ctx).Err()
}
