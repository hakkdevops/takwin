package platform

import (
	"path/filepath"
	"runtime"
	"strings"
)

// Adapter interface for platform-specific operations
type Adapter interface {
	ExecutableExtension() string
	StaticLibraryExtension() string
	SharedLibraryExtension() string
	ObjectFileExtension() string
	PathSeparator() string
	AddExtension(filename, targetType string) string
}

// NewAdapter creates a platform adapter for the current OS
func NewAdapter() Adapter {
	switch runtime.GOOS {
	case "windows":
		return NewWindowsAdapter()
	case "darwin":
		return NewMacOSAdapter()
	default:
		return NewUnixAdapter()
	}
}

// WindowsAdapter implements platform-specific behavior for Windows
type WindowsAdapter struct{}

func NewWindowsAdapter() *WindowsAdapter {
	return &WindowsAdapter{}
}

func (w *WindowsAdapter) ExecutableExtension() string {
	return ".exe"
}

func (w *WindowsAdapter) StaticLibraryExtension() string {
	return ".lib"
}

func (w *WindowsAdapter) SharedLibraryExtension() string {
	return ".dll"
}

func (w *WindowsAdapter) ObjectFileExtension() string {
	return ".obj"
}

func (w *WindowsAdapter) PathSeparator() string {
	return "\\"
}

func (w *WindowsAdapter) AddExtension(filename, targetType string) string {
	switch targetType {
	case "executable":
		if !strings.HasSuffix(filename, w.ExecutableExtension()) {
			return filename + w.ExecutableExtension()
		}
	case "static_library":
		if !strings.HasSuffix(filename, w.StaticLibraryExtension()) {
			// Add lib prefix if not present
			base := filepath.Base(filename)
			if !strings.HasPrefix(base, "lib") {
				dir := filepath.Dir(filename)
				if dir == "." {
					filename = "lib" + filename
				} else {
					filename = filepath.Join(dir, "lib"+base)
				}
			}
			return filename + w.StaticLibraryExtension()
		}
	case "shared_library":
		if !strings.HasSuffix(filename, w.SharedLibraryExtension()) {
			return filename + w.SharedLibraryExtension()
		}
	}
	return filename
}

// UnixAdapter implements platform-specific behavior for Unix-like systems
type UnixAdapter struct{}

func NewUnixAdapter() *UnixAdapter {
	return &UnixAdapter{}
}

func (u *UnixAdapter) ExecutableExtension() string {
	return ""
}

func (u *UnixAdapter) StaticLibraryExtension() string {
	return ".a"
}

func (u *UnixAdapter) SharedLibraryExtension() string {
	return ".so"
}

func (u *UnixAdapter) ObjectFileExtension() string {
	return ".o"
}

func (u *UnixAdapter) PathSeparator() string {
	return "/"
}

func (u *UnixAdapter) AddExtension(filename, targetType string) string {
	switch targetType {
	case "executable":
		// No extension needed for Unix executables
		return filename
	case "static_library":
		if !strings.HasSuffix(filename, u.StaticLibraryExtension()) {
			// Add lib prefix if not present
			base := filepath.Base(filename)
			if !strings.HasPrefix(base, "lib") {
				dir := filepath.Dir(filename)
				if dir == "." {
					filename = "lib" + filename
				} else {
					filename = filepath.Join(dir, "lib"+base)
				}
			}
			return filename + u.StaticLibraryExtension()
		}
	case "shared_library":
		if !strings.HasSuffix(filename, u.SharedLibraryExtension()) {
			// Add lib prefix if not present
			base := filepath.Base(filename)
			if !strings.HasPrefix(base, "lib") {
				dir := filepath.Dir(filename)
				if dir == "." {
					filename = "lib" + filename
				} else {
					filename = filepath.Join(dir, "lib"+base)
				}
			}
			return filename + u.SharedLibraryExtension()
		}
	}
	return filename
}

// MacOSAdapter implements platform-specific behavior for macOS
type MacOSAdapter struct {
	*UnixAdapter
}

func NewMacOSAdapter() *MacOSAdapter {
	return &MacOSAdapter{
		UnixAdapter: NewUnixAdapter(),
	}
}

func (m *MacOSAdapter) SharedLibraryExtension() string {
	return ".dylib"
}
