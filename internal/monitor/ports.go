package monitor

import (
	"net"
	"time"

	"tfnsysmonitor/internal/config"
	"tfnsysmonitor/internal/notify"
)

func CheckPort(cfg *config.Config, p config.PortMonitor) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(p.Host, string(rune(p.Port))), 5*time.Second)
	if err != nil {
		notify.NotifyFailure(p.Name+" Port", p.Name+" port down: "+p.Host, cfg)
		return
	}
	conn.Close()
}
