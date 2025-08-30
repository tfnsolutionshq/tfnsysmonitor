package main

import (
	"flag"
	"log"
	"time"

	"github.com/coreos/go-systemd/daemon"

	"tfnsysmonitor/internal/config"
	"tfnsysmonitor/internal/monitor"
)

func main() {
	// Command-line flag for config file
	configPath := flag.String("config", "./config.yaml", "Path to configuration YAML file")
	flag.Parse()

	// Load configuration
	cfg := config.MustLoadConfig(*configPath)

	// Monitoring loop interval
	interval := time.Duration(cfg.IntervalSeconds) * time.Second
	log.Printf("Starting monitoring loop (interval: %v)...", interval)

	// Notify systemd that we are ready
	sent, err := daemon.SdNotify(false, daemon.SdNotifyReady)
	if err != nil {
		log.Printf("Failed to send systemd READY notification: %v", err)
	} else if sent {
		log.Println("Systemd READY notification sent")
	}

	for {
		// Website monitors
		for _, w := range cfg.Monitors.Websites {
			go monitor.CheckWebsite(cfg, w)
		}

		// Database monitors
		for _, d := range cfg.Monitors.Databases {
			go monitor.CheckDatabase(cfg, d)
		}

		// Redis monitors
		for _, r := range cfg.Monitors.Redis {
			go monitor.CheckRedis(cfg, r)
		}

		// RabbitMQ monitors
		for _, q := range cfg.Monitors.RabbitMQ {
			go monitor.CheckRabbitMQ(cfg, q)
		}

		// Port monitors
		for _, p := range cfg.Monitors.Ports {
			go monitor.CheckPort(cfg, p)
		}

		// SSL monitors
		for _, s := range cfg.Monitors.SSL {
			go monitor.CheckSSL(cfg, s)
		}

		// Docker monitors
		for _, d := range cfg.Monitors.Docker {
			go monitor.CheckDocker(cfg, d)
		}

		// Notify systemd watchdog
		_, _ = daemon.SdNotify(false, daemon.SdNotifyWatchdog)

		time.Sleep(interval)
	}
}
