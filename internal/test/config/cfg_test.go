package config_test

import (
	"main/internal/config"
	"os"
	"strconv"
	"testing"
)

func TestConfig(t *testing.T) {
	os.Setenv("CONFIG_PATH", "/home/moxcelix/Documents/projects/cosmo-messenger/config.yaml")

	cfg, err := config.NewConfig()

	if err != nil {
		t.Fatalf("cannot load config: %v", err)
	}

	t.Run("YAML Config", func(t *testing.T) {
		tests := []struct {
			name     string
			value    string
			expected string
		}{
			{name: "Message.DeleteDuration", value: cfg.Policies.Message.DeleteDuration.String(), expected: "1h0m0s"},
			{name: "Message.EditDuration", value: cfg.Policies.Message.EditDuration.String(), expected: "15m0s"},
			{name: "Message.MaxLength", value: strconv.Itoa(cfg.Policies.Message.MaxLength), expected: "4000"},
			{name: "Message.MinLength", value: strconv.Itoa(cfg.Policies.Message.MinLength), expected: "1"},
			{name: "Chat.MaxChatNameLength", value: strconv.Itoa(cfg.Policies.Chat.MaxChatNameLength), expected: "100"},
			{name: "Chat.MaxGroupMembers", value: strconv.Itoa(cfg.Policies.Chat.MaxGroupMembers), expected: "100"},
			{name: "Chat.MinChatNameLength", value: strconv.Itoa(cfg.Policies.Chat.MinChatNameLength), expected: "1"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if tt.expected != tt.value {
					t.Errorf("value (%q), wanted (%q)", tt.value, tt.expected)
				}
			})
		}
	})
}
