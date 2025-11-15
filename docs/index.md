# Takwin

<div align="center">

**Modern, lightning-fast build tool for C and C++ projects**

[![CI Pipeline](https://github.com/hakkim/takwin/actions/workflows/ci.yml/badge.svg)](https://github.com/hakkim/takwin/actions/workflows/ci.yml)
[![Release](https://github.com/hakkim/takwin/actions/workflows/release.yml/badge.svg)](https://github.com/hakkim/takwin/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/hakkim/takwin)](https://goreportcard.com/report/github.com/hakkim/takwin)
[![GitHub release](https://img.shields.io/github/release/hakkim/takwin.svg)](https://github.com/hakkim/takwin/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

</div>

---

## Why Takwin?

**Lightning-Fast Build Tool**

- **Instant startup**: ~50ms startup time
- **Single binary**: No runtime dependencies or complex installations
- **Cross-platform**: Native support for Linux, macOS, and Windows
- **Enterprise-ready**: Comprehensive testing, CI/CD, and security scanning

## Quick Start

=== "Linux/macOS"

    ```bash
    # Download and install
    curl -L https://github.com/hakkim/takwin/releases/latest/download/takwin-linux-amd64.tar.gz | tar -xz
    sudo mv takwin-linux-amd64 /usr/local/bin/takwin
    
    # Verify installation
    takwin --version
    ```

=== "Windows"

    ```powershell
    # Download from releases
    Invoke-WebRequest -Uri "https://github.com/hakkim/takwin/releases/latest/download/takwin-windows-amd64.zip" -OutFile "takwin.zip"
    Expand-Archive -Path "takwin.zip" -DestinationPath "C:\Program Files\Takwin"
    
    # Add to PATH and verify
    takwin --version
    ```

=== "Docker"

    ```bash
    # Pull the image
    docker pull ghcr.io/hakkim/takwin:latest
    
    # Run in current directory
    docker run --rm -v $(pwd):/workspace -w /workspace ghcr.io/hakkim/takwin:latest build
    ```

## Your First Build

Create a simple C++ project:

```bash
# Create project directory
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

## Features

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

## Performance Metrics

| Metric | Value | Benefit |
|--------|-------|---------|
| **Startup Time** | ~50ms | Instant feedback |
| **Memory Usage** | ~5MB | Lightweight footprint |
| **Binary Size** | ~8MB | Easy distribution |
| **Build Speed** | Fast | Efficient compilation |
| **Distribution** | Single binary | No dependencies |

## Next Steps

<div class="grid cards" markdown>

-   :material-clock-fast:{ .lg .middle } __Quick Start__

    ---

    Get up and running in 5 minutes

    [:octicons-arrow-right-24: Quick Start Guide](quickstart.md)

-   :material-cog:{ .lg .middle } __Configuration__

    ---

    Learn how to configure your builds

    [:octicons-arrow-right-24: Configuration Reference](configuration.md)

-   :material-console:{ .lg .middle } __CLI Reference__

    ---

    Explore all available commands

    [:octicons-arrow-right-24: CLI Documentation](cli.md)

-   :material-code-braces:{ .lg .middle } __Examples__

    ---

    See real-world usage examples

    [:octicons-arrow-right-24: View Examples](examples.md)

</div>

## Community & Support

- **GitHub Issues**: [Report bugs or request features](https://github.com/hakkim/takwin/issues)
- **Discussions**: [Ask questions and share ideas](https://github.com/hakkim/takwin/discussions)
- **Contributing**: [Learn how to contribute](CONTRIBUTING.md)

## License

Takwin is released under the [MIT License](LICENSE.md).

---

<div align="center">

**Made with ❤️ for the C/C++ community**

*Takwin: Where enterprise meets simplicity in C/C++ builds.*

</div>
