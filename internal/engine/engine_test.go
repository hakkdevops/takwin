package engine

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/hakkdevops/takwin/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEngine(t *testing.T) {
	cfg := &config.Config{
		Build: config.Build{
			Compiler: "gcc",
		},
	}

	engine := NewEngine(cfg)
	assert.NotNil(t, engine)
	assert.Equal(t, cfg, engine.config)
	assert.NotNil(t, engine.compiler)
	assert.NotNil(t, engine.platform)
	assert.NotNil(t, engine.resolver)
}

func TestBuildDefault(t *testing.T) {
	// Create temporary directory for test
	tmpDir, err := os.MkdirTemp("", "takwin_test_*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create test source file
	srcFile := filepath.Join(tmpDir, "main.cpp")
	err = os.WriteFile(srcFile, []byte(`#include <iostream>
int main() { return 0; }`), 0600)
	require.NoError(t, err)

	// Change to temp directory
	oldWd, err := os.Getwd()
	require.NoError(t, err)
	defer func() {
		if chdirErr := os.Chdir(oldWd); chdirErr != nil {
			t.Logf("Failed to restore working directory: %v", chdirErr)
		}
	}()
	require.NoError(t, os.Chdir(tmpDir))

	cfg := &config.Config{
		Project: config.Project{
			Name:    "test",
			Version: "1.0.0",
		},
		Build: config.Build{
			Compiler: "gcc",
		},
		Targets: []config.Target{
			{
				Name:    "test",
				Type:    "executable",
				Sources: []string{"main.cpp"},
			},
		},
	}

	engine := NewEngine(cfg)
	err = engine.BuildDefault()
	assert.NoError(t, err)
}

func TestBuildTarget(t *testing.T) {
	cfg := &config.Config{
		Targets: []config.Target{
			{
				Name:    "myapp",
				Type:    "executable",
				Sources: []string{"main.cpp"},
			},
		},
	}

	engine := NewEngine(cfg)

	// Test existing target
	err := engine.BuildTarget("myapp")
	// Will fail due to missing source file, but should find the target
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no source files found")

	// Test non-existing target
	err = engine.BuildTarget("nonexistent")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "target 'nonexistent' not found")
}

func TestGetOutputPath(t *testing.T) {
	cfg := &config.Config{
		Build: config.Build{
			OutputDir: "build",
		},
	}

	engine := NewEngine(cfg)

	tests := []struct {
		name   string
		target config.Target
	}{
		{
			name: "executable with custom output",
			target: config.Target{
				Name:   "myapp",
				Type:   "executable",
				Output: "custom_name",
			},
		},
		{
			name: "executable without custom output",
			target: config.Target{
				Name: "myapp",
				Type: "executable",
			},
		},
		{
			name: "static library",
			target: config.Target{
				Name: "mylib",
				Type: "static_library",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := engine.getOutputPath(&tt.target)

			// Verify the path structure is correct
			assert.Contains(t, result, "build")

			// Verify correct subdirectory based on type
			if tt.target.Type == "executable" {
				assert.Contains(t, result, filepath.Join("build", "bin"))
			} else {
				assert.Contains(t, result, filepath.Join("build", "lib"))
			}

			// Verify the filename is present (platform-specific extension is added by platform adapter)
			if tt.target.Output != "" {
				assert.Contains(t, result, tt.target.Output)
			} else {
				assert.Contains(t, result, tt.target.Name)
			}
		})
	}
}

func TestMergeSettings(t *testing.T) {
	cfg := &config.Config{
		Build: config.Build{
			IncludePaths: []string{"global/include"},
			Libraries:    []string{"globallib"},
			CompileFlags: []string{"-Wall"},
		},
	}

	engine := NewEngine(cfg)

	target := &config.Target{
		IncludePaths: []string{"target/include"},
		Libraries:    []string{"targetlib"},
		CompileFlags: []string{"-Wextra"},
	}

	// Test merging
	includes := engine.mergeIncludePaths(target)
	assert.Equal(t, []string{"global/include", "target/include"}, includes)

	libs := engine.mergeLibraries(target)
	assert.Equal(t, []string{"globallib", "targetlib"}, libs)

	flags := engine.mergeCompileFlags(target)
	assert.Equal(t, []string{"-Wall", "-Wextra"}, flags)
}

func TestBuildStaticLibrary(t *testing.T) {
	cfg := &config.Config{
		Targets: []config.Target{
			{
				Name:    "mylib",
				Type:    "static_library",
				Sources: []string{"lib.cpp"},
			},
		},
	}

	engine := NewEngine(cfg)
	err := engine.BuildTarget("mylib")
	// Will fail due to missing source file, but tests the code path
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no source files found")
}

func TestBuildSharedLibrary(t *testing.T) {
	cfg := &config.Config{
		Targets: []config.Target{
			{
				Name:    "myshared",
				Type:    "shared_library",
				Sources: []string{"shared.cpp"},
			},
		},
	}

	engine := NewEngine(cfg)
	err := engine.BuildTarget("myshared")
	// Will fail due to missing source file, but tests the code path
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no source files found")
}

func TestGetOutputDir(t *testing.T) {
	cfg := &config.Config{
		Build: config.Build{
			OutputDir: "custom_build",
		},
	}

	engine := NewEngine(cfg)

	//nolint:govet // fieldalignment: minor optimization not worth readability loss in test
	tests := []struct {
		name     string
		target   config.Target
		expected string
	}{
		{
			name: "executable",
			target: config.Target{
				Type: "executable",
			},
			expected: filepath.Join("custom_build", "bin"),
		},
		{
			name: "static library",
			target: config.Target{
				Type: "static_library",
			},
			expected: filepath.Join("custom_build", "lib"),
		},
		{
			name: "shared library",
			target: config.Target{
				Type: "shared_library",
			},
			expected: filepath.Join("custom_build", "lib"),
		},
		{
			name: "unknown type",
			target: config.Target{
				Type: "unknown",
			},
			expected: "custom_build",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := engine.getOutputDir(&tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRunCommand(t *testing.T) {
	engine := NewEngine(&config.Config{})

	// Test with valid command
	err := engine.runCommand([]string{"gcc", "main.cpp", "-o", "main"})
	assert.NoError(t, err)

	// Test with empty command
	err = engine.runCommand([]string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "empty command")
}
