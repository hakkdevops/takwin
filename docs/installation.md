# Installation Guide

## Quick Install

### Linux/macOS

```bash
# Download and install latest version
curl -L https://github.com/hakkim/takwin/releases/latest/download/takwin-linux-amd64.tar.gz | tar -xz
sudo mv takwin-linux-amd64 /usr/local/bin/takwin

# Verify installation
takwin --version
```

### Windows

1. Download `takwin-windows-amd64.zip` from [releases](https://github.com/hakkim/takwin/releases)
2. Extract the ZIP file
3. Add the directory to your PATH
4. Verify: `takwin --version`

### macOS (Homebrew)

```bash
# Add tap (once)
brew tap hakkim/tap

# Install
brew install takwin

# Verify
takwin --version
```

## Platform-Specific Instructions

### Linux

#### Ubuntu/Debian
```bash
# Download
wget https://github.com/hakkim/takwin/releases/latest/download/takwin-linux-amd64.tar.gz

# Extract and install
tar -xzf takwin-linux-amd64.tar.gz
sudo mv takwin-linux-amd64 /usr/local/bin/takwin
sudo chmod +x /usr/local/bin/takwin

# Verify
takwin --version
```

#### CentOS/RHEL/Fedora
```bash
# Download
curl -L -o takwin-linux-amd64.tar.gz https://github.com/hakkim/takwin/releases/latest/download/takwin-linux-amd64.tar.gz

# Extract and install
tar -xzf takwin-linux-amd64.tar.gz
sudo mv takwin-linux-amd64 /usr/local/bin/takwin
sudo chmod +x /usr/local/bin/takwin
```

#### Arch Linux
```bash
# Using AUR helper (yay)
yay -S takwin-bin

# Or manual installation
curl -L https://github.com/hakkim/takwin/releases/latest/download/takwin-linux-amd64.tar.gz | tar -xz
sudo mv takwin-linux-amd64 /usr/local/bin/takwin
```

### macOS

#### Intel Macs
```bash
curl -L https://github.com/hakkim/takwin/releases/latest/download/takwin-darwin-amd64.tar.gz | tar -xz
sudo mv takwin-darwin-amd64 /usr/local/bin/takwin
```

#### Apple Silicon (M1/M2)
```bash
curl -L https://github.com/hakkim/takwin/releases/latest/download/takwin-darwin-arm64.tar.gz | tar -xz
sudo mv takwin-darwin-arm64 /usr/local/bin/takwin
```

### Windows

#### PowerShell (Recommended)
```powershell
# Download
Invoke-WebRequest -Uri "https://github.com/hakkim/takwin/releases/latest/download/takwin-windows-amd64.zip" -OutFile "takwin.zip"

# Extract
Expand-Archive -Path "takwin.zip" -DestinationPath "C:\Program Files\Takwin"

# Add to PATH (requires admin)
$env:PATH += ";C:\Program Files\Takwin"
[Environment]::SetEnvironmentVariable("PATH", $env:PATH, [EnvironmentVariableTarget]::Machine)
```

#### Chocolatey
```powershell
# Install Chocolatey package (if available)
choco install takwin
```

#### Scoop
```powershell
# Add bucket
scoop bucket add hakkim https://github.com/hakkim/scoop-bucket

# Install
scoop install takwin
```

## Docker

### Run with Docker
```bash
# Pull image
docker pull ghcr.io/hakkim/takwin:latest

# Run in current directory
docker run --rm -v $(pwd):/workspace -w /workspace ghcr.io/hakkim/takwin:latest build

# Create alias for convenience
alias takwin='docker run --rm -v $(pwd):/workspace -w /workspace ghcr.io/hakkim/takwin:latest'
```

### Dockerfile for CI/CD
```dockerfile
FROM ghcr.io/hakkim/takwin:latest
COPY . /workspace
WORKDIR /workspace
RUN takwin build
```

## Building from Source

### Prerequisites
- Go 1.20 or later
- Git

### Build Steps
```bash
# Clone repository
git clone https://github.com/hakkim/takwin.git
cd takwin

# Build
go build -o takwin main.go

# Install (optional)
sudo mv takwin /usr/local/bin/

# Or install with go install
go install github.com/hakkim/takwin@latest
```

### Development Build
```bash
# Clone and build
git clone https://github.com/hakkim/takwin.git
cd takwin

# Install dependencies
go mod download

# Run tests
go test ./...

# Build with debug info
go build -o takwin main.go

# Run directly
go run main.go --help
```

## Verification

### Check Installation
```bash
# Version
takwin --version

# Help
takwin --help

# List commands
takwin help
```

### Test with Example
```bash
# Create test project
mkdir test-project && cd test-project

# Create simple C++ file
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

# Build
takwin build

# Check output
ls build/bin/
```

## Troubleshooting

### Common Issues

#### Command Not Found
```bash
# Check if binary is in PATH
which takwin

# Check PATH
echo $PATH

# Add to PATH (Linux/macOS)
export PATH=$PATH:/usr/local/bin

# Add to PATH (Windows)
set PATH=%PATH%;C:\Program Files\Takwin
```

#### Permission Denied
```bash
# Make executable (Linux/macOS)
chmod +x /usr/local/bin/takwin

# Run as administrator (Windows)
# Right-click PowerShell -> "Run as Administrator"
```

#### SSL/TLS Errors
```bash
# Update certificates (Linux)
sudo apt-get update && sudo apt-get install ca-certificates

# Use insecure download (not recommended)
curl -k -L https://github.com/hakkim/takwin/releases/latest/download/takwin-linux-amd64.tar.gz
```

### Getting Help

- **GitHub Issues**: https://github.com/hakkim/takwin/issues
- **Discussions**: https://github.com/hakkim/takwin/discussions
- **Documentation**: https://github.com/hakkim/takwin/docs

## Uninstallation

### Linux/macOS
```bash
# Remove binary
sudo rm /usr/local/bin/takwin

# Remove Homebrew installation
brew uninstall takwin
```

### Windows
```bash
# Remove from Program Files
Remove-Item "C:\Program Files\Takwin" -Recurse

# Remove from PATH
# Use System Properties -> Environment Variables
```

## Next Steps

After installation:

1. **Read the [Quick Start Guide](quickstart.md)**
2. **Check out [Examples](examples.md)**
3. **Review [Configuration Reference](configuration.md)**
4. **Join the community discussions**
