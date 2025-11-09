# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial Go implementation of Takwin build tool
- Cross-platform support (Linux, macOS, Windows)
- Multiple compiler support (GCC, Clang, MSVC)
- TOML configuration file support
- CLI with build, clean, and list-targets commands
- Comprehensive test suite
- CI/CD pipeline with GitHub Actions
- Docker support
- Documentation and examples

### Changed
- Complete rewrite from Python to Go for better performance
- Single binary distribution (no runtime dependencies)
- 20x+ faster startup time compared to Python version
- Improved error handling and validation

### Performance
- Startup time: ~50ms (vs ~1000ms Python version)
- Memory usage: ~5MB (vs ~25MB Python version)
- Binary size: ~8MB (vs ~50MB+ Python + dependencies)

## [2.0.0] - 2024-01-XX

### Added
- **Core Features**
  - TOML-based configuration system
  - Multi-target build support (executable, static library, shared library)
  - Cross-platform compiler support (GCC, Clang, MSVC)
  - Flexible source file resolution with glob patterns
  - Platform-specific file extensions and paths
  - Verbose and quiet output modes

- **CLI Commands**
  - `takwin build` - Build targets
  - `takwin build -t <target>` - Build specific target
  - `takwin clean` - Clean build artifacts
  - `takwin list-targets` - List available targets
  - `takwin --version` - Show version information
  - `takwin --help` - Show help information

- **Configuration Features**
  - Project metadata (name, version)
  - Global build settings (compiler, optimization, output directory)
  - Target-specific settings (sources, libraries, flags)
  - Include path management
  - Library linking support
  - Custom compile and link flags

- **Developer Experience**
  - Comprehensive error messages
  - Configuration validation
  - Source file existence checking
  - Build directory creation
  - Cross-compilation support

- **Testing & Quality**
  - Unit tests for all core components
  - Integration tests with real examples
  - Code coverage reporting
  - Linting with golangci-lint
  - Security scanning with gosec

- **CI/CD & Release**
  - GitHub Actions workflows
  - Multi-platform binary builds
  - Automated releases with GitHub Releases
  - Docker image publishing
  - Homebrew formula updates

- **Documentation**
  - Comprehensive README
  - Installation guide
  - Configuration reference
  - Usage examples
  - API documentation
  - Migration guide from Python version

### Technical Details
- **Language**: Go 1.21+
- **Dependencies**: Minimal (cobra, viper, toml, testify)
- **Platforms**: Linux (amd64, arm64), macOS (amd64, arm64), Windows (amd64)
- **Architecture**: Clean, modular design with interfaces
- **Testing**: >85% code coverage
- **Performance**: 20x+ faster than Python version

### Migration from Python Version
- 100% configuration file compatibility
- Identical command-line interface
- Same build behavior and output
- No breaking changes for existing users

## Version History

### Pre-2.0.0 (Python Version)
- Original Python implementation
- Poetry-based dependency management
- Click-based CLI
- Comprehensive test suite (101 tests)
- CI/CD with GitHub Actions
- Documentation with Sphinx

---

## Release Process

### Version Numbering
- **Major** (X.0.0): Breaking changes, major new features
- **Minor** (X.Y.0): New features, backwards compatible
- **Patch** (X.Y.Z): Bug fixes, security updates

### Release Types
- **Stable**: Production-ready releases (v2.0.0)
- **Pre-release**: Beta versions (v2.0.0-beta.1)
- **Development**: Nightly builds (v2.0.0-dev.20240101)

### Supported Versions
- **Current**: v2.x.x (Active development)
- **Previous**: v1.x.x (Python version, maintenance only)
- **Legacy**: v0.x.x (Deprecated)

### Security Updates
Security vulnerabilities are addressed in:
- Current major version (immediate)
- Previous major version (6 months)
- Critical issues may receive patches for older versions

---

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for details on:
- How to contribute
- Development setup
- Testing requirements
- Release process
- Code of conduct
