package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config represents the application configuration
type Config struct {
	// Server configuration
	Server struct {
		Port string
		Env  string
	}

	// Mysql configuration
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		DSN      string
	}

	// Redis configuration
	Redis struct {
		Host     string
		Port     string
		Password string
		DB       int
	}

	// Kafka configuration
	Kafka struct {
		Brokers []string
		Topics  struct {
			Task   string
			Result string
		}
	}

	// COS configuration
	COS struct {
		SecretID  string
		SecretKey string
		Region    string
		Buckets   struct {
			Code string
			Logs string
		}
	}

	// Kubernetes configuration
	K8s struct {
		Namespace      string
		AgentImage     string
		ServiceAccount string
	}

	// OpenAI configuration
	OpenAI struct {
		APIKey string
	}

	// GitHub configuration
	GitHub struct {
		Token string
	}
}

// LoadFromYAML 加载 YAML 配置文件
func LoadFromYAML(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func buildDSN(user, password, host, port, dbName string) string {
	return user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true"
}
