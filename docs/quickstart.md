# Quick Start Guide

Get up and running with Takwin in under 5 minutes!

## Prerequisites

- **C++ Compiler**: GCC, Clang, or MSVC
- **Operating System**: Linux, macOS, or Windows

## Installation

### Option 1: Download Binary (Recommended)

```bash
# Linux (x64)
curl -L https://github.com/hakkim/takwin-go/releases/latest/download/takwin-linux-amd64.tar.gz | tar -xz
sudo mv takwin-linux-amd64 /usr/local/bin/takwin

# macOS (x64)
curl -L https://github.com/hakkim/takwin-go/releases/latest/download/takwin-darwin-amd64.tar.gz | tar -xz
sudo mv takwin-darwin-amd64 /usr/local/bin/takwin

# macOS (Homebrew)
brew tap hakkim/tap && brew install takwin

# Windows (PowerShell)
Invoke-WebRequest -Uri "https://github.com/hakkim/takwin-go/releases/latest/download/takwin-windows-amd64.zip" -OutFile "takwin.zip"
Expand-Archive -Path "takwin.zip" -DestinationPath "C:\Program Files\Takwin"
```

### Option 2: Build from Source

```bash
git clone https://github.com/hakkim/takwin-go.git
cd takwin-go
go build -o takwin main.go
sudo mv takwin /usr/local/bin/  # Linux/macOS
```

## Verify Installation

```bash
takwin --version
# Should output: takwin version 2.0.0
```

## Your First Project

### 1. Create Project Directory

```bash
mkdir hello-takwin && cd hello-takwin
```

### 2. Create Source File

```bash
cat > main.cpp << EOF
#include <iostream>

int main() {
    std::cout << "Hello, Takwin!" << std::endl;
    return 0;
}
EOF
```

### 3. Create Build Configuration

```bash
cat > build.toml << EOF
[project]
name = "hello"
version = "1.0.0"

[[targets]]
name = "hello"
type = "executable"
sources = ["main.cpp"]
EOF
```

### 4. Build and Run

```bash
# Build the project
takwin build

# Run the executable
./build/bin/hello          # Linux/macOS
# or
build\bin\hello.exe        # Windows
```

**Output:**
```
Hello, Takwin!
```

## Common Commands

```bash
# Build all targets
takwin build

# Build specific target
takwin build -t hello

# List available targets
takwin list-targets

# Clean build artifacts
takwin clean

# Verbose output
takwin build -v

# Help
takwin --help
```

## Next Steps

### Multi-File Project

Create a more complex project:

```bash
mkdir math-project && cd math-project

# Create header file
mkdir include
cat > include/math_utils.h << EOF
#pragma once

int add(int a, int b);
int multiply(int a, int b);
EOF

# Create source files
mkdir src
cat > src/math_utils.cpp << EOF
#include "math_utils.h"

int add(int a, int b) {
    return a + b;
}

int multiply(int a, int b) {
    return a * b;
}
EOF

cat > src/main.cpp << EOF
#include <iostream>
#include "math_utils.h"

int main() {
    std::cout << "5 + 3 = " << add(5, 3) << std::endl;
    std::cout << "5 * 3 = " << multiply(5, 3) << std::endl;
    return 0;
}
EOF

# Create build configuration
cat > build.toml << EOF
[project]
name = "math_project"
version = "1.0.0"

[build]
compiler = "gcc"
optimization = "O2"
include_paths = ["include"]

[[targets]]
name = "math_app"
type = "executable"
sources = ["src/*.cpp"]
EOF

# Build and run
takwin build
./build/bin/math_app
```

### Library Project

Create a project with both library and executable:

```bash
mkdir library-project && cd library-project

# Create the same structure as above, then:
cat > build.toml << EOF
[project]
name = "library_project"
version = "1.0.0"

[build]
compiler = "gcc"
optimization = "O2"
include_paths = ["include"]

# Static library target
[[targets]]
name = "mathlib"
type = "static_library"
sources = ["src/math_utils.cpp"]
output = "math"

# Executable target using the library
[[targets]]
name = "app"
type = "executable"
sources = ["src/main.cpp"]
libraries = ["math"]
library_paths = ["build/lib"]
EOF

# Build library first, then executable
takwin build -t mathlib
takwin build -t app
```

## Configuration Options

### Basic Configuration

```toml
[project]
name = "my_project"
version = "1.0.0"

[build]
compiler = "gcc"           # gcc, clang, msvc
optimization = "O2"        # O0, O1, O2, O3, Os
output_dir = "build"

[[targets]]
name = "my_app"
type = "executable"        # executable, static_library, shared_library
sources = ["src/*.cpp"]    # Supports glob patterns
```

### Advanced Configuration

```toml
[project]
name = "advanced_project"
version = "2.0.0"

[build]
compiler = "gcc"
optimization = "O2"
output_dir = "build"
include_paths = ["include", "third_party/include"]
libraries = ["pthread", "m"]
compile_flags = ["-Wall", "-Wextra", "-std=c++17"]
link_flags = ["-static"]

[[targets]]
name = "my_app"
type = "executable"
sources = ["src/main.cpp", "src/utils.cpp"]
include_paths = ["app/include"]  # Target-specific includes
libraries = ["boost_system"]     # Target-specific libraries
compile_flags = ["-DAPP_VERSION=\"2.0.0\""]
```

## Troubleshooting

### Common Issues

**1. Command not found**
```bash
# Check if takwin is in PATH
which takwin

# Add to PATH if needed
export PATH=$PATH:/usr/local/bin
```

**2. Compiler not found**
```bash
# Install GCC (Ubuntu/Debian)
sudo apt-get install build-essential

# Install GCC (macOS)
xcode-select --install

# Install GCC (Windows)
# Download and install MinGW-w64 or use Visual Studio
```

**3. Build fails**
```bash
# Check verbose output
takwin build -v

# Verify source files exist
ls -la src/

# Check configuration syntax
takwin list-targets
```

## What's Next?

1. **[Configuration Reference](configuration.md)** - Learn all configuration options
2. **[Examples](examples.md)** - See real-world usage examples
3. **[CLI Reference](cli.md)** - Master the command-line interface
4. **[Contributing](../CONTRIBUTING.md)** - Help improve Takwin

## Getting Help

- **Documentation**: Browse the [docs](.) directory
- **Issues**: Report bugs on [GitHub Issues](https://github.com/hakkim/takwin-go/issues)
- **Discussions**: Ask questions in [GitHub Discussions](https://github.com/hakkim/takwin-go/discussions)

---

**Congratulations!** You've successfully built your first project with Takwin. 🎉