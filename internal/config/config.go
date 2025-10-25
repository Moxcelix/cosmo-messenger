package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type MessagePolicyConfig struct {
	EditDuration   time.Duration `yaml:"edit_duration"`
	DeleteDuration time.Duration `yaml:"delete_duration"`
	MaxLength      int           `yaml:"max_length"`
	MinLength      int           `yaml:"min_length"`
}

type ChatPolicyConfig struct {
	MaxGroupMembers   int `yaml:"max_group_members"`
	MaxChatNameLength int `yaml:"max_chat_name_length"`
	MinChatNameLength int `yaml:"min_chat_name_length"`
}

type PolicyConfig struct {
	Message MessagePolicyConfig `yaml:"message"`
	Chat    ChatPolicyConfig    `yaml:"chat"`
}

type Config struct {
	Policies PolicyConfig `yaml:"policies"`
}

func NewConfig() (*Config, error) {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "config.yaml"
	}

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
