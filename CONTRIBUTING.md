# Contributing to Takwin

Thank you for your interest in contributing to Takwin! This document provides guidelines and information for contributors.

## 🚀 Quick Start

### Development Setup

1. **Fork and Clone**
   ```bash
   git clone https://github.com/yourusername/takwin-go.git
   cd takwin-go
   ```

2. **Install Dependencies**
   ```bash
   go mod download
   ```

3. **Install Development Tools**
   ```bash
   # Install golangci-lint
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
   
   # Install other tools
   go install github.com/securecodewarrior/github-action-gosec@latest
   ```

4. **Verify Setup**
   ```bash
   go test ./...
   golangci-lint run
   ```

## 🛠️ Development Workflow

### 1. Create Feature Branch
```bash
git checkout -b feature/your-feature-name
# or
git checkout -b fix/issue-number
```

### 2. Make Changes
- Write clean, well-documented code
- Follow Go conventions and idioms
- Add tests for new functionality
- Update documentation as needed

### 3. Test Your Changes
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -race -coverprofile=coverage.out ./...

# Run linting
golangci-lint run

# Test with examples
cd examples/simple && ../../takwin build
cd examples/complex && ../../takwin build
```

### 4. Commit Changes
```bash
# Stage changes
git add .

# Commit with descriptive message
git commit -m "feat: add support for custom compiler flags"

# Push to your fork
git push origin feature/your-feature-name
```

### 5. Create Pull Request
- Open a PR against the `main` branch
- Fill out the PR template
- Link any related issues
- Wait for review and address feedback

## 📝 Code Guidelines

### Go Style Guide
- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use [gofmt](https://golang.org/cmd/gofmt/) for formatting
- Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use meaningful variable and function names
- Write clear, concise comments

### Code Structure
```go
// Package comment
package main

import (
    // Standard library imports first
    "fmt"
    "os"
    
    // Third-party imports
    "github.com/spf13/cobra"
    
    // Local imports
    "github.com/hakkim/takwin-go/internal/config"
)

// Exported functions have comments
func BuildProject(cfg *config.Config) error {
    // Implementation
}

// Private functions can have comments too
func validateConfig(cfg *config.Config) error {
    // Implementation
}
```

### Testing Guidelines
- Write tests for all new functionality
- Use table-driven tests where appropriate
- Mock external dependencies
- Aim for >85% code coverage
- Include both positive and negative test cases

```go
func TestBuildProject(t *testing.T) {
    tests := []struct {
        name    string
        config  *config.Config
        wantErr bool
    }{
        {
            name: "valid config",
            config: &config.Config{
                Project: config.Project{Name: "test"},
                Targets: []config.Target{{Name: "main", Type: "executable"}},
            },
            wantErr: false,
        },
        // More test cases...
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := BuildProject(tt.config)
            if (err != nil) != tt.wantErr {
                t.Errorf("BuildProject() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

## 🐛 Bug Reports

### Before Reporting
1. Check existing issues
2. Try the latest version
3. Reproduce with minimal example
4. Check documentation

### Bug Report Template
```markdown
**Describe the bug**
A clear description of what the bug is.

**To Reproduce**
Steps to reproduce the behavior:
1. Create config file with...
2. Run command...
3. See error

**Expected behavior**
What you expected to happen.

**Environment:**
- OS: [e.g. Ubuntu 20.04]
- Go version: [e.g. 1.21.0]
- Takwin version: [e.g. 2.0.0]
- Compiler: [e.g. GCC 9.4.0]

**Additional context**
Any other context about the problem.
```

## 💡 Feature Requests

### Before Requesting
1. Check existing issues and discussions
2. Consider if it fits the project scope
3. Think about implementation complexity
4. Consider backwards compatibility

### Feature Request Template
```markdown
**Is your feature request related to a problem?**
A clear description of what the problem is.

**Describe the solution you'd like**
A clear description of what you want to happen.

**Describe alternatives you've considered**
Other solutions you've considered.

**Additional context**
Any other context or screenshots.
```

## 🔄 Pull Request Process

### PR Requirements
- [ ] Tests pass (`go test ./...`)
- [ ] Linting passes (`golangci-lint run`)
- [ ] Documentation updated
- [ ] CHANGELOG.md updated (for significant changes)
- [ ] Examples work with changes
- [ ] No breaking changes (or clearly documented)

### PR Template
```markdown
## Description
Brief description of changes.

## Type of Change
- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Documentation update

## Testing
- [ ] Unit tests added/updated
- [ ] Integration tests pass
- [ ] Manual testing completed

## Checklist
- [ ] Code follows style guidelines
- [ ] Self-review completed
- [ ] Documentation updated
- [ ] Tests added for new functionality
```

### Review Process
1. **Automated Checks**: CI/CD pipeline runs
2. **Code Review**: Maintainer reviews code
3. **Testing**: Manual testing if needed
4. **Approval**: Maintainer approves changes
5. **Merge**: Changes merged to main branch

## 📚 Documentation

### Types of Documentation
- **Code Comments**: Explain complex logic
- **README**: Project overview and quick start
- **API Docs**: Generated from code comments
- **User Guides**: Step-by-step instructions
- **Examples**: Working code samples

### Documentation Standards
- Use clear, concise language
- Include code examples
- Keep examples up-to-date
- Test all code examples
- Use consistent formatting

## 🏗️ Architecture

### Project Structure
```
takwin-go/
├── main.go                 # Entry point
├── cmd/                    # CLI commands (Cobra)
│   ├── root.go            # Root command
│   ├── build.go           # Build command
│   ├── clean.go           # Clean command
│   └── list.go            # List command
├── internal/              # Internal packages
│   ├── config/            # Configuration handling
│   ├── build/             # Build engine
│   ├── compiler/          # Compiler adapters
│   ├── platform/          # Platform-specific code
│   └── sources/           # Source file resolution
├── examples/              # Example projects
├── docs/                  # Documentation
└── .github/               # GitHub workflows
```

### Design Principles
- **Simplicity**: Keep interfaces simple and focused
- **Modularity**: Separate concerns into distinct packages
- **Testability**: Design for easy testing
- **Performance**: Optimize for speed and memory usage
- **Compatibility**: Maintain backwards compatibility

## 🔒 Security

### Security Guidelines
- Never commit secrets or credentials
- Validate all user inputs
- Use secure defaults
- Follow Go security best practices
- Report security issues privately

### Security Review Process
1. **Automated Scanning**: gosec, nancy
2. **Dependency Checking**: go mod verify
3. **Manual Review**: Security-focused code review
4. **Testing**: Security-specific tests

## 🎯 Release Process

### Version Numbering
- **Major** (X.0.0): Breaking changes
- **Minor** (X.Y.0): New features, backwards compatible
- **Patch** (X.Y.Z): Bug fixes, security updates

### Release Checklist
- [ ] All tests pass
- [ ] Documentation updated
- [ ] CHANGELOG.md updated
- [ ] Version bumped
- [ ] Tag created
- [ ] Release notes written
- [ ] Binaries built and tested
- [ ] Release published

## 🤝 Community

### Communication Channels
- **GitHub Issues**: Bug reports and feature requests
- **GitHub Discussions**: General questions and ideas
- **Pull Requests**: Code contributions
- **Email**: security@takwin.dev (security issues only)

### Code of Conduct
We follow the [Contributor Covenant Code of Conduct](https://www.contributor-covenant.org/version/2/1/code_of_conduct/).

### Recognition
Contributors are recognized in:
- CONTRIBUTORS.md file
- Release notes
- GitHub contributors page
- Special thanks in documentation

## 📞 Getting Help

### For Contributors
- Read this contributing guide
- Check existing issues and PRs
- Ask questions in GitHub Discussions
- Join community discussions

### For Users
- Check documentation first
- Search existing issues
- Create new issue with template
- Provide minimal reproduction case

---

Thank you for contributing to Takwin! Your efforts help make C/C++ development better for everyone. 🚀