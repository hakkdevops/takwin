package sources

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGlobResolver(t *testing.T) {
	// Create temporary directory with test files
	tmpDir, err := os.MkdirTemp("", "takwin_test_*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create test files
	files := []string{
		"main.cpp",
		"utils.cpp",
		"test.cpp",
		"header.h",
		"readme.txt",
	}

	for _, file := range files {
		path := filepath.Join(tmpDir, file)
		writeErr := os.WriteFile(path, []byte("test content"), 0600)
		require.NoError(t, writeErr)
	}

	// Change to temp directory for relative path testing
	oldWd, err := os.Getwd()
	require.NoError(t, err)
	defer func() {
		if chdirErr := os.Chdir(oldWd); chdirErr != nil {
			t.Logf("Failed to restore working directory: %v", chdirErr)
		}
	}()
	require.NoError(t, os.Chdir(tmpDir))

	resolver := NewGlobResolver()

	t.Run("glob pattern matching", func(t *testing.T) {
		patterns := []string{"*.cpp"}
		result, err := resolver.Resolve(patterns, nil)
		require.NoError(t, err)

		assert.Len(t, result, 3)
		// Results should be absolute paths
		for _, path := range result {
			assert.True(t, filepath.IsAbs(path))
			assert.True(t, filepath.Ext(path) == ".cpp")
		}
	})

	t.Run("multiple patterns", func(t *testing.T) {
		patterns := []string{"*.cpp", "*.h"}
		result, err := resolver.Resolve(patterns, nil)
		require.NoError(t, err)

		assert.Len(t, result, 4) // 3 cpp + 1 h
	})

	t.Run("exclude patterns", func(t *testing.T) {
		patterns := []string{"*.cpp"}
		excludePatterns := []string{"test.cpp"}
		result, err := resolver.Resolve(patterns, excludePatterns)
		require.NoError(t, err)

		assert.Len(t, result, 2) // main.cpp and utils.cpp
		for _, path := range result {
			assert.NotContains(t, path, "test.cpp")
		}
	})

	t.Run("explicit file", func(t *testing.T) {
		patterns := []string{"main.cpp"}
		result, err := resolver.Resolve(patterns, nil)
		require.NoError(t, err)

		assert.Len(t, result, 1)
		assert.Contains(t, result[0], "main.cpp")
	})

	t.Run("no matches", func(t *testing.T) {
		patterns := []string{"*.java"}
		result, err := resolver.Resolve(patterns, nil)
		require.NoError(t, err)

		assert.Len(t, result, 0)
	})

	t.Run("duplicate removal", func(t *testing.T) {
		patterns := []string{"main.cpp", "*.cpp"}
		result, err := resolver.Resolve(patterns, nil)
		require.NoError(t, err)

		// Should not have duplicates
		seen := make(map[string]bool)
		for _, path := range result {
			assert.False(t, seen[path], "Duplicate path found: %s", path)
			seen[path] = true
		}
	})
}

func TestExplicitResolver(t *testing.T) {
	// Create temporary directory with test files
	tmpDir, err := os.MkdirTemp("", "takwin_test_*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create test file
	testFile := filepath.Join(tmpDir, "main.cpp")
	err = os.WriteFile(testFile, []byte("test content"), 0600)
	require.NoError(t, err)

	resolver := NewExplicitResolver()

	t.Run("explicit files", func(t *testing.T) {
		patterns := []string{testFile}
		result, err := resolver.Resolve(patterns, nil)
		require.NoError(t, err)

		assert.Len(t, result, 1)
		assert.True(t, filepath.IsAbs(result[0]))
		assert.Contains(t, result[0], "main.cpp")
	})

	t.Run("multiple explicit files", func(t *testing.T) {
		// Create another file
		testFile2 := filepath.Join(tmpDir, "utils.cpp")
		err := os.WriteFile(testFile2, []byte("test content"), 0600)
		require.NoError(t, err)

		patterns := []string{testFile, testFile2}
		result, err := resolver.Resolve(patterns, nil)
		require.NoError(t, err)

		assert.Len(t, result, 2)
	})

	t.Run("duplicate removal", func(t *testing.T) {
		patterns := []string{testFile, testFile}
		result, err := resolver.Resolve(patterns, nil)
		require.NoError(t, err)

		assert.Len(t, result, 1)
	})
}

func TestSmartResolver(t *testing.T) {
	// Create temporary directory with test files
	tmpDir, err := os.MkdirTemp("", "takwin_test_*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create test files
	files := []string{"main.cpp", "utils.cpp", "test.cpp"}
	for _, file := range files {
		path := filepath.Join(tmpDir, file)
		writeErr := os.WriteFile(path, []byte("test content"), 0600)
		require.NoError(t, writeErr)
	}

	// Change to temp directory
	oldWd, err := os.Getwd()
	require.NoError(t, err)
	defer func() {
		if chdirErr := os.Chdir(oldWd); chdirErr != nil {
			t.Logf("Failed to restore working directory: %v", chdirErr)
		}
	}()
	require.NoError(t, os.Chdir(tmpDir))

	resolver := NewSmartResolver()

	t.Run("mixed patterns and explicit files", func(t *testing.T) {
		patterns := []string{"*.cpp", "main.cpp"} // mix of glob and explicit
		result, err := resolver.Resolve(patterns, nil)
		require.NoError(t, err)

		assert.Len(t, result, 3) // Should not have duplicates
		for _, path := range result {
			assert.True(t, filepath.IsAbs(path))
		}
	})

	t.Run("only glob patterns", func(t *testing.T) {
		patterns := []string{"*.cpp"}
		result, err := resolver.Resolve(patterns, nil)
		require.NoError(t, err)

		assert.Len(t, result, 3)
	})

	t.Run("only explicit files", func(t *testing.T) {
		patterns := []string{"main.cpp", "utils.cpp"}
		result, err := resolver.Resolve(patterns, nil)
		require.NoError(t, err)

		assert.Len(t, result, 2)
	})
}
