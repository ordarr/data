package core

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBuildConfig(t *testing.T) {
	tests := []struct {
		name        string
		env         map[string]string
		expectation Config
		err         error
	}{
		{
			name: "Loads Default Config from File",
			env: map[string]string{
				"CORE_ENVIRONMENT_FILE": ".env",
			},
			expectation: Config{
				Type:    "sqlite",
				Name:    "ordarr.db",
				User:    "ordarr",
				Pass:    "ordarr",
				Host:    "localhost",
				Port:    "5432",
				LogMode: true,
			},
			err: nil,
		},
		{
			name: "Overrides Default Config from Environment",
			env: map[string]string{
				"CORE_ENVIRONMENT_FILE":  ".env",
				"CORE_DATABASE_TYPE":     "postgres",
				"CORE_DATABASE_NAME":     "db_name",
				"CORE_DATABASE_USER":     "db_user",
				"CORE_DATABASE_PASS":     "db_pass",
				"CORE_DATABASE_HOST":     "db_host",
				"CORE_DATABASE_PORT":     "db_port",
				"CORE_DATABASE_LOG_MODE": "false",
			},
			expectation: Config{
				Type:    "postgres",
				Name:    "db_name",
				User:    "db_user",
				Pass:    "db_pass",
				Host:    "db_host",
				Port:    "db_port",
				LogMode: false,
			},
			err: nil,
		},
		{
			name: "Errors on Missing File",
			env: map[string]string{
				"CORE_DATABASE_ENVIRONMENT_FILE": ".envasdfasdf",
			},
			expectation: Config{},
			err:         errors.New("couldn't read .env file"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.env {
				_ = os.Setenv(k, v)
			}
			config, err := BuildConfig()
			if tt.err != nil {
				assert.Equal(t, tt.err, err)
			} else {
				assert.Equal(t, &tt.expectation, config)
			}
		})
	}
}
