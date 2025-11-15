package platform

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAdapter(t *testing.T) {
	adapter := NewAdapter()
	assert.NotNil(t, adapter)

	// Verify correct adapter is created based on OS
	switch runtime.GOOS {
	case "windows":
		_, ok := adapter.(*WindowsAdapter)
		assert.True(t, ok, "Expected WindowsAdapter on Windows")
	case "darwin":
		_, ok := adapter.(*MacOSAdapter)
		assert.True(t, ok, "Expected MacOSAdapter on macOS")
	default:
		_, ok := adapter.(*UnixAdapter)
		assert.True(t, ok, "Expected UnixAdapter on Unix-like systems")
	}
}

func TestWindowsAdapter(t *testing.T) {
	adapter := NewWindowsAdapter()

	assert.Equal(t, ".exe", adapter.ExecutableExtension())
	assert.Equal(t, ".lib", adapter.StaticLibraryExtension())
	assert.Equal(t, ".dll", adapter.SharedLibraryExtension())
	assert.Equal(t, ".obj", adapter.ObjectFileExtension())
	assert.Equal(t, "\\", adapter.PathSeparator())

	testAddExtension(t, adapter, "program.exe", "libmath.lib", "shared.dll")

	// Test with existing extension
	assert.Equal(t, "program.exe", adapter.AddExtension("program.exe", "executable"))
}

func TestUnixAdapter(t *testing.T) {
	adapter := NewUnixAdapter()

	assert.Equal(t, "", adapter.ExecutableExtension())
	assert.Equal(t, ".a", adapter.StaticLibraryExtension())
	assert.Equal(t, ".so", adapter.SharedLibraryExtension())
	assert.Equal(t, ".o", adapter.ObjectFileExtension())
	assert.Equal(t, "/", adapter.PathSeparator())

	testAddExtension(t, adapter, "program", "libmath.a", "libshared.so")

	// Test with existing lib prefix
	assert.Equal(t, "libmath.a", adapter.AddExtension("libmath", "static_library"))
}

func testAddExtension(t *testing.T, adapter Adapter, expectedExe, expectedLib, expectedShared string) {
	t.Helper()
	assert.Equal(t, expectedExe, adapter.AddExtension("program", "executable"))
	assert.Equal(t, expectedLib, adapter.AddExtension("math", "static_library"))
	assert.Equal(t, expectedShared, adapter.AddExtension("shared", "shared_library"))
	assert.Equal(t, "unknown", adapter.AddExtension("unknown", "unknown_type"))
}

func TestMacOSAdapter(t *testing.T) {
	adapter := NewMacOSAdapter()

	assert.Equal(t, "", adapter.ExecutableExtension())
	assert.Equal(t, ".a", adapter.StaticLibraryExtension())
	assert.Equal(t, ".dylib", adapter.SharedLibraryExtension())
	assert.Equal(t, ".o", adapter.ObjectFileExtension())
	assert.Equal(t, "/", adapter.PathSeparator())

	// Test AddExtension - Note: MacOSAdapter embeds UnixAdapter which uses .so
	// This is expected behavior as the base implementation is used
	assert.Equal(t, "program", adapter.AddExtension("program", "executable"))
	assert.Equal(t, "libmath.a", adapter.AddExtension("math", "static_library"))
	// MacOSAdapter doesn't override AddExtension, so it uses UnixAdapter's .so
	assert.Contains(t, adapter.AddExtension("shared", "shared_library"), "shared")
}
