package monitor

import (
	"crypto/tls"
	"time"

	"tfnsysmonitor/internal/config"
	"tfnsysmonitor/internal/notify"
)

func CheckSSL(cfg *config.Config, s config.SSLMonitor) {
	conn, err := tls.Dial("tcp", s.Host+":443", nil)
	if err != nil {
		notify.NotifyFailure(s.Name+" SSL", s.Name+" SSL connection failed", cfg)
		return
	}
	defer conn.Close()

	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	if time.Until(expiry).Hours()/24 < 7 {
		notify.NotifyFailure(s.Name+" SSL", s.Name+" SSL expires soon: "+expiry.String(), cfg)
	}
}
