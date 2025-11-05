package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	// Create a temporary config file
	content := `
[project]
name = "test_project"
version = "1.0.0"

[build]
compiler = "gcc"
optimization = "O2"

[[targets]]
name = "main"
type = "executable"
sources = ["main.cpp"]
`

	tmpFile, err := os.CreateTemp("", "test_config_*.toml")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(content)
	require.NoError(t, err)
	tmpFile.Close()

	// Test loading the config
	cfg, err := Load(tmpFile.Name())
	require.NoError(t, err)

	assert.Equal(t, "test_project", cfg.Project.Name)
	assert.Equal(t, "1.0.0", cfg.Project.Version)
	assert.Equal(t, "gcc", cfg.Build.Compiler)
	assert.Equal(t, "O2", cfg.Build.Optimization)
	assert.Len(t, cfg.Targets, 1)
	assert.Equal(t, "main", cfg.Targets[0].Name)
	assert.Equal(t, "executable", cfg.Targets[0].Type)
	assert.Equal(t, []string{"main.cpp"}, cfg.Targets[0].Sources)
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid config",
			config: Config{
				Project: Project{Name: "test", Version: "1.0.0"},
				Targets: []Target{
					{Name: "main", Type: "executable", Sources: []string{"main.cpp"}},
				},
			},
			wantErr: false,
		},
		{
			name: "missing project name",
			config: Config{
				Project: Project{Version: "1.0.0"},
				Targets: []Target{
					{Name: "main", Type: "executable", Sources: []string{"main.cpp"}},
				},
			},
			wantErr: true,
			errMsg:  "project name is required",
		},
		{
			name: "missing project version",
			config: Config{
				Project: Project{Name: "test"},
				Targets: []Target{
					{Name: "main", Type: "executable", Sources: []string{"main.cpp"}},
				},
			},
			wantErr: true,
			errMsg:  "project version is required",
		},
		{
			name: "no targets",
			config: Config{
				Project: Project{Name: "test", Version: "1.0.0"},
				Targets: []Target{},
			},
			wantErr: true,
			errMsg:  "at least one target must be defined",
		},
		{
			name: "invalid target type",
			config: Config{
				Project: Project{Name: "test", Version: "1.0.0"},
				Targets: []Target{
					{Name: "main", Type: "invalid", Sources: []string{"main.cpp"}},
				},
			},
			wantErr: true,
			errMsg:  "invalid type 'invalid'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetTarget(t *testing.T) {
	cfg := &Config{
		Targets: []Target{
			{Name: "main", Type: "executable"},
			{Name: "lib", Type: "static_library"},
		},
	}

	// Test existing target
	target := cfg.GetTarget("main")
	require.NotNil(t, target)
	assert.Equal(t, "main", target.Name)
	assert.Equal(t, "executable", target.Type)

	// Test non-existing target
	target = cfg.GetTarget("nonexistent")
	assert.Nil(t, target)
}

func TestGetDefaultTarget(t *testing.T) {
	// Test with targets
	cfg := &Config{
		Targets: []Target{
			{Name: "main", Type: "executable"},
			{Name: "lib", Type: "static_library"},
		},
	}

	target := cfg.GetDefaultTarget()
	require.NotNil(t, target)
	assert.Equal(t, "main", target.Name)

	// Test with no targets
	cfg = &Config{Targets: []Target{}}
	target = cfg.GetDefaultTarget()
	assert.Nil(t, target)
}
