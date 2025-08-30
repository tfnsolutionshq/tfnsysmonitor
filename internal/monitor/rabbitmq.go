package monitor

import (
	"tfnsysmonitor/internal/config"
	"tfnsysmonitor/internal/notify"

	"github.com/rabbitmq/amqp091-go"
)

func CheckRabbitMQ(cfg *config.Config, r config.RabbitMQMonitor) {
	conn, err := amqp091.Dial(r.URL)
	if err != nil {
		notify.NotifyFailure(r.Name+" RabbitMQ", r.Name+" RabbitMQ unreacheable", cfg)
		return
	}
	defer conn.Close()
}
