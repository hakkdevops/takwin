package sources

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Resolver interface for resolving source file patterns
type Resolver interface {
	Resolve(patterns, excludePatterns []string) ([]string, error)
}

// GlobResolver resolves source files using glob patterns
type GlobResolver struct{}

func NewGlobResolver() *GlobResolver {
	return &GlobResolver{}
}

func (g *GlobResolver) Resolve(patterns, excludePatterns []string) ([]string, error) {
	var allFiles []string
	seen := make(map[string]bool)

	for _, pattern := range patterns {
		files, err := filepath.Glob(pattern)
		if err != nil {
			return nil, fmt.Errorf("invalid glob pattern '%s': %w", pattern, err)
		}

		// If no files found with glob, try treating as explicit file
		if len(files) == 0 {
			// Check if the pattern exists as a file
			if _, err := os.Stat(pattern); err == nil {
				files = []string{pattern}
			}
		}

		for _, file := range files {
			// Convert to absolute path for consistency
			absFile, err := filepath.Abs(file)
			if err != nil {
				continue
			}

			// Skip if already seen (remove duplicates)
			if seen[absFile] {
				continue
			}

			// Check if file should be excluded
			excluded := false
			for _, excludePattern := range excludePatterns {
				matched, err := filepath.Match(excludePattern, filepath.Base(file))
				if err == nil && matched {
					excluded = true
					break
				}
			}

			if !excluded {
				allFiles = append(allFiles, absFile)
				seen[absFile] = true
			}
		}
	}

	return allFiles, nil
}

// ExplicitResolver resolves explicit file paths (no globbing)
type ExplicitResolver struct{}

func NewExplicitResolver() *ExplicitResolver {
	return &ExplicitResolver{}
}

func (e *ExplicitResolver) Resolve(patterns, excludePatterns []string) ([]string, error) {
	var allFiles []string
	seen := make(map[string]bool)

	for _, file := range patterns {
		// Convert to absolute path
		absFile, err := filepath.Abs(file)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve path '%s': %w", file, err)
		}

		// Skip if already seen
		if seen[absFile] {
			continue
		}

		// For explicit resolver, we don't apply exclude patterns
		// as the user explicitly specified these files
		allFiles = append(allFiles, absFile)
		seen[absFile] = true
	}

	return allFiles, nil
}

// SmartResolver automatically chooses between glob and explicit resolution
type SmartResolver struct {
	globResolver     *GlobResolver
	explicitResolver *ExplicitResolver
}

func NewSmartResolver() *SmartResolver {
	return &SmartResolver{
		globResolver:     NewGlobResolver(),
		explicitResolver: NewExplicitResolver(),
	}
}

func (s *SmartResolver) Resolve(patterns, excludePatterns []string) ([]string, error) {
	var globPatterns []string
	var explicitFiles []string

	// Separate glob patterns from explicit files
	for _, pattern := range patterns {
		if strings.ContainsAny(pattern, "*?[]") {
			globPatterns = append(globPatterns, pattern)
		} else {
			explicitFiles = append(explicitFiles, pattern)
		}
	}

	var allFiles []string
	seen := make(map[string]bool)

	// Resolve glob patterns
	if len(globPatterns) > 0 {
		globFiles, err := s.globResolver.Resolve(globPatterns, excludePatterns)
		if err != nil {
			return nil, err
		}
		for _, file := range globFiles {
			if !seen[file] {
				allFiles = append(allFiles, file)
				seen[file] = true
			}
		}
	}

	// Resolve explicit files
	if len(explicitFiles) > 0 {
		explicitResolved, err := s.explicitResolver.Resolve(explicitFiles, nil)
		if err != nil {
			return nil, err
		}
		for _, file := range explicitResolved {
			if !seen[file] {
				allFiles = append(allFiles, file)
				seen[file] = true
			}
		}
	}

	return allFiles, nil
}
