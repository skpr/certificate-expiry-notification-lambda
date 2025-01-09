package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("testdata")
	assert.NoError(t, err)
	assert.Equal(t, []string{"http://example.com/slack-hook1", "http://example.com/slack-hook2"}, config.SlackWebhookURL)
}

func TestValidate(t *testing.T) {
	var tests = []struct {
		name   string
		config Config
		fails  bool
	}{
		{
			name:  "Missing Config",
			fails: true,
		},
		{
			name: "All the Config",
			config: Config{
				SlackWebhookURL: []string{"http://example.com/slack-webhook"},
			},
			fails: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.config.Validate()
			if len(ans) > 0 != tt.fails {
				t.Errorf("got %s, want %v", ans, tt.fails)
			}
		})
	}
}
