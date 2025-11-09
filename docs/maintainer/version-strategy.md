# Takwin Version Strategy

## Current Version: v0.1.0

Takwin follows [Semantic Versioning 2.0.0](https://semver.org/).

## Version Format

```
v<MAJOR>.<MINOR>.<PATCH>[-<PRERELEASE>]
```

### Examples:
- `v0.1.0-beta.1` - First beta release
- `v0.1.0-beta.2` - Second beta release  
- `v0.1.0-rc.1` - Release candidate
- `v0.1.0` - First stable release
- `v0.1.1` - Patch release
- `v0.2.0` - Minor version bump
- `v1.0.0` - Major version (production ready)

## Version Progression

### Phase 1: Initial Development (v0.x.x)
**Current Phase**

- `v0.1.0-beta.1` - First beta release (current target)
- `v0.1.0-beta.2` - Bug fixes and improvements
- `v0.1.0-rc.1` - Release candidate
- `v0.1.0` - First stable release

**Characteristics:**
- API may change
- Breaking changes allowed
- Focus on core functionality
- Gathering user feedback

### Phase 2: Stabilization (v0.2.x - v0.9.x)

- `v0.2.0` - Additional features
- `v0.3.0` - More features and improvements
- `v0.9.0` - Feature complete, preparing for v1.0.0

**Characteristics:**
- API stabilizing
- Fewer breaking changes
- Focus on stability and performance
- Comprehensive testing

### Phase 3: Production Ready (v1.0.0+)

- `v1.0.0` - Production ready release
- `v1.1.0` - New features (backward compatible)
- `v1.0.1` - Bug fixes only
- `v2.0.0` - Major changes (may break compatibility)

**Characteristics:**
- Stable API
- Semantic versioning strictly followed
- Long-term support
- Enterprise ready

## When to Bump Versions

### MAJOR version (v1.0.0 → v2.0.0)
- Breaking API changes
- Incompatible configuration changes
- Major architectural changes
- **Only after v1.0.0**

### MINOR version (v0.1.0 → v0.2.0)
- New features (backward compatible)
- New compiler support
- New platform support
- Significant improvements

### PATCH version (v0.1.0 → v0.1.1)
- Bug fixes
- Security patches
- Documentation updates
- Performance improvements (no API changes)

### PRE-RELEASE versions
- `-beta.X` - Beta testing phase
- `-rc.X` - Release candidate
- `-alpha.X` - Alpha testing (if needed)

## Version 0.x.x Special Rules

During the 0.x.x phase:
- **MINOR** version bumps MAY include breaking changes
- **PATCH** version bumps should be backward compatible
- API is not yet stable
- Users should expect changes

## Roadmap

### v0.1.0 (First Release)
- [x] Core build functionality
- [x] Multi-compiler support
- [x] Cross-platform support
- [ ] Command execution (currently dry-run)
- [ ] Comprehensive logging

### v0.2.0 (Planned)
- [ ] Actual command execution
- [ ] Incremental builds
- [ ] Build caching
- [ ] Dependency tracking
- [ ] Enhanced error messages

### v0.3.0 (Planned)
- [ ] Plugin system
- [ ] Package manager integration
- [ ] IDE integration
- [ ] Build visualization

### v1.0.0 (Production Ready)
- [ ] Stable API
- [ ] Complete documentation
- [ ] Enterprise features
- [ ] Long-term support commitment

## Release Checklist

Before any release:
- [ ] All tests pass
- [ ] Documentation updated
- [ ] CHANGELOG.md updated
- [ ] Version bumped in code
- [ ] Examples tested
- [ ] Cross-platform builds verified

## Version in Code

The version is defined in `cmd/root.go`:

```go
Version: "0.1.0",
```

This should be updated before each release.

## Git Tags

All releases are tagged in git:

```bash
# Beta release
git tag -a v0.1.0-beta.1 -m "Release v0.1.0-beta.1"

# Stable release
git tag -a v0.1.0 -m "Release v0.1.0"
```

## Communication

### Beta Releases
- Marked as "Pre-release" on GitHub
- Announced in GitHub Discussions
- Feedback requested from community

### Stable Releases
- Full release notes
- Migration guide (if needed)
- Announcement on all channels
- Homebrew formula updated

## Support Policy

### Current (v0.x.x)
- Active development
- Bug fixes and features
- Breaking changes allowed

### Stable (v1.x.x+)
- Bug fixes for current major version
- Security patches for previous major version
- No breaking changes in minor/patch versions

## Questions?

For questions about versioning:
- Check this document
- Review CHANGELOG.md
- Ask in GitHub Discussions
