package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name   string
		env    string
		region string
		want   *APIConfig
	}{
		{
			name:   "local config",
			env:    "local",
			region: "local",
			want: &APIConfig{
				Port:   "8080",
				Env:    "local",
				Region: "local",
			},
		},
		{
			name:   "dev config",
			env:    "dev",
			region: "us-east-1",
			want: &APIConfig{
				Port:   "3000",
				Env:    "dev",
				Region: "us-east-1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set the environment variables in test environment
			t.Setenv("ENV", tt.env)
			t.Setenv("REGION", tt.region)

			// Load the config
			if err := LoadConfig("."); err != nil {
				t.Fatalf("error loading config: %v", err)
			}

			// Check the config
			assert.Equal(t, Config.Port, tt.want.Port)
			assert.Equal(t, Config.Env, tt.want.Env)
			assert.Equal(t, Config.Region, tt.want.Region)

		})
	}

}
