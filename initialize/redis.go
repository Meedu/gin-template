package initialize

import (
	"context"

	"github.com/Meedu/gin-template/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.MD_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.MD_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.MD_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.MD_REDIS = client
	}
}
