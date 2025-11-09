# Takwin

[![CI Pipeline](https://github.com/hakkim/takwin/actions/workflows/ci.yml/badge.svg)](https://github.com/hakkim/takwin/actions/workflows/ci.yml)
[![Release](https://github.com/hakkim/takwin/actions/workflows/release.yml/badge.svg)](https://github.com/hakkim/takwin/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/hakkim/takwin)](https://goreportcard.com/report/github.com/hakkim/takwin)
[![codecov](https://codecov.io/gh/hakkim/takwin/branch/main/graph/badge.svg)](https://codecov.io/gh/hakkim/takwin)
[![GitHub release](https://img.shields.io/github/release/hakkim/takwin.svg)](https://github.com/hakkim/takwin/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A modern, lightning-fast build tool for C and C++ projects. Built with Go for enterprise-grade performance and reliability.

## Table of Contents

- [Why Takwin?](#-why-takwin)
- [Quick Start](#-quick-start)
- [Features](#️-features)
- [Documentation](#-documentation)
- [Configuration](#-configuration)
- [Configuration Reference](#configuration-reference)
- [Examples](#examples)
- [Performance Comparison](#-performance-comparison)
- [Development](#️-development)
- [Development Tasks](#development-tasks)
- [Enterprise Usage](#-enterprise-usage)
- [Contributing](#-contributing)
- [Project Status](#-project-status)
- [Roadmap](#️-roadmap)
- [Resources](#-resources)

## 🚀 Why Takwin?

**20x Faster Than Traditional Build Tools**
- **Instant startup**: ~50ms vs ~1000ms+ for Python-based tools
- **Single binary**: No runtime dependencies or complex installations
- **Cross-platform**: Native support for Linux, macOS, and Windows
- **Enterprise-ready**: Comprehensive testing, CI/CD, and security scanning

## ⚡ Quick Start

### Installation

```bash
# Linux/macOS
curl -L https://github.com/hakkim/takwin/releases/latest/download/takwin-linux-amd64.tar.gz | tar -xz
sudo mv takwin-linux-amd64 /usr/local/bin/takwin

# macOS (Homebrew)
brew tap hakkim/tap && brew install takwin

# Windows (PowerShell)
Invoke-WebRequest -Uri "https://github.com/hakkim/takwin/releases/latest/download/takwin-windows-amd64.zip" -OutFile "takwin.zip"
Expand-Archive -Path "takwin.zip" -DestinationPath "C:\Program Files\Takwin"

# Docker
docker pull ghcr.io/hakkim/takwin:latest
```

### Your First Build

```bash
# Create a new C++ project
mkdir hello-world && cd hello-world

# Create source file
cat > main.cpp << EOF
#include <iostream>
int main() {
    std::cout << "Hello, Takwin!" << std::endl;
    return 0;
}
EOF

# Create build configuration
cat > build.toml << EOF
[project]
name = "hello"
version = "1.0.0"

[[targets]]
name = "hello"
type = "executable"
sources = ["main.cpp"]
EOF

# Build and run
takwin build
./build/bin/hello
```

### Basic Commands

```bash
# Build all targets
takwin build

# Build specific target
takwin build -t my_target

# Build with verbose output
takwin build -v

# List available targets
takwin list-targets

# Clean build artifacts
takwin clean

# Use custom config file
takwin build -c my_config.toml

# Show version
takwin --version

# Show help
takwin --help
```

## 🏗️ Features

### Core Capabilities
- ✅ **Multi-target builds** - Executables, static libraries, shared libraries
- ✅ **Cross-platform compilers** - GCC, Clang, MSVC support
- ✅ **Flexible source resolution** - Glob patterns and explicit file lists
- ✅ **Dependency management** - Library linking and include paths
- ✅ **Optimization levels** - O0, O1, O2, O3, Os support
- ✅ **Custom build flags** - Compile and link flag customization

### Developer Experience
- ✅ **Simple TOML configuration** - Human-readable build files
- ✅ **Comprehensive validation** - Early error detection
- ✅ **Verbose output modes** - Detailed build information
- ✅ **Clean build management** - Artifact cleanup
- ✅ **Target listing** - Easy project overview

### Enterprise Features
- ✅ **Security scanning** - Automated vulnerability detection
- ✅ **Comprehensive testing** - >85% code coverage
- ✅ **CI/CD ready** - GitHub Actions workflows included
- ✅ **Docker support** - Containerized builds
- ✅ **Cross-compilation** - Build for any platform from any platform

## 📖 Documentation

| Resource | Description |
|----------|-------------|
| **[Installation Guide](docs/installation.md)** | Detailed installation instructions for all platforms |
| **[Quick Start Tutorial](docs/quickstart.md)** | Get up and running in 5 minutes |
| **[Configuration Reference](docs/configuration.md)** | Complete configuration options |
| **[CLI Reference](docs/cli.md)** | Command-line interface documentation |
| **[Examples](docs/examples.md)** | Real-world usage examples |
| **[API Documentation](https://pkg.go.dev/github.com/hakkim/takwin)** | Go package documentation |

## 🔧 Configuration

### Simple Project
```toml
[project]
name = "my_app"
version = "1.0.0"

[build]
compiler = "gcc"
optimization = "O2"

[[targets]]
name = "main"
type = "executable"
sources = ["src/*.cpp"]
include_paths = ["include"]
```

### Complex Multi-Target Project
```toml
[project]
name = "advanced_project"
version = "2.1.0"

[build]
compiler = "gcc"
optimization = "O2"
output_dir = "build"
include_paths = ["include", "third_party/include"]
compile_flags = ["-Wall", "-Wextra", "-std=c++17"]

# Static library
[[targets]]
name = "core"
type = "static_library"
sources = ["src/core/*.cpp"]
output = "libcore"

# Shared library
[[targets]]
name = "plugin"
type = "shared_library"
sources = ["src/plugin/*.cpp"]
libraries = ["core"]
library_paths = ["build/lib"]

# Main executable
[[targets]]
name = "app"
type = "executable"
sources = ["src/main.cpp"]
libraries = ["core", "pthread"]
library_paths = ["build/lib"]
compile_flags = ["-DVERSION=\"2.1.0\""]
```

## Configuration Reference

### Essential Configuration Options

**Project Section:**
```toml
[project]
name = "my_project"        # Project name (required)
version = "1.0.0"          # Project version (required)
```

**Build Section:**
```toml
[build]
compiler = "gcc"           # Compiler: gcc, clang, msvc
optimization = "O2"        # Optimization: O0, O1, O2, O3, Os
output_dir = "build"       # Output directory
include_paths = ["include"] # Include directories
libraries = ["pthread"]    # Libraries to link
compile_flags = ["-Wall"]  # Custom compile flags
link_flags = ["-static"]   # Custom link flags
```

**Targets Section:**
```toml
[[targets]]
name = "my_app"           # Target name (required)
type = "executable"       # Type: executable, static_library, shared_library
sources = ["src/*.cpp"]   # Source files (glob patterns supported)
output = "custom_name"    # Custom output name (optional)
include_paths = ["inc"]   # Target-specific includes
libraries = ["mylib"]     # Target-specific libraries
library_paths = ["lib"]   # Library search paths
```

> **📖 For complete configuration options, see the [Configuration Reference](docs/configuration.md)**

## Examples

### Simple Executable
```toml
[project]
name = "hello"
version = "1.0.0"

[[targets]]
name = "hello"
type = "executable"
sources = ["main.cpp"]
```

### Library with Executable
```toml
[project]
name = "math_project"
version = "1.0.0"

[build]
compiler = "gcc"
optimization = "O2"

# Static library
[[targets]]
name = "mathlib"
type = "static_library"
sources = ["src/math/*.cpp"]

# Executable using the library
[[targets]]
name = "calculator"
type = "executable"
sources = ["src/main.cpp"]
libraries = ["mathlib"]
library_paths = ["build/lib"]
```

### Cross-Platform Configuration
```toml
[project]
name = "cross_platform"
version = "1.0.0"

[build]
compiler = "gcc"
optimization = "O2"
include_paths = ["include"]

# Windows-specific settings
[build.windows]
libraries = ["ws2_32", "user32"]
compile_flags = ["-DWIN32"]

# Linux-specific settings  
[build.linux]
libraries = ["pthread", "dl"]
compile_flags = ["-DLINUX"]

[[targets]]
name = "app"
type = "executable"
sources = ["src/*.cpp"]
```

> **📚 More examples available in the [Examples Documentation](docs/examples.md)**

## 🚀 Performance Comparison

| Metric | Python Takwin | Go Takwin | Improvement |
|--------|---------------|-----------|-------------|
| **Startup Time** | ~1,000ms | ~50ms | **20x faster** |
| **Memory Usage** | ~25MB | ~5MB | **5x less** |
| **Binary Size** | ~50MB+ deps | ~8MB | **6x smaller** |
| **Build Speed** | Moderate | Fast | **2-5x faster** |
| **Distribution** | Complex | Single binary | **Simplified** |

## 🛠️ Development

### Prerequisites
- Go 1.20 or later
- GCC/Clang/MSVC (for building C++ projects)

### Building from Source
```bash
git clone https://github.com/hakkim/takwin.git
cd takwin

# Install dependencies
go mod download

# Run tests
go test ./...

# Build
go build -o takwin main.go

# Cross-compile
GOOS=linux GOARCH=amd64 go build -o takwin-linux main.go
GOOS=windows GOARCH=amd64 go build -o takwin.exe main.go
GOOS=darwin GOARCH=amd64 go build -o takwin-macos main.go
```

### Development Workflow
```bash
# Format code
go fmt ./...

# Lint code
golangci-lint run

# Run tests with coverage
go test -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run integration tests
cd examples/simple && ../../takwin build
cd examples/complex && ../../takwin build -t mathlib
```

### Development Tasks

Takwin includes convenient development workflows:

```bash
# Quick validation (build + test examples)
go test ./... && cd examples/simple && ../../takwin build

# Code quality checks
go fmt ./...              # Format code
go vet ./...              # Check for issues
golangci-lint run         # Comprehensive linting

# Testing
go test ./...                                    # Run all tests
go test -race ./...                             # Run with race detection
go test -coverprofile=coverage.out ./...       # Generate coverage
go tool cover -html=coverage.out               # View coverage report

# Building
go build -o takwin main.go                     # Build for current platform
go build -ldflags="-s -w" -o takwin main.go   # Optimized build

# Cross-compilation
GOOS=linux GOARCH=amd64 go build -o takwin-linux main.go
GOOS=windows GOARCH=amd64 go build -o takwin.exe main.go
GOOS=darwin GOARCH=amd64 go build -o takwin-macos main.go

# Example testing
cd examples/simple && ../../takwin build       # Test simple example
cd examples/complex && ../../takwin build -t mathlib  # Test complex example
```

### Project Structure

```
takwin/                  # Repository root
├── main.go                 # Entry point
├── cmd/                    # CLI commands (Cobra)
│   ├── root.go            # Root command and global flags
│   ├── build.go           # Build command implementation
│   ├── clean.go           # Clean command implementation
│   └── list.go            # List targets command
├── internal/              # Internal packages
│   ├── config/            # Configuration handling
│   │   ├── config.go      # TOML parsing and validation
│   │   └── config_test.go # Configuration tests
│   ├── build/             # Build engine
│   │   ├── engine.go      # Main build orchestration
│   │   └── engine_test.go # Build engine tests
│   ├── compiler/          # Compiler adapters
│   │   ├── adapter.go     # GCC, Clang, MSVC implementations
│   │   └── adapter_test.go # Compiler tests
│   ├── platform/          # Platform-specific code
│   │   └── adapter.go     # Windows, Linux, macOS support
│   └── sources/           # Source file resolution
│       ├── resolver.go    # Glob and explicit file handling
│       └── resolver_test.go # Source resolution tests
├── examples/              # Example projects
│   ├── simple/           # Basic executable example
│   └── complex/          # Multi-target library example
├── docs/                 # Documentation
├── .github/              # GitHub Actions workflows
├── go.mod               # Go module definition
├── go.sum               # Dependency checksums
├── Dockerfile           # Container build
├── CHANGELOG.md         # Version history
├── CONTRIBUTING.md      # Contribution guidelines
├── SECURITY.md          # Security policy
└── LICENSE              # MIT license
```

## 🏢 Enterprise Usage

### CI/CD Integration
```yaml
# GitHub Actions example
- name: Setup Takwin
  run: |
    curl -L https://github.com/hakkim/takwin/releases/latest/download/takwin-linux-amd64.tar.gz | tar -xz
    sudo mv takwin-linux-amd64 /usr/local/bin/takwin

- name: Build Project
  run: takwin build

- name: Run Tests
  run: ./build/bin/tests
```

### Docker Usage
```dockerfile
FROM ghcr.io/hakkim/takwin:latest AS builder
COPY . /workspace
WORKDIR /workspace
RUN takwin build

FROM alpine:latest
COPY --from=builder /workspace/build/bin/myapp /usr/local/bin/
CMD ["myapp"]
```

### Security & Compliance
- **Vulnerability scanning** with gosec and nancy
- **Dependency verification** with go mod verify
- **SBOM generation** for supply chain security
- **Signed releases** with checksums
- **Regular security updates**

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Quick Contribution Setup
```bash
# Fork and clone
git clone https://github.com/yourusername/takwin.git
cd takwin

# Create feature branch
git checkout -b feature/amazing-feature

# Make changes and test
go test ./...
golangci-lint run

# Commit and push
git commit -m "Add amazing feature"
git push origin feature/amazing-feature

# Create pull request
```

## 📊 Project Status

- **Stability**: Production Ready
- **Maintenance**: Actively Maintained
- **Support**: Community + Enterprise
- **License**: MIT
- **Go Version**: 1.20+

### Release Cycle
- **Major releases**: Every 6-12 months
- **Minor releases**: Monthly
- **Patch releases**: As needed
- **Security updates**: Immediate

## 🛣️ Roadmap

### Planned Features

**Performance & Scalability:**
- [ ] Incremental builds (only rebuild changed files)
- [ ] Dependency tracking and caching
- [ ] Parallel compilation support
- [ ] Build performance profiling
- [ ] Remote build caching

**Developer Experience:**
- [ ] IDE integration (VS Code, CLion, Vim)
- [ ] Interactive configuration wizard
- [ ] Build visualization and debugging
- [ ] Hot reload for development builds
- [ ] Build templates and scaffolding

**Enterprise Features:**
- [ ] Plugin system for extensibility
- [ ] Package management integration (Conan, vcpkg)
- [ ] Advanced cross-compilation support
- [ ] Build artifact signing
- [ ] Enterprise authentication and authorization

**Ecosystem Integration:**
- [ ] CMake project import/export
- [ ] Bazel compatibility layer
- [ ] Continuous integration templates
- [ ] Cloud build support (GitHub Actions, GitLab CI)
- [ ] Container-native builds

### Current Status

**✅ Completed (v2.0.0):**
- Core build functionality
- Cross-platform support (Linux, macOS, Windows)
- Multiple compiler support (GCC, Clang, MSVC)
- TOML configuration system
- Comprehensive testing (>85% coverage)
- CI/CD pipeline with security scanning
- Docker support and multi-arch builds
- Enterprise-grade documentation

**🚧 In Progress:**
- Advanced configuration validation
- Enhanced error messages and diagnostics
- Performance optimizations
- Extended platform support

**📋 Planned (v2.1.0):**
- Incremental build support
- IDE integration plugins
- Build caching system
- Enhanced cross-compilation

## 🔗 Links

- **[GitHub Repository](https://github.com/hakkim/takwin)**
- **[Release Notes](CHANGELOG.md)**
- **[Issue Tracker](https://github.com/hakkim/takwin/issues)**
- **[Discussions](https://github.com/hakkim/takwin/discussions)**
- **[Security Policy](SECURITY.md)**

## 📚 Resources

### Documentation & Guides
- **[Complete Documentation](docs/)** - Comprehensive guides and references
- **[Installation Guide](docs/installation.md)** - Platform-specific installation
- **[Quick Start Tutorial](docs/quickstart.md)** - Get started in 5 minutes
- **[Configuration Reference](docs/configuration.md)** - Complete configuration options
- **[CLI Reference](docs/cli.md)** - Command-line interface guide
- **[Examples](docs/examples.md)** - Real-world usage examples
- **[Contributing Guide](CONTRIBUTING.md)** - How to contribute
- **[Security Policy](SECURITY.md)** - Security and vulnerability reporting

### Development Resources
- **[API Documentation](https://pkg.go.dev/github.com/hakkim/takwin)** - Go package documentation
- **[Architecture Guide](docs/architecture.md)** - Internal design and structure
- **[Testing Guide](docs/testing.md)** - Testing strategies and practices
- **[Release Process](docs/release.md)** - How releases are made

### Community & Support
- **[GitHub Repository](https://github.com/hakkim/takwin)** - Source code and development
- **[Issue Tracker](https://github.com/hakkim/takwin/issues)** - Bug reports and feature requests
- **[Discussions](https://github.com/hakkim/takwin/discussions)** - Community discussions and Q&A
- **[Release Notes](CHANGELOG.md)** - Version history and changes

### Integration & Deployment
- **[Docker Hub](https://hub.docker.com/r/hakkim/takwin)** - Official Docker images
- **[GitHub Packages](https://github.com/hakkim/takwin/packages)** - Container registry
- **[GitHub Actions](https://github.com/hakkim/takwin/actions)** - CI/CD pipelines
- **[Homebrew Formula](https://github.com/hakkim/homebrew-tap)** - macOS package manager

### Comparison & Migration
- **[Python vs Go Comparison](MIGRATION_COMPARISON.md)** - Detailed feature comparison
- **[Migration Guide](docs/migration.md)** - Moving from Python version
- **[Performance Benchmarks](docs/benchmarks.md)** - Speed and efficiency metrics
- **[Compatibility Matrix](docs/compatibility.md)** - Platform and compiler support

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Made with ❤️ for the C/C++ community**

*Takwin: Where enterprise meets simplicity in C/C++ builds.*
