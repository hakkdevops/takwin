# Frequently Asked Questions

## General

### What is Takwin?

Takwin is a modern, fast build tool for C and C++ projects. It uses simple TOML configuration files instead of complex Makefiles.

### Why use Takwin instead of Make/CMake?

- **Simpler**: TOML configuration vs complex Makefile syntax
- **Faster**: ~50ms startup time
- **Single binary**: No dependencies to install
- **Cross-platform**: Works on Linux, macOS, Windows
- **Modern**: Built with Go for reliability

### Is Takwin production-ready?

Takwin is currently in beta (v0.1.0-beta.1). It's suitable for:
- ✅ Personal projects
- ✅ Small to medium projects
- ✅ Evaluation and testing
- ⚠️ Production use (with caution)

### What license is Takwin under?

MIT License - free for commercial and personal use.

## Installation

### How do I install Takwin?

See the [Installation Guide](installation.md) for detailed instructions.

Quick install:
```bash
curl -L https://github.com/hakkim/takwin/releases/latest/download/takwin-linux-amd64.tar.gz | tar -xz
sudo mv takwin-linux-amd64 /usr/local/bin/takwin
```

### Do I need to install any dependencies?

No! Takwin is a single binary with no runtime dependencies.

### Can I use Takwin with Docker?

Yes! We provide official Docker images:
```bash
docker pull ghcr.io/hakkim/takwin:latest
```

## Usage

### How do I create a build configuration?

Create a `build.toml` file:
```toml
[project]
name = "myproject"
version = "1.0.0"

[[targets]]
name = "myapp"
type = "executable"
sources = ["src/*.cpp"]
```

### What compilers does Takwin support?

- GCC
- Clang
- MSVC (Windows)

### Can I use Takwin for C projects?

Yes! Takwin supports both C and C++ projects.

### How do I build multiple targets?

Define multiple `[[targets]]` sections:
```toml
[[targets]]
name = "mylib"
type = "static_library"
sources = ["lib/*.cpp"]

[[targets]]
name = "myapp"
type = "executable"
sources = ["src/*.cpp"]
libraries = ["mylib"]
```

### Can I use glob patterns for source files?

Yes! Takwin supports glob patterns:
```toml
sources = ["src/**/*.cpp"]  # Recursive
sources = ["src/*.cpp"]     # Single directory
sources = ["src/main.cpp", "src/utils.cpp"]  # Explicit files
```

## Configuration

### Where should I put my build.toml?

In your project root directory, next to your source code.

### Can I use a different config file name?

Yes, use the `--config` flag:
```bash
takwin build --config my-config.toml
```

### How do I set compiler flags?

```toml
[build]
compile_flags = ["-Wall", "-Wextra", "-std=c++17"]
link_flags = ["-static"]
```

### How do I add include paths?

```toml
[build]
include_paths = ["include", "third_party/include"]
```

### How do I link libraries?

```toml
[build]
libraries = ["pthread", "m"]
library_paths = ["/usr/local/lib"]
```

## Features

### Does Takwin support incremental builds?

Not yet. This is planned for v0.2.0.

### Can I run tests with Takwin?

Takwin focuses on building. Use your test framework separately:
```bash
takwin build
./build/bin/tests
```

### Does Takwin support cross-compilation?

Yes! Set the target platform:
```bash
GOOS=linux GOARCH=amd64 takwin build
```

### Can I use Takwin in CI/CD?

Yes! Takwin is designed for CI/CD:
```yaml
- name: Build with Takwin
  run: |
    curl -L https://github.com/hakkim/takwin/releases/latest/download/takwin-linux-amd64.tar.gz | tar -xz
    ./takwin-linux-amd64 build
```

## Troubleshooting

### Build fails with "compiler not found"

Install a C++ compiler:
- Linux: `sudo apt-get install build-essential`
- macOS: `xcode-select --install`
- Windows: Install MinGW or Visual Studio

### "No source files found"

Check your glob patterns and file paths in `build.toml`.

### Build succeeds but no output

Check the `build/` directory. Use `-v` for verbose output:
```bash
takwin build -v
```

### More troubleshooting help?

See the [Troubleshooting Guide](troubleshooting.md).

## Comparison

### Takwin vs Make

| Feature | Takwin | Make |
|---------|--------|------|
| Config | TOML | Makefile |
| Learning Curve | Easy | Steep |
| Cross-platform | ✅ | ⚠️ |
| Dependencies | None | make |

### Takwin vs CMake

| Feature | Takwin | CMake |
|---------|--------|-------|
| Config | TOML | CMakeLists.txt |
| Complexity | Simple | Complex |
| Features | Basic | Advanced |
| Best For | Small/Medium | Large projects |

### Takwin vs Meson

| Feature | Takwin | Meson |
|---------|--------|-------|
| Config | TOML | meson.build |
| Speed | Fast | Fast |
| Dependencies | None | Python |
| Maturity | Beta | Stable |

## Contributing

### How can I contribute?

See the [Contributing Guide](../contributing/index.md).

### Where do I report bugs?

[GitHub Issues](https://github.com/hakkim/takwin/issues)

### Can I request features?

Yes! Open a [feature request](https://github.com/hakkim/takwin/issues/new).

## Roadmap

### What's planned for future releases?

See [Version Strategy](../maintainer/version-strategy.md) for the roadmap.

**v0.2.0:**
- Incremental builds
- Build caching
- Dependency tracking

**v1.0.0:**
- Stable API
- Production ready
- Long-term support

### When will v1.0.0 be released?

When the API is stable and all core features are complete. Follow our [releases](https://github.com/hakkim/takwin/releases).

## Support

### Where can I get help?

- **Documentation**: You're reading it!
- **GitHub Discussions**: https://github.com/hakkim/takwin/discussions
- **GitHub Issues**: https://github.com/hakkim/takwin/issues

### Is there commercial support?

Not currently. Community support is available through GitHub.

## Didn't find your answer?

Ask in [GitHub Discussions](https://github.com/hakkim/takwin/discussions) or check the full [documentation](../index.md).
