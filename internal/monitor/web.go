package monitor

import (
	"net/http"
	"time"

	"tfnsysmonitor/internal/config"
	"tfnsysmonitor/internal/notify"
)

func CheckWebsite(cfg *config.Config, w config.WebsiteMonitor) {
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(w.URL)
	if err != nil || resp.StatusCode != 200 {
		notify.NotifyFailure(w.Name+" Website", w.Name+" website down: "+w.URL, cfg)
	}
}
