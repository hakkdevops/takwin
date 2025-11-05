# Takwin Go Setup Guide

## Prerequisites

### 1. Install Go

Download and install Go from https://golang.org/dl/

**Windows:**
- Download the Windows installer (.msi)
- Run the installer and follow the prompts
- Verify installation: `go version`

**Linux/macOS:**
```bash
# Using package manager (Ubuntu/Debian)
sudo apt update
sudo apt install golang-go

# Or download from official site
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

### 2. Verify Go Installation

```bash
go version
# Should output: go version go1.21.x ...
```

## Setup Project

### 1. Initialize Go Module

```bash
cd takwin-go
go mod tidy
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Build the Project

```bash
# Build for current platform
go build -o takwin main.go

# Or use Makefile
make build
```

### 4. Run Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
make test-coverage
```

### 5. Test with Examples

```bash
# Build and test simple example
make examples

# Or manually
./takwin build --config examples/simple/build.toml
./takwin build --config examples/complex/build.toml
```

## Development Workflow

### 1. Code Formatting

```bash
go fmt ./...
# Or: make fmt
```

### 2. Linting (optional)

```bash
# Install golangci-lint first
make dev-setup

# Run linting
make lint
```

### 3. Testing

```bash
# Run tests
make test

# Run with coverage
make test-coverage
```

### 4. Cross-compilation

```bash
# Build for all platforms
make cross-compile

# Manual cross-compilation
GOOS=windows GOARCH=amd64 go build -o takwin.exe main.go
GOOS=linux GOARCH=amd64 go build -o takwin main.go
GOOS=darwin GOARCH=amd64 go build -o takwin main.go
```

## Project Structure Explained

```
takwin-go/
├── main.go                 # Entry point - calls cmd.Execute()
├── go.mod                  # Go module definition
├── go.sum                  # Dependency checksums (auto-generated)
├── Makefile               # Build automation
├── README.md              # Project documentation
├── SETUP.md               # This setup guide
│
├── cmd/                   # CLI command definitions (Cobra)
│   ├── root.go           # Root command and global flags
│   ├── build.go          # Build command implementation
│   ├── clean.go          # Clean command implementation
│   └── list.go           # List targets command
│
├── internal/             # Internal packages (not importable by others)
│   ├── config/           # Configuration file handling
│   │   ├── config.go     # Config structs and loading logic
│   │   └── config_test.go # Config tests
│   │
│   ├── build/            # Build engine
│   │   └── engine.go     # Main build orchestration
│   │
│   ├── compiler/         # Compiler adapters
│   │   └── adapter.go    # GCC, Clang, MSVC implementations
│   │
│   ├── platform/         # Platform-specific code
│   │   └── adapter.go    # Windows, Linux, macOS differences
│   │
│   └── sources/          # Source file resolution
│       └── resolver.go   # Glob pattern and explicit file handling
│
└── examples/             # Example projects for testing
    ├── simple/           # Basic executable example
    │   ├── build.toml
    │   └── main.cpp
    └── complex/          # Library + executable example
        ├── build.toml
        ├── include/
        └── src/
```

## Key Go Concepts Used

### 1. Cobra CLI Framework
- Provides command-line interface structure
- Handles flags, arguments, and subcommands
- Used by kubectl, Hugo, and many other CLI tools

### 2. Viper Configuration
- Handles configuration file loading
- Supports TOML, JSON, YAML formats
- Environment variable integration

### 3. Interface-based Design
- Compiler adapters implement common interface
- Platform adapters implement common interface
- Easy to extend with new compilers/platforms

### 4. Internal Packages
- Code in `internal/` cannot be imported by external packages
- Enforces clean API boundaries
- Prevents accidental dependencies

## Migration Benefits

### Performance Improvements
- **Startup time**: Go binary starts instantly vs Python interpreter
- **Memory usage**: Lower memory footprint
- **Execution speed**: Compiled Go is much faster than interpreted Python

### Distribution Benefits
- **Single binary**: No Python runtime or dependencies needed
- **Cross-compilation**: Build for any platform from any platform
- **Easy deployment**: Just copy the binary

### Development Benefits
- **Static typing**: Catch errors at compile time
- **Built-in testing**: `go test` is part of the standard toolchain
- **Fast compilation**: Go compiles very quickly
- **Excellent tooling**: Built-in formatter, linter, profiler

## Next Steps

1. **Install Go** if not already installed
2. **Run setup commands** listed above
3. **Test with examples** to verify everything works
4. **Compare performance** with Python version
5. **Extend functionality** as needed

## Troubleshooting

### Go Not Found
- Ensure Go is installed and in PATH
- Restart terminal after installation
- Check `go version` works

### Module Issues
- Run `go mod tidy` to fix dependency issues
- Delete `go.sum` and run `go mod download` if needed

### Build Issues
- Ensure you're in the `takwin-go` directory
- Check that all source files are present
- Run `go mod verify` to check dependencies

### Cross-compilation Issues
- Some packages don't support all platforms
- Use build tags if needed: `//go:build windows`
- Check GOOS/GOARCH combinations are valid