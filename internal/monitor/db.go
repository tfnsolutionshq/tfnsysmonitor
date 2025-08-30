package monitor

import (
	"database/sql"
	"tfnsysmonitor/internal/config"
	"tfnsysmonitor/internal/notify"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func CheckDatabase(cfg *config.Config, d config.DatabaseMonitor) {
	db, err := sql.Open(d.Driver, d.DSN)
	if err != nil {
		notify.NotifyFailure(d.Name+" DB", d.Name+" DB Connection failed", cfg)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		notify.NotifyFailure(d.Name+" DB", d.Name+" DB Ping failed", cfg)
	}
}
