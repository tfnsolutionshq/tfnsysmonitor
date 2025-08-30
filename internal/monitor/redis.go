package monitor

import (
	"context"
	"strconv"

	"tfnsysmonitor/internal/config"
	"tfnsysmonitor/internal/notify"

	"github.com/redis/go-redis/v9"
)

func CheckRedis(cfg *config.Config, r config.RedisMonitor) {
	client := redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + strconv.Itoa(r.Port),
		Password: r.Password,
		DB:       r.DB,
	})
	defer client.Close()

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		notify.NotifyFailure(r.Name+" Redis", r.Name+" Redis unreacheable", cfg)
	}
}
