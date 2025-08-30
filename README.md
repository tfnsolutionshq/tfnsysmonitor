# TFN System Monitor (`tfnsysmonitor`)

`tfnsysmonitor` is a lightweight system monitoring daemon written in Go.  
It monitors websites, databases, Redis, RabbitMQ, ports, SSL certificates, and Docker containers.  
If failures are detected, it sends notifications (email, SMS, or chat) based on your configuration.  

It is designed to run under **systemd**, with `sd_notify` support so that systemd can restart the service if it fails.

## 🚀 Features

- Website availability monitoring (HTTP/HTTPS).
- Database connection checks.
- Redis, RabbitMQ service checks.
- Port availability monitoring.
- SSL certificate expiry monitoring.
- Docker container monitoring.
- Notifications via Email, SMS, and Chat (configurable).
- Systemd integration with automatic restart on failure.

## 📦 Installation

Clone the repository and build from source

```bash
git clone https://github.com/tfnsolutionshq/tfnsysmonitor.git
cd tfnsysmonitor
go build -o tfnsysmonitor
```

Install binary and man page

```bash
sudo cp tfnsysmonitor /usr/local/bin/
sudo cp docs/tfnsysmonitor.1 /usr/share/man/man1/
sudo mandb
```

Install systemd service

```bash
sudo cp packaging/tfnsysmonitor.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable tfnsysmonitor
sudo systemctl start tfnsysmonitor
```

## ⚙️ Configuration

The main config file lives at `/etc/tfnsysmonitor/config.yaml`
Edit the `config.yaml` file to specify monitoring targets and notification methods.

Example:

```bash
interval_seconds: 60
notifier:
  type: slack
  webhook_url: "https://hooks.slack.com/services/XXXX/YYYY/ZZZZ"

monitors:
  websites:
    - name: Afriwok API
      url: "https://api.afriwok.com/health"

  databases:
    - name: MainDB
      driver: postgres
      dsn: "host=localhost user=postgres password=secret dbname=afriwok sslmode=disable"

  redis:
    - name: Cache
      address: "localhost:6379"

  rabbitmq:
    - name: MQ
      url: "amqp://guest:guest@localhost:5672/"

  ports:
    - name: SSH
      address: "localhost:22"

  ssl:
    - name: Afriwok SSL
      url: "https://afriwok.com"
      days_before_expiry: 15

  docker:
    - name: Laravel Service
      container: "afriwok_laravel"
```

## 🚀 Running with systemd

1. Copy the `tfnsysmonitor.service` to the location systemd service location

    ```bash
    sudo cp tfnsysmonitor.service /etc/systemd/system/tfnsysmonitor.service
    ```

2. Enable and start the service

   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable tfnsysmonitor
   sudo systemctl start tfnsysmonitor
   ```

3. Check the status of the `tfnsysmonitor` service

   ```bash
   sudo systemctl status tfnsysmonitor
   journalctl -u tfnsysmonitor -f
   ```

## 🖥 Usage

Start manually

```bash
/usr/local/bin/tfnsysmonitor
```

Start/Stop via systemd to run as a background service

```bash
sudo systemctl start tfnsysmonitor
sudo systemctl stop tfnsysmonitor
```

Reload config (restart service)

```bash
sudo systemctl restart tfnsysmonitor
```

Run directly

```bash
./tfnsysmonitor -config config.yaml
```

## 🔍 Logs

Logs are written to journalctl when run under systemd

```bash
journalctl -u tfnsysmonitor -f
```

## 🛠 Development

Build locally

```bash
go build -o tfnsysmonitor
```

Run with debugging

```bash
go run . -config config.yaml
```

## 🧑‍💻 Contributing

- Fork repo, submit PRs

- Add new monitors under `internal/monitor/`

- Update `config.yaml` schema

## 📄 License

MIT License.

© 2025 TurboFlux Network Solutions Ltd.
