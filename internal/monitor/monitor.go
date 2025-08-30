package monitor

import (
	"log"
	"tfnsysmonitor/internal/config"
)

func RunAll(cfg *config.Config) {
	log.Println("Starting monitoring cycle...")

	for _, w := range cfg.Monitors.Websites {
		go CheckWebsite(cfg, w)
	}

	for _, p := range cfg.Monitors.Ports {
		go CheckPort(cfg, p)
	}

	for _, s := range cfg.Monitors.SSL {
		go CheckSSL(cfg, s)
	}

	for _, db := range cfg.Monitors.Databases {
		go CheckDatabase(cfg, db)
	}

	for _, r := range cfg.Monitors.Redis {
		go CheckRedis(cfg, r)
	}

	for _, rmq := range cfg.Monitors.RabbitMQ {
		go CheckRabbitMQ(cfg, rmq)
	}

	for _, d := range cfg.Monitors.Docker {
		go CheckDocker(cfg, d)
	}
}
