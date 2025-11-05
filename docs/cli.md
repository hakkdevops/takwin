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
| `--config` | `-c` | Configuration file path | `build.toml` |
| `--verbose` | `-v` | Enable verbose output | `false` |
| `--help` | `-h` | Show help information | - |
| `--version` | | Show version information | - |

### Examples

```bash
# Use custom config file
takwin --config my-config.toml build

# Enable verbose output
takwin --verbose build

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
| `--target` | `-t` | Specific target to build | All targets |

#### Examples

```bash
# Build all targets
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
takwin build -t mylib -v -c debug.toml
```

#### Output

```bash
$ takwin build
Building target 'myapp' (executable)
Source files: [/path/to/src/main.cpp, /path/to/src/utils.cpp]
Output: build/bin/myapp
Running: g++ -O2 -Iinclude /path/to/src/main.cpp /path/to/src/utils.cpp -o build/bin/myapp
Build completed successfully
```

#### Verbose Output

```bash
$ takwin build -v
Using config file: /path/to/build.toml
Loading configuration...
Validating configuration...
Building target 'myapp' (executable)
Creating output directory: build/bin
Resolving source files...
Source files: [/path/to/src/main.cpp, /path/to/src/utils.cpp]
Output: build/bin/myapp
Compiler: gcc
Optimization: O2
Include paths: [include]
Libraries: [pthread]
Running: g++ -O2 -Iinclude /path/to/src/main.cpp /path/to/src/utils.cpp -o build/bin/myapp -lpthread
Command completed successfully
Build completed successfully
```

### `clean`

Remove build artifacts and output directories.

```bash
takwin clean [options]
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
takwin list-targets [options]
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
    Libraries: mylib
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

### Default Locations

Takwin looks for configuration files in this order:

1. File specified with `--config` flag
2. `build.toml` in current directory
3. `takwin.toml` in current directory

### Custom Configuration

```bash
# Use different config file
takwin --config release.toml build

# Use config from different directory
takwin --config ../configs/debug.toml build
```

## Exit Codes

| Code | Description |
|------|-------------|
| `0` | Success |
| `1` | General error |
| `2` | Configuration error |
| `3` | Build error |
| `4` | File not found |

## Environment Variables

### `TAKWIN_CONFIG`

Override default configuration file:

```bash
export TAKWIN_CONFIG=my-config.toml
takwin build  # Uses my-config.toml
```

### `TAKWIN_VERBOSE`

Enable verbose output by default:

```bash
export TAKWIN_VERBOSE=1
takwin build  # Runs with verbose output
```

### `TAKWIN_COMPILER`

Override default compiler:

```bash
export TAKWIN_COMPILER=clang
takwin build  # Uses clang instead of gcc
```

## Shell Completion

### Bash

```bash
# Generate completion script
takwin completion bash > /etc/bash_completion.d/takwin

# Or for current user
takwin completion bash > ~/.bash_completion.d/takwin
```

### Zsh

```bash
# Generate completion script
takwin completion zsh > "${fpath[1]}/_takwin"
```

### Fish

```bash
# Generate completion script
takwin completion fish > ~/.config/fish/completions/takwin.fish
```

### PowerShell

```powershell
# Generate completion script
takwin completion powershell > takwin.ps1
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
FROM ghcr.io/hakkim/takwin-go:latest
COPY . /workspace
WORKDIR /workspace
RUN takwin build
```

## Debugging

### Common Issues

**1. Configuration not found**
```bash
$ takwin build
Error: configuration file 'build.toml' not found

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
# Enable maximum verbosity
takwin build -v

# Check configuration parsing
takwin list-targets -v

# Verify file resolution
takwin build -v | grep "Source files"
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

### Reduced Output

```bash
# Quiet mode (minimal output)
takwin build 2>/dev/null

# Only show errors
takwin build 2>&1 | grep -i error
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
        }
    ]
}
```

---

For more information, see the [Configuration Reference](configuration.md) and [Examples](examples.md).