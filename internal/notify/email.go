package notify

import (
	"fmt"
	"net/smtp"
	"strings"

	"tfnsysmonitor/internal/config"
)

func SendEmail(serviceName, message string, cfg *config.Config) error {
	c := cfg.Alerts.Emails

	auth := smtp.PlainAuth("", c.SMTPUser, c.SMTPPass, c.SMTPHost)

	subject := fmt.Sprintf("[ALERT] Service Down: %s", serviceName)
	body := fmt.Sprintf("Service: %s\nIssue: %s\n", serviceName, message)

	msg := "From: " + c.SMTPUser + "\n" +
		"To: " + strings.Join(c.Recipients, ",") + "\n" +
		"Subject: " + subject + "\n\n" + body

	addr := fmt.Sprintf("%s:%d", c.SMTPHost, c.SMTPPort)
	return smtp.SendMail(addr, auth, c.SMTPUser, c.Recipients, []byte(msg))
}
