package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"tfnsysmonitor/internal/config"
)

func SendChat(serviceName, message string, cfg *config.Config) error {
	c := cfg.Alerts.GoogleChat

	text := fmt.Sprintf("[ALERT] %s - %s", serviceName, message)

	data := map[string]string{"text": text}
	jsonData, _ := json.Marshal(data)

	resp, err := http.Post(c.WebhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf(" Google Chat webhook failed: %s", resp.Status)
	}
	return nil
}
