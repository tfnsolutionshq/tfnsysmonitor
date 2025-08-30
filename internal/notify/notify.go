package notify

import (
	"log"

	"tfnsysmonitor/internal/config"
)

func NotifyFailure(serviceName, message string, cfg *config.Config) {
	if cfg.Alerts.Emails.Enabled {
		err := SendEmail(serviceName, message, cfg)
		if err != nil {
			log.Printf("[ALERT][EMAIL] failed: %v", err)
		}
	}

	if cfg.Alerts.SMS.Enabled {
		err := SendSMS(serviceName, message, cfg)
		if err != nil {
			log.Printf("[ALERT][SMS] failed: %v", err)
		}
	}

	if cfg.Alerts.GoogleChat.Enabled {
		err := SendChat(serviceName, message, cfg)
		if err != nil {
			log.Printf("[ALERT][CHAT] failed: %v", err)
		}
	}

	log.Printf("[ALERT][%s] %s. Loaded config: %+v", serviceName, message, cfg)
}
