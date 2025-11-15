# Takwin Beta Release Guide

This guide walks you through creating a beta release for Takwin.

## Prerequisites

Before creating a beta release, ensure:

- [ ] All tests pass locally (`task test`)
- [ ] Code is formatted (`task fmt`)
- [ ] Examples work (`task examples`)
- [ ] Documentation is up to date
- [ ] CHANGELOG.md is updated with changes

## Beta Release Process

### 1. Update Version Information

Update the CHANGELOG.md with beta release notes:

```markdown
## [0.1.0-beta.1] - 2024-XX-XX

### Added
- List new features

### Changed
- List changes

### Fixed
- List bug fixes

### Known Issues
- List any known issues for beta testers
```

### 2. Run Local Tests

```bash
# Run all tests
task test

# Run with race detection
task test-race

# Test examples
task examples

# Build for all platforms
task build-all
```

### 3. Commit Changes

```bash
# Stage all changes
git add .

# Commit with descriptive message
git commit -m "chore: prepare for v0.1.0-beta.1 release"

# Push to remote
git push origin main
```

### 4. Create Beta Tag

Beta releases use semantic versioning with a `-beta.X` suffix:

```bash
# Create annotated tag for beta release
git tag -a v0.1.0-beta.1 -m "Release v0.1.0-beta.1"

# Push the tag to trigger release workflow
git push origin v0.1.0-beta.1
```

**Tag Format:**
- `v0.1.0-beta.1` - First beta release
- `v0.1.0-beta.2` - Second beta release
- `v0.1.0-rc.1` - Release candidate
- `v0.1.0` - First stable release
- `v0.2.0` - Minor version bump
- `v1.0.0` - Major version (production ready)

### 5. Monitor Release Workflow

After pushing the tag, the GitHub Actions workflow will automatically:

1. ✅ Run all tests
2. ✅ Build binaries for all platforms (Linux, macOS, Windows)
3. ✅ Create checksums
4. ✅ Generate release archives
5. ✅ Create GitHub release (marked as pre-release)
6. ✅ Build and push Docker images

**Monitor the workflow:**
- Go to: https://github.com/hakkdevops/takwin/actions
- Check the "Release" workflow
- Ensure all jobs complete successfully

### 6. Verify Release

Once the workflow completes:

1. **Check GitHub Release:**
   - Go to: https://github.com/hakkdevops/takwin/releases
   - Verify the beta release is marked as "Pre-release"
   - Check all binaries are uploaded

2. **Test Download:**
   ```bash
   # Test Linux binary
   curl -L https://github.com/hakkdevops/takwin/releases/download/v0.1.0-beta.1/takwin-linux-amd64.tar.gz | tar -xz
   ./takwin-linux-amd64 --version
   
   # Test Windows binary
   # Download and test on Windows
   ```

3. **Verify Checksums:**
   ```bash
   # Download checksums
   curl -L https://github.com/hakkdevops/takwin/releases/download/v0.1.0-beta.1/checksums.txt
   
   # Verify
   sha256sum -c checksums.txt
   ```

### 7. Announce Beta Release

After verification, announce the beta release:

1. **GitHub Discussions:**
   - Create a post in Discussions
   - Link to the release
   - Ask for feedback

2. **Update README (Optional):**
   - Add beta release badge
   - Link to beta installation instructions

## Beta Testing Checklist

Share this checklist with beta testers:

- [ ] Download and install beta release
- [ ] Test basic commands (`build`, `clean`, `list-targets`)
- [ ] Test with simple C++ project
- [ ] Test with complex multi-target project
- [ ] Test on different platforms (Linux, macOS, Windows)
- [ ] Report any issues on GitHub

## Troubleshooting

### Release Workflow Fails

**If tests fail:**
```bash
# Fix the issues locally
task test

# Commit fixes
git commit -am "fix: resolve test failures"
git push

# Delete the tag
git tag -d v0.1.0-beta.1
git push origin :refs/tags/v0.1.0-beta.1

# Create new tag
git tag -a v0.1.0-beta.1 -m "Release v0.1.0-beta.1"
git push origin v0.1.0-beta.1
```

**If build fails:**
- Check the GitHub Actions logs
- Ensure all dependencies are in go.mod
- Test cross-compilation locally

### Delete a Release

If you need to delete a beta release:

```bash
# Delete local tag
git tag -d v0.1.0-beta.1

# Delete remote tag
git push origin :refs/tags/v0.1.0-beta.1

# Delete GitHub release manually from the web interface
```

## Moving from Beta to Stable

When ready for stable release:

1. **Update CHANGELOG.md:**
   ```markdown
   ## [0.1.0] - 2024-XX-XX
   
   ### Release Notes
   - First stable release based on beta feedback
   - All beta issues resolved
   ```

2. **Create stable tag:**
   ```bash
   git tag -a v0.1.0 -m "Release v0.1.0"
   git push origin v0.1.0
   ```

3. **Stable releases will:**
   - Not be marked as pre-release
   - Update Homebrew formula automatically
   - Be tagged as `latest` in Docker

## Quick Reference

### Common Commands

```bash
# Local testing
task test              # Run tests
task test-race         # Race detection
task build             # Build binary
task examples          # Test examples

# Release process
git tag -a v0.1.0-beta.1 -m "Release v0.1.0-beta.1"
git push origin v0.1.0-beta.1

# Delete tag (if needed)
git tag -d v0.1.0-beta.1
git push origin :refs/tags/v0.1.0-beta.1
```

### Version Naming

- **Beta**: `v0.1.0-beta.1`, `v0.1.0-beta.2`
- **RC**: `v0.1.0-rc.1`, `v0.1.0-rc.2`
- **Stable**: `v0.1.0` (first release)
- **Patch**: `v0.1.1`
- **Minor**: `v0.2.0`
- **Major**: `v1.0.0` (production ready)

## Notes

- Beta releases are automatically marked as "Pre-release" on GitHub
- The workflow detects beta/rc versions by checking for `-` in the tag
- Homebrew formula updates only happen for stable releases
- Docker images are tagged with both version and `latest`
- All binaries are built with version info embedded

## Support

If you encounter issues during the release process:
- Check GitHub Actions logs
- Review this guide
- Ask in GitHub Discussions
