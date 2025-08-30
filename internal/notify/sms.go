package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"tfnsysmonitor/internal/config"
)

type TermiiRequest struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Sms     string `json:"sms"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	ApiKey  string `json:"api_key"`
}

func SendSMS(serviceName, message string, cfg *config.Config) error {
	c := cfg.Alerts.SMS
	text := fmt.Sprintf("[ALERT] %s - %s", serviceName, message)

	for _, recipient := range c.Recipients {
		reqBody := TermiiRequest{
			To:      recipient,
			From:    c.Sender,
			Sms:     text,
			Type:    "plain",
			Channel: "generic",
			ApiKey:  c.APIKey,
		}

		jsonData, _ := json.Marshal(reqBody)

		resp, err := http.Post("https://api.ng.termii.com/api/sms/send", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			return err
		}
		defer resp.Body.Close()
	}
	return nil
}
