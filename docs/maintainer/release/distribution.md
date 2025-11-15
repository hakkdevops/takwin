# Takwin Release Distribution Guide

## Where Your Product Will Be Released

When you create a release tag (e.g., `v0.1.0-beta.1`), your product will be automatically distributed to multiple platforms:

## 1. 🐙 GitHub Releases (Primary)

**Location:** https://github.com/hakkdevops/takwin/releases

**What Gets Published:**
- ✅ Release binaries for all platforms:
  - `takwin-linux-amd64.tar.gz` (Linux x64)
  - `takwin-linux-arm64.tar.gz` (Linux ARM64)
  - `takwin-darwin-amd64.tar.gz` (macOS Intel)
  - `takwin-darwin-arm64.tar.gz` (macOS Apple Silicon)
  - `takwin-windows-amd64.zip` (Windows x64)
- ✅ `checksums.txt` (SHA256 checksums for verification)
- ✅ Release notes (auto-generated)
- ✅ Source code (automatic GitHub archive)

**Beta vs Stable:**
- Beta releases: Marked as "Pre-release" 🟡
- Stable releases: Marked as "Latest release" 🟢

**Users Can Download:**
```bash
# Direct download
curl -L https://github.com/hakkdevops/takwin/releases/download/v0.1.0-beta.1/takwin-linux-amd64.tar.gz | tar -xz

# Or browse and download from web interface
```

---

## 2. 🐳 GitHub Container Registry (Docker)

**Location:** https://github.com/hakkdevops/takwin/packages

**What Gets Published:**
- ✅ Docker images for multiple architectures:
  - `ghcr.io/hakkdevops/takwin:v0.1.0-beta.1` (version tag)
  - `ghcr.io/hakkdevops/takwin:latest` (latest tag)
- ✅ Multi-arch support (linux/amd64, linux/arm64)

**Users Can Use:**
```bash
# Pull the image
docker pull ghcr.io/hakkdevops/takwin:v0.1.0-beta.1

# Or use latest
docker pull ghcr.io/hakkdevops/takwin:latest

# Run in current directory
docker run --rm -v $(pwd):/workspace -w /workspace ghcr.io/hakkdevops/takwin:latest build
```

---

## 3. 🍺 Homebrew (macOS - Stable Only)

**Location:** https://github.com/hakkim/homebrew-tap

**What Gets Published:**
- ✅ Homebrew formula (only for stable releases, not beta)
- ✅ Automatic formula updates

**Users Can Install:**
```bash
# Add tap
brew tap hakkim/tap

# Install
brew install takwin

# Update
brew upgrade takwin
```

**Note:** Homebrew updates only happen for stable releases (v0.1.0), not beta releases (v0.1.0-beta.1)

---

## 4. 📦 Go Package Registry

**Location:** https://pkg.go.dev/github.com/hakkdevops/takwin

**What Gets Published:**
- ✅ Go module documentation
- ✅ API reference
- ✅ Package information

**Developers Can Install:**
```bash
# Install as a Go tool
go install github.com/hakkdevops/takwin@latest

# Or specific version
go install github.com/hakkdevops/takwin@v0.1.0-beta.1
```

---

## Release Workflow Summary

When you push a tag, here's what happens automatically:

### Step 1: Tests Run
```
✓ Run all unit tests
✓ Run race detection tests
✓ Verify code quality
```

### Step 2: Build Binaries
```
✓ Build for Linux (amd64, arm64)
✓ Build for macOS (amd64, arm64)
✓ Build for Windows (amd64)
✓ Generate checksums
✓ Create archives (.tar.gz, .zip)
```

### Step 3: Create GitHub Release
```
✓ Upload all binaries
✓ Upload checksums
✓ Generate release notes
✓ Mark as pre-release (for beta) or latest (for stable)
```

### Step 4: Build Docker Images
```
✓ Build multi-arch Docker images
✓ Push to GitHub Container Registry
✓ Tag with version and 'latest'
```

### Step 5: Update Homebrew (Stable Only)
```
✓ Update Homebrew formula
✓ Push to homebrew-tap repository
```

---

## Distribution Channels Comparison

| Channel | Beta Support | Auto-Update | Platform | Installation |
|---------|--------------|-------------|----------|--------------|
| **GitHub Releases** | ✅ Yes | Manual | All | Download binary |
| **Docker** | ✅ Yes | `docker pull` | All | Container |
| **Homebrew** | ❌ No | `brew upgrade` | macOS | Package manager |
| **Go Install** | ✅ Yes | Manual | All | Go toolchain |

---

## User Installation Methods

### Method 1: Direct Download (Recommended for Beta)
```bash
# Linux
curl -L https://github.com/hakkdevops/takwin/releases/download/v0.1.0-beta.1/takwin-linux-amd64.tar.gz | tar -xz
sudo mv takwin-linux-amd64 /usr/local/bin/takwin

# macOS
curl -L https://github.com/hakkdevops/takwin/releases/download/v0.1.0-beta.1/takwin-darwin-amd64.tar.gz | tar -xz
sudo mv takwin-darwin-amd64 /usr/local/bin/takwin

# Windows
# Download from GitHub releases page and extract
```

### Method 2: Docker
```bash
# Pull and run
docker pull ghcr.io/hakkdevops/takwin:v0.1.0-beta.1
docker run --rm -v $(pwd):/workspace -w /workspace ghcr.io/hakkdevops/takwin:v0.1.0-beta.1 build

# Create alias
alias takwin='docker run --rm -v $(pwd):/workspace -w /workspace ghcr.io/hakkdevops/takwin:latest'
```

### Method 3: Go Install
```bash
# Install latest
go install github.com/hakkdevops/takwin@latest

# Install specific version
go install github.com/hakkdevops/takwin@v0.1.0-beta.1
```

### Method 4: Homebrew (Stable Only)
```bash
# Only works for stable releases (v0.1.0, not v0.1.0-beta.1)
brew tap hakkim/tap
brew install takwin
```

---

## Monitoring Your Release

### 1. GitHub Actions
**URL:** https://github.com/hakkdevops/takwin/actions

Watch the "Release" workflow to ensure:
- ✅ Tests pass
- ✅ Builds complete
- ✅ Docker images push successfully
- ✅ Release is created

### 2. GitHub Releases
**URL:** https://github.com/hakkdevops/takwin/releases

Verify:
- ✅ All binaries are uploaded
- ✅ Checksums are present
- ✅ Release notes are correct
- ✅ Pre-release badge (for beta)

### 3. GitHub Packages
**URL:** https://github.com/hakkdevops/takwin/packages

Check:
- ✅ Docker image is published
- ✅ Tags are correct
- ✅ Multi-arch support

### 4. Download Statistics
**URL:** https://github.com/hakkdevops/takwin/releases

GitHub provides download counts for each binary.

---

## Release Artifacts

After a successful release, users can access:

### Binaries
```
takwin-linux-amd64.tar.gz      (~8MB)
takwin-linux-arm64.tar.gz      (~8MB)
takwin-darwin-amd64.tar.gz     (~8MB)
takwin-darwin-arm64.tar.gz     (~8MB)
takwin-windows-amd64.zip       (~8MB)
checksums.txt                  (~1KB)
```

### Docker Images
```
ghcr.io/hakkdevops/takwin:v0.1.0-beta.1
ghcr.io/hakkdevops/takwin:latest
```

### Source Code
```
Source code (zip)
Source code (tar.gz)
```

---

## Security & Verification

### Checksums
Users can verify downloads:
```bash
# Download checksums
curl -L https://github.com/hakkdevops/takwin/releases/download/v0.1.0-beta.1/checksums.txt

# Verify
sha256sum -c checksums.txt
```

### Docker Image Verification
```bash
# Inspect image
docker inspect ghcr.io/hakkdevops/takwin:v0.1.0-beta.1

# Check labels
docker inspect ghcr.io/hakkdevops/takwin:v0.1.0-beta.1 | grep -A 10 Labels
```

---

## Announcement Strategy

After release is published:

### 1. GitHub Discussions
- Create announcement post
- Link to release
- Highlight new features
- Request feedback

### 2. README Badge
Update README.md with latest release badge:
```markdown
[![GitHub release](https://img.shields.io/github/release/hakkdevops/takwin.svg)](https://github.com/hakkdevops/takwin/releases)
```

### 3. Social Media (Optional)
- Twitter/X
- LinkedIn
- Reddit (r/cpp, r/golang)
- Dev.to

---

## Troubleshooting

### Release Workflow Fails
1. Check GitHub Actions logs
2. Verify all tests pass locally
3. Check go.mod dependencies
4. Ensure tag format is correct

### Binaries Not Uploaded
1. Check build step in workflow
2. Verify dist/ directory is created
3. Check file permissions

### Docker Push Fails
1. Verify GITHUB_TOKEN permissions
2. Check Docker build logs
3. Ensure Dockerfile is valid

### Homebrew Not Updated
1. Check if release is stable (not beta)
2. Verify HOMEBREW_TAP_TOKEN secret
3. Check homebrew-tap repository

---

## Quick Reference

### Release Command
```bash
git tag -a v0.1.0-beta.1 -m "Release v0.1.0-beta.1"
git push origin v0.1.0-beta.1
```

### Monitor Release
```bash
# Watch workflow
https://github.com/hakkdevops/takwin/actions

# Check release
https://github.com/hakkdevops/takwin/releases

# Verify Docker
docker pull ghcr.io/hakkdevops/takwin:v0.1.0-beta.1
```

### Test Installation
```bash
# Download and test
curl -L https://github.com/hakkdevops/takwin/releases/download/v0.1.0-beta.1/takwin-linux-amd64.tar.gz | tar -xz
./takwin-linux-amd64 --version
```

---

## Support

For release issues:
- Check GitHub Actions logs
- Review this guide
- Ask in GitHub Discussions
- Open an issue if needed
