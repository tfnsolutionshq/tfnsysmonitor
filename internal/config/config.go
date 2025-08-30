package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	IntervalSeconds int `yaml:"interval_seconds"`

	Alerts struct {
		Emails     EmailConfig `yaml:"emails"`
		SMS        SMSConfig   `yaml:"sms"`
		GoogleChat ChatConfig  `yaml:"google_chat"`
	} `yaml:"alerts"`

	Monitors struct {
		Websites  []WebsiteMonitor  `yaml:"websites"`
		Ports     []PortMonitor     `yaml:"ports"`
		SSL       []SSLMonitor      `yaml:"ssl"`
		Databases []DatabaseMonitor `yaml:"databases"`
		Redis     []RedisMonitor    `yaml:"redis"`
		RabbitMQ  []RabbitMQMonitor `yaml:"rabbitmq"`
		Docker    []DockerMonitor   `yaml:"docker"`
	} `yaml:"monitors"`
}

type WebsiteMonitor struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type DatabaseMonitor struct {
	Name   string `yaml:"name"`
	Driver string `yaml:"driver"`
	DSN    string `yaml:"dsn"`
}

type RedisMonitor struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type RabbitMQMonitor struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type PortMonitor struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type SSLMonitor struct {
	Name          string `yaml:"name"`
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	DaysThreshold int    `yaml:"days_threshold"`
}

type DockerMonitor struct {
	Name          string `yaml:"name"`
	ContainerName string `yaml:"container_name"`
}

type EmailConfig struct {
	Enabled    bool     `yaml:"enabled"`
	SMTPHost   string   `yaml:"smtp_host"`
	SMTPPort   int      `yaml:"smtp_port"`
	SMTPUser   string   `yaml:"smtp_user"`
	SMTPPass   string   `yaml:"smtp_pass"`
	Recipients []string `yaml:"recipients"`
}

type SMSConfig struct {
	Enabled    bool     `yaml:"enabled"`
	APIKey     string   `yaml:"api_key"`
	Sender     string   `yaml:"sender"`
	Recipients []string `yaml:"recipients"`
}

type ChatConfig struct {
	Enabled    bool   `yaml:"enabled"`
	WebhookURL string `yaml:"webhook_url"`
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	var cfg Config
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	if cfg.IntervalSeconds == 0 {
		cfg.IntervalSeconds = 60
	}

	return &cfg, nil
}

func MustLoadConfig(path string) *Config {
	cfg, err := LoadConfig(path)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return cfg
}
