# CLI Reference

Complete command-line interface reference for Takwin.

## Global Options

These options are available for all commands:

```bash
takwin [global options] <command> [command options]
```

### Global Flags

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--config` | | Configuration file path | `build.toml` |
| `--verbose` | `-v` | Enable verbose output | `false` |
| `--help` | `-h` | Show help information | - |
| `--version` | | Show version information | - |

### Examples

```bash
# Use custom config file
takwin --config my-config.toml build

# Enable verbose output
takwin --verbose build
takwin -v build

# Show version
takwin --version

# Show help
takwin --help
```

## Commands

### `build`

Build one or more targets defined in the configuration.

```bash
takwin build [options]
```

#### Options

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--target` | `-t` | Specific target to build | Default target |

#### Examples

```bash
# Build default target
takwin build

# Build specific target
takwin build --target myapp
takwin build -t myapp

# Build with verbose output
takwin build --verbose
takwin build -v

# Build using custom config
takwin build --config release.toml

# Combine options
takwin build -t mylib -v --config debug.toml
```

#### Output

```bash
$ takwin build
Building target 'myapp' (executable)
Source files: [/path/to/src/main.cpp, /path/to/src/utils.cpp]
Output: build/bin/myapp
Running: g++ /path/to/src/main.cpp /path/to/src/utils.cpp -o build/bin/myapp
Command execution completed (dry run)
```

### `clean`

Remove build artifacts and output directories.

```bash
takwin clean
```

#### Examples

```bash
# Clean build directory
takwin clean

# Clean with verbose output
takwin clean -v

# Clean using custom config
takwin clean --config my-config.toml
```

#### Output

```bash
$ takwin clean
Cleaned build directory: /path/to/project/build
```

### `list-targets`

Display all available build targets defined in the configuration.

```bash
takwin list-targets
```

#### Examples

```bash
# List all targets
takwin list-targets

# List targets with verbose output
takwin list-targets -v

# List targets from custom config
takwin list-targets --config my-config.toml
```

#### Output

```bash
$ takwin list-targets
Available targets (3):

  myapp
    Type: executable
    Sources: src/main.cpp, src/utils.cpp

  mylib
    Type: static_library
    Sources: src/lib/*.cpp
    Output: libmylib

  tests
    Type: executable
    Sources: tests/*.cpp
```

### `help`

Show help information for commands.

```bash
takwin help [command]
```

#### Examples

```bash
# Show general help
takwin help

# Show help for specific command
takwin help build
takwin help clean
takwin help list-targets
```

## Configuration File

### Default Location

Takwin looks for the configuration file in this order:

1. File specified with `--config` flag
2. `build.toml` in current directory

### Custom Configuration

```bash
# Use different config file
takwin --config release.toml build

# Use config from different directory
takwin --config ../configs/debug.toml build
```

## Advanced Usage

### Scripting

```bash
#!/bin/bash
set -e

# Build all targets
takwin build

# Check if specific target exists
if takwin list-targets | grep -q "tests"; then
    echo "Running tests..."
    ./build/bin/tests
fi

# Clean up
takwin clean
```

### CI/CD Integration

```yaml
# GitHub Actions
- name: Build with Takwin
  run: |
    takwin build
    takwin list-targets
    
# Check build artifacts
- name: Verify Build
  run: |
    test -f build/bin/myapp
    ./build/bin/myapp --version
```

### Docker Usage

```dockerfile
FROM ghcr.io/hakkdevops/takwin:latest
COPY . /workspace
WORKDIR /workspace
RUN takwin build
```

## Debugging

### Common Issues

**1. Configuration not found**
```bash
$ takwin build
Error: failed to load config: open build.toml: no such file or directory

# Solution: Create build.toml or specify path
takwin build --config path/to/config.toml
```

**2. Target not found**
```bash
$ takwin build -t nonexistent
Error: target 'nonexistent' not found

# Solution: List available targets
takwin list-targets
```

**3. Compiler not found**
```bash
$ takwin build
Error: compiler 'gcc' not found in PATH

# Solution: Install compiler or specify different one
# In build.toml:
[build]
compiler = "clang"  # or "msvc"
```

### Debug Mode

```bash
# Enable verbose output
takwin build -v

# Check configuration parsing
takwin list-targets -v
```

## Performance Tips

### Faster Builds

```bash
# Build specific target only
takwin build -t myapp

# Use parallel compilation (if supported by compiler)
# In build.toml:
[build]
compile_flags = ["-j4"]  # GCC parallel compilation
```

## Integration Examples

### Make Integration

```makefile
# Makefile
.PHONY: build clean test

build:
	takwin build

clean:
	takwin clean

test: build
	./build/bin/tests

install: build
	cp build/bin/myapp /usr/local/bin/
```

### CMake Integration

```cmake
# CMakeLists.txt
add_custom_target(takwin_build
    COMMAND takwin build
    WORKING_DIRECTORY ${CMAKE_SOURCE_DIR}
    COMMENT "Building with Takwin"
)
```

### VS Code Integration

```json
// .vscode/tasks.json
{
    "version": "2.0.0",
    "tasks": [
        {
            "label": "Takwin Build",
            "type": "shell",
            "command": "takwin",
            "args": ["build"],
            "group": {
                "kind": "build",
                "isDefault": true
            }
        },
        {
            "label": "Takwin Clean",
            "type": "shell",
            "command": "takwin",
            "args": ["clean"]
        }
    ]
}
```

## Platform Support

Takwin supports the following platforms:

- **Linux** (x64, ARM64)
- **macOS** (x64, ARM64/Apple Silicon)
- **Windows** (x64)

Supported compilers:

- **GCC** (default)
- **Clang**
- **MSVC** (Windows)

---

For more information, see the [Configuration Reference](configuration.md) and [Quick Start Guide](quickstart.md).
