package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hakkdevops/takwin/internal/compiler"
	"github.com/hakkdevops/takwin/internal/config"
	"github.com/hakkdevops/takwin/internal/platform"
	"github.com/hakkdevops/takwin/internal/sources"
)

// Target type constants
const (
	targetTypeExecutable    = "executable"
	targetTypeStaticLibrary = "static_library"
	targetTypeSharedLibrary = "shared_library"
)

// Engine handles the build process
type Engine struct {
	config   *config.Config
	compiler compiler.Adapter
	platform platform.Adapter
	resolver sources.Resolver
}

// NewEngine creates a new build engine
func NewEngine(cfg *config.Config) *Engine {
	compilerName := cfg.Build.Compiler
	if compilerName == "" {
		compilerName = "gcc" // default
	}

	return &Engine{
		config:   cfg,
		compiler: compiler.NewAdapter(compilerName),
		platform: platform.NewAdapter(),
		resolver: sources.NewGlobResolver(),
	}
}

// BuildDefault builds the default (first) target
func (e *Engine) BuildDefault() error {
	target := e.config.GetDefaultTarget()
	if target == nil {
		return fmt.Errorf("no targets defined")
	}

	return e.buildTarget(target)
}

// BuildTarget builds a specific target by name
func (e *Engine) BuildTarget(name string) error {
	target := e.config.GetTarget(name)
	if target == nil {
		return fmt.Errorf("target '%s' not found", name)
	}

	return e.buildTarget(target)
}

// buildTarget performs the actual build process
func (e *Engine) buildTarget(target *config.Target) error {
	fmt.Printf("Building target '%s' (%s)\n", target.Name, target.Type)

	// Resolve source files
	sourceFiles, err := e.resolver.Resolve(target.Sources, nil)
	if err != nil {
		return fmt.Errorf("failed to resolve sources: %w", err)
	}

	if len(sourceFiles) == 0 {
		return fmt.Errorf("no source files found for target '%s'", target.Name)
	}

	fmt.Printf("Source files: %v\n", sourceFiles)

	// Create output directory
	outputDir := e.getOutputDir(target)
	const dirPerm = 0755
	if err := os.MkdirAll(outputDir, dirPerm); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Get output path
	outputPath := e.getOutputPath(target)
	fmt.Printf("Output: %s\n", outputPath)

	// Build based on target type
	switch target.Type {
	case targetTypeExecutable:
		return e.buildExecutable(target, sourceFiles, outputPath)
	case targetTypeStaticLibrary:
		return e.buildStaticLibrary(target, sourceFiles, outputPath)
	case targetTypeSharedLibrary:
		return e.buildSharedLibrary(target, sourceFiles, outputPath)
	default:
		return fmt.Errorf("unsupported target type: %s", target.Type)
	}
}

// buildExecutable builds an executable target
func (e *Engine) buildExecutable(target *config.Target, sources []string, output string) error {
	ctx := e.buildContext(target, sources, output, compiler.Executable)
	cmd := e.compiler.BuildCompileCommand(ctx)
	return e.runCommand(cmd)
}

// buildStaticLibrary builds a static library target
func (e *Engine) buildStaticLibrary(target *config.Target, sources []string, output string) error {
	ctx := e.buildContext(target, sources, output, compiler.StaticLibrary)
	cmd := e.compiler.BuildStaticLibraryCommand(ctx)
	return e.runCommand(cmd)
}

// buildSharedLibrary builds a shared library target
func (e *Engine) buildSharedLibrary(target *config.Target, sources []string, output string) error {
	ctx := e.buildContext(target, sources, output, compiler.SharedLibrary)
	cmd := e.compiler.BuildCompileCommand(ctx)
	return e.runCommand(cmd)
}

// buildContext creates a compiler context with merged settings
func (e *Engine) buildContext(
	target *config.Target,
	sources []string,
	output string,
	targetType compiler.TargetType,
) *compiler.Context {
	return &compiler.Context{
		Sources:      sources,
		Output:       output,
		TargetType:   targetType,
		IncludePaths: e.mergeIncludePaths(target),
		Libraries:    e.mergeLibraries(target),
		LibraryPaths: e.mergeLibraryPaths(target),
		CompileFlags: e.mergeCompileFlags(target),
		LinkFlags:    e.mergeLinkFlags(target),
		Optimization: e.config.Build.Optimization,
	}
}

// Helper methods for merging global and target-specific settings
func (e *Engine) mergeIncludePaths(target *config.Target) []string {
	return append(e.config.Build.IncludePaths, target.IncludePaths...)
}

func (e *Engine) mergeLibraries(target *config.Target) []string {
	return append(e.config.Build.Libraries, target.Libraries...)
}

func (e *Engine) mergeLibraryPaths(target *config.Target) []string {
	return append(e.config.Build.LibraryPaths, target.LibraryPaths...)
}

func (e *Engine) mergeCompileFlags(target *config.Target) []string {
	return append(e.config.Build.CompileFlags, target.CompileFlags...)
}

func (e *Engine) mergeLinkFlags(target *config.Target) []string {
	return append(e.config.Build.LinkFlags, target.LinkFlags...)
}

// getOutputDir returns the output directory for a target
func (e *Engine) getOutputDir(target *config.Target) string {
	baseDir := e.config.Build.OutputDir
	if baseDir == "" {
		baseDir = "build"
	}

	switch target.Type {
	case targetTypeExecutable:
		return filepath.Join(baseDir, "bin")
	case targetTypeStaticLibrary, targetTypeSharedLibrary:
		return filepath.Join(baseDir, "lib")
	default:
		return baseDir
	}
}

// getOutputPath returns the full output path for a target
func (e *Engine) getOutputPath(target *config.Target) string {
	outputDir := e.getOutputDir(target)

	if target.Output != "" {
		filename := target.Output
		// Add platform-specific extension if not present
		filename = e.platform.AddExtension(filename, target.Type)
		return filepath.Join(outputDir, filename)
	}

	// Generate filename from target name
	filename := target.Name
	filename = e.platform.AddExtension(filename, target.Type)
	return filepath.Join(outputDir, filename)
}

// runCommand executes a shell command
func (e *Engine) runCommand(cmd []string) error {
	if len(cmd) == 0 {
		return fmt.Errorf("empty command")
	}

	fmt.Printf("Running: %s\n", strings.Join(cmd, " "))

	// For now, just print the command (actual execution would be added here)
	// This allows us to see what commands would be run
	fmt.Println("Command execution completed (dry run)")
	return nil
}
