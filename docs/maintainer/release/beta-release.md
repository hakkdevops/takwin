# Beta Release Checklist for v0.1.0-beta.1

## Pre-Release

- [ ] All tests pass locally
  ```bash
  task test
  task test-race
  ```

- [ ] Code is formatted
  ```bash
  task fmt
  ```

- [ ] Examples work
  ```bash
  task examples
  ```

- [ ] Update CHANGELOG.md with beta release notes

- [ ] Review and update documentation

## Release Steps

### 1. Commit Changes
```bash
git add .
git commit -m "chore: prepare for v0.1.0-beta.1 release"
git push origin main
```

### 2. Create and Push Tag
```bash
# Create annotated tag
git tag -a v0.1.0-beta.1 -m "Release v0.1.0-beta.1"

# Push tag to trigger release
git push origin v0.1.0-beta.1
```

### 3. Monitor Workflow
- [ ] Go to https://github.com/hakkdevops/takwin/actions
- [ ] Watch "Release" workflow
- [ ] Ensure all jobs pass (test, build, docker)

### 4. Verify Release
- [ ] Check https://github.com/hakkdevops/takwin/releases
- [ ] Verify "Pre-release" badge is shown
- [ ] Check all binaries are uploaded:
  - [ ] takwin-linux-amd64.tar.gz
  - [ ] takwin-linux-arm64.tar.gz
  - [ ] takwin-darwin-amd64.tar.gz
  - [ ] takwin-darwin-arm64.tar.gz
  - [ ] takwin-windows-amd64.zip
  - [ ] checksums.txt

### 5. Test Downloads
```bash
# Test Linux binary
curl -L https://github.com/hakkdevops/takwin/releases/download/v0.1.0-beta.1/takwin-linux-amd64.tar.gz | tar -xz
./takwin-linux-amd64 --version
./takwin-linux-amd64 --help
```

- [ ] Linux binary works
- [ ] macOS binary works (if available)
- [ ] Windows binary works (if available)

### 6. Verify Checksums
```bash
curl -L https://github.com/hakkdevops/takwin/releases/download/v0.1.0-beta.1/checksums.txt
sha256sum -c checksums.txt
```

- [ ] Checksums match

## Post-Release

- [ ] Announce in GitHub Discussions
- [ ] Share with beta testers
- [ ] Monitor for issues
- [ ] Collect feedback

## If Something Goes Wrong

### Delete and Recreate Release
```bash
# Delete local tag
git tag -d v0.1.0-beta.1

# Delete remote tag
git push origin :refs/tags/v0.1.0-beta.1

# Delete GitHub release from web interface

# Fix issues, then recreate tag
git tag -a v0.1.0-beta.1 -m "Release v0.1.0-beta.1"
git push origin v0.1.0-beta.1
```

## Notes

- Beta releases are automatically marked as "Pre-release"
- Homebrew formula will NOT be updated for beta releases
- Docker images will be tagged with version number
- All binaries include version info

## Ready to Release?

If all pre-release checks pass, run:

```bash
git tag -a v0.1.0-beta.1 -m "Release v0.1.0-beta.1" && git push origin v0.1.0-beta.1
```

Then monitor: https://github.com/hakkdevops/takwin/actions
