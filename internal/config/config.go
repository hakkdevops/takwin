package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// Config represents the complete build configuration
type Config struct {
	Project Project  `toml:"project"`
	Build   Build    `toml:"build"`
	Targets []Target `toml:"targets"`
}

// Project contains project metadata
type Project struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
}

// Build contains global build settings
type Build struct {
	Compiler     string   `toml:"compiler"`
	Optimization string   `toml:"optimization"`
	OutputDir    string   `toml:"output_dir"`
	IncludePaths []string `toml:"include_paths"`
	Libraries    []string `toml:"libraries"`
	LibraryPaths []string `toml:"library_paths"`
	CompileFlags []string `toml:"compile_flags"`
	LinkFlags    []string `toml:"link_flags"`
}

// Target represents a build target
type Target struct {
	Name         string   `toml:"name"`
	Type         string   `toml:"type"` // executable, static_library, shared_library
	Sources      []string `toml:"sources"`
	Output       string   `toml:"output"`
	IncludePaths []string `toml:"include_paths"`
	Libraries    []string `toml:"libraries"`
	LibraryPaths []string `toml:"library_paths"`
	CompileFlags []string `toml:"compile_flags"`
	LinkFlags    []string `toml:"link_flags"`
}

// Load reads and parses a TOML configuration file
func Load(filename string) (*Config, error) {
	if filename == "" {
		filename = "build.toml"
	}

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("configuration file '%s' not found", filename)
	}

	var config Config
	if _, err := toml.DecodeFile(filename, &config); err != nil {
		return nil, fmt.Errorf("failed to parse configuration file: %w", err)
	}

	// Validate configuration
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return &config, nil
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.Project.Name == "" {
		return fmt.Errorf("project name is required")
	}

	if c.Project.Version == "" {
		return fmt.Errorf("project version is required")
	}

	if len(c.Targets) == 0 {
		return fmt.Errorf("at least one target must be defined")
	}

	for i, target := range c.Targets {
		if target.Name == "" {
			return fmt.Errorf("target %d: name is required", i)
		}

		if target.Type == "" {
			return fmt.Errorf("target '%s': type is required", target.Name)
		}

		validTypes := []string{"executable", "static_library", "shared_library"}
		valid := false
		for _, validType := range validTypes {
			if target.Type == validType {
				valid = true
				break
			}
		}
		if !valid {
			return fmt.Errorf("target '%s': invalid type '%s', must be one of: %v",
				target.Name, target.Type, validTypes)
		}

		if len(target.Sources) == 0 {
			return fmt.Errorf("target '%s': at least one source file is required", target.Name)
		}
	}

	return nil
}

// GetDefaultTarget returns the first target or nil if no targets exist
func (c *Config) GetDefaultTarget() *Target {
	if len(c.Targets) == 0 {
		return nil
	}
	return &c.Targets[0]
}

// GetTarget returns a target by name or nil if not found
func (c *Config) GetTarget(name string) *Target {
	for _, target := range c.Targets {
		if target.Name == name {
			return &target
		}
	}
	return nil
}
