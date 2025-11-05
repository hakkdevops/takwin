package compiler

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAdapter(t *testing.T) {
	tests := []struct {
		name     string
		compiler string
		expected string
	}{
		{"gcc", "gcc", "gcc"},
		{"clang", "clang", "clang"},
		{"msvc", "msvc", "msvc"},
		{"unknown", "unknown", "gcc"}, // defaults to gcc
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := NewAdapter(tt.compiler)
			assert.Equal(t, tt.expected, adapter.Name())
		})
	}
}

func TestGccAdapter(t *testing.T) {
	adapter := NewGccAdapter()

	assert.Equal(t, "gcc", adapter.Name())
	assert.True(t, adapter.SupportsLanguage("c"))
	assert.True(t, adapter.SupportsLanguage("c++"))
	assert.True(t, adapter.SupportsLanguage("cpp"))
	assert.False(t, adapter.SupportsLanguage("java"))
}

func TestGccBuildCompileCommand(t *testing.T) {
	adapter := NewGccAdapter()

	ctx := &Context{
		Sources:      []string{"main.cpp", "utils.cpp"},
		Output:       "myapp.exe",
		TargetType:   Executable,
		IncludePaths: []string{"include", "third_party"},
		Libraries:    []string{"pthread", "m"},
		LibraryPaths: []string{"lib", "/usr/lib"},
		CompileFlags: []string{"-Wall", "-std=c++17"},
		LinkFlags:    []string{"-static"},
		Optimization: "O2",
	}

	cmd := adapter.BuildCompileCommand(ctx)

	expected := []string{
		"g++",
		"-O2",
		"-Iinclude",
		"-Ithird_party",
		"-Wall",
		"-std=c++17",
		"main.cpp",
		"utils.cpp",
		"-o",
		"myapp.exe",
		"-Llib",
		"-L/usr/lib",
		"-lpthread",
		"-lm",
		"-static",
	}

	assert.Equal(t, expected, cmd)
}

func TestGccBuildSharedLibraryCommand(t *testing.T) {
	adapter := NewGccAdapter()

	ctx := &Context{
		Sources:    []string{"lib.cpp"},
		Output:     "libmylib.so",
		TargetType: SharedLibrary,
	}

	cmd := adapter.BuildCompileCommand(ctx)

	assert.Contains(t, cmd, "g++")
	assert.Contains(t, cmd, "-shared")
	if runtime.GOOS != "windows" {
		assert.Contains(t, cmd, "-fPIC")
	}
}

func TestGccBuildStaticLibraryCommand(t *testing.T) {
	adapter := NewGccAdapter()

	ctx := &Context{
		Sources: []string{"lib.cpp"},
		Output:  "libmylib.a",
	}

	cmd := adapter.BuildStaticLibraryCommand(ctx)

	expected := []string{
		"ar",
		"rcs",
		"libmylib.a",
		"lib.cpp.o",
	}

	assert.Equal(t, expected, cmd)
}

func TestClangAdapter(t *testing.T) {
	adapter := NewClangAdapter()

	assert.Equal(t, "clang", adapter.Name())

	ctx := &Context{
		Sources:    []string{"main.cpp"},
		Output:     "app",
		TargetType: Executable,
	}

	cmd := adapter.BuildCompileCommand(ctx)
	assert.Equal(t, "clang++", cmd[0])
}

func TestMsvcAdapter(t *testing.T) {
	adapter := NewMsvcAdapter()

	assert.Equal(t, "msvc", adapter.Name())
	assert.True(t, adapter.SupportsLanguage("c++"))

	ctx := &Context{
		Sources:      []string{"main.cpp"},
		Output:       "app.exe",
		TargetType:   Executable,
		IncludePaths: []string{"include"},
		Libraries:    []string{"kernel32"},
		Optimization: "O2",
	}

	cmd := adapter.BuildCompileCommand(ctx)

	assert.Equal(t, "cl.exe", cmd[0])
	assert.Contains(t, cmd, "/O2")
	assert.Contains(t, cmd, "/Iinclude")
	assert.Contains(t, cmd, "/Fe:app.exe")
	assert.Contains(t, cmd, "kernel32.lib")
}

func TestMsvcOptimizationMapping(t *testing.T) {
	adapter := NewMsvcAdapter()

	tests := []struct {
		input    string
		expected string
	}{
		{"O0", "/Od"},
		{"O1", "/O1"},
		{"O2", "/O2"},
		{"O3", "/Ox"},
		{"Os", "/Ox"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			ctx := &Context{
				Sources:      []string{"main.cpp"},
				Output:       "app.exe",
				Optimization: tt.input,
			}

			cmd := adapter.BuildCompileCommand(ctx)
			assert.Contains(t, cmd, tt.expected)
		})
	}
}
