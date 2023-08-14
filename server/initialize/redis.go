package initialize

import (
	"context"

	"gin-vue-admin/server/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.ECOVACS_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.ECOVACS_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.ECOVACS_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.ECOVACS_REDIS = client
	}
}
