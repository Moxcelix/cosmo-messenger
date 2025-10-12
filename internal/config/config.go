package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
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

func NewConfig() *Config {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Can't read the config.yaml file: ", err)
		return nil
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatal("Can't unmarshal the config.yaml file: ", err)
		return nil
	}

	return &config
}
