package compiler

import (
	"runtime"
)

// TargetType represents the type of build target
type TargetType int

const (
	Executable TargetType = iota
	StaticLibrary
	SharedLibrary
)

// Context contains all information needed for compilation
type Context struct {
	Sources      []string
	Output       string
	TargetType   TargetType
	IncludePaths []string
	Libraries    []string
	LibraryPaths []string
	CompileFlags []string
	LinkFlags    []string
	Optimization string
	Language     string
}

// Adapter interface for different compilers
type Adapter interface {
	Name() string
	BuildCompileCommand(ctx *Context) []string
	BuildStaticLibraryCommand(ctx *Context) []string
	SupportsLanguage(lang string) bool
}

// NewAdapter creates a compiler adapter based on the compiler name
func NewAdapter(name string) Adapter {
	switch name {
	case "gcc":
		return NewGccAdapter()
	case "clang":
		return NewClangAdapter()
	case "msvc":
		return NewMsvcAdapter()
	default:
		// Default to GCC
		return NewGccAdapter()
	}
}

// GccAdapter implements the Adapter interface for GCC
type GccAdapter struct{}

func NewGccAdapter() *GccAdapter {
	return &GccAdapter{}
}

func (g *GccAdapter) Name() string {
	return "gcc"
}

func (g *GccAdapter) SupportsLanguage(lang string) bool {
	return lang == "c" || lang == "c++" || lang == "cpp"
}

func (g *GccAdapter) BuildCompileCommand(ctx *Context) []string {
	cmd := []string{"g++"}

	// Add optimization
	if ctx.Optimization != "" {
		cmd = append(cmd, "-"+ctx.Optimization)
	}

	// Add include paths
	for _, path := range ctx.IncludePaths {
		cmd = append(cmd, "-I"+path)
	}

	// Add compile flags
	cmd = append(cmd, ctx.CompileFlags...)

	// Add sources
	cmd = append(cmd, ctx.Sources...)

	// Add output
	cmd = append(cmd, "-o", ctx.Output)

	// Add library paths
	for _, path := range ctx.LibraryPaths {
		cmd = append(cmd, "-L"+path)
	}

	// Add libraries
	for _, lib := range ctx.Libraries {
		cmd = append(cmd, "-l"+lib)
	}

	// Add link flags
	cmd = append(cmd, ctx.LinkFlags...)

	// Add shared library flag if needed
	if ctx.TargetType == SharedLibrary {
		cmd = append(cmd, "-shared")
		if runtime.GOOS != "windows" {
			cmd = append(cmd, "-fPIC")
		}
	}

	return cmd
}

func (g *GccAdapter) BuildStaticLibraryCommand(ctx *Context) []string {
	// For static libraries, we need to compile objects first, then archive
	// For simplicity, we'll use a single command approach
	cmd := []string{"ar", "rcs", ctx.Output}

	// Add object files (we'd need to compile these first in a real implementation)
	for _, source := range ctx.Sources {
		objFile := source + ".o"
		cmd = append(cmd, objFile)
	}

	return cmd
}

// ClangAdapter implements the Adapter interface for Clang
type ClangAdapter struct {
	*GccAdapter // Clang is mostly compatible with GCC
}

func NewClangAdapter() *ClangAdapter {
	return &ClangAdapter{
		GccAdapter: NewGccAdapter(),
	}
}

func (c *ClangAdapter) Name() string {
	return "clang"
}

func (c *ClangAdapter) BuildCompileCommand(ctx *Context) []string {
	cmd := c.GccAdapter.BuildCompileCommand(ctx)
	// Replace gcc with clang++
	if len(cmd) > 0 && cmd[0] == "g++" {
		cmd[0] = "clang++"
	}
	return cmd
}

// MsvcAdapter implements the Adapter interface for MSVC
type MsvcAdapter struct{}

func NewMsvcAdapter() *MsvcAdapter {
	return &MsvcAdapter{}
}

func (m *MsvcAdapter) Name() string {
	return "msvc"
}

func (m *MsvcAdapter) SupportsLanguage(lang string) bool {
	return lang == "c" || lang == "c++" || lang == "cpp"
}

func (m *MsvcAdapter) BuildCompileCommand(ctx *Context) []string {
	cmd := []string{"cl.exe"}

	// Add optimization
	switch ctx.Optimization {
	case "O0":
		cmd = append(cmd, "/Od")
	case "O1":
		cmd = append(cmd, "/O1")
	case "O2":
		cmd = append(cmd, "/O2")
	case "O3", "Os":
		cmd = append(cmd, "/Ox")
	}

	// Add include paths
	for _, path := range ctx.IncludePaths {
		cmd = append(cmd, "/I"+path)
	}

	// Add compile flags
	cmd = append(cmd, ctx.CompileFlags...)

	// Add sources
	cmd = append(cmd, ctx.Sources...)

	// Add output
	cmd = append(cmd, "/Fe:"+ctx.Output)

	// Add library paths
	for _, path := range ctx.LibraryPaths {
		cmd = append(cmd, "/LIBPATH:"+path)
	}

	// Add libraries
	for _, lib := range ctx.Libraries {
		cmd = append(cmd, lib+".lib")
	}

	// Add link flags
	cmd = append(cmd, ctx.LinkFlags...)

	// Add shared library flag if needed
	if ctx.TargetType == SharedLibrary {
		cmd = append(cmd, "/LD")
	}

	return cmd
}

func (m *MsvcAdapter) BuildStaticLibraryCommand(ctx *Context) []string {
	cmd := []string{"lib.exe"}

	// Add output
	cmd = append(cmd, "/OUT:"+ctx.Output)

	// Add object files (we'd need to compile these first in a real implementation)
	for _, source := range ctx.Sources {
		objFile := source + ".obj"
		cmd = append(cmd, objFile)
	}

	return cmd
}
