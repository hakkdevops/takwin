# Maintainer Guide

This section contains documentation for Takwin maintainers and core contributors.

## Overview

As a maintainer, you have additional responsibilities beyond regular contributors:

- 🔄 **Release Management** - Creating and publishing releases
- 🏗️ **Architecture Decisions** - Making technical decisions
- 👥 **Code Review** - Reviewing and merging pull requests
- 📊 **Project Planning** - Roadmap and feature planning
- 🔒 **Security** - Handling security issues
- 📈 **Metrics** - Monitoring project health

## Quick Links

### Release Management
- [Release Overview](release/overview.md) - Complete release process
- [Beta Releases](release/beta-release.md) - Beta release checklist
- [Stable Releases](release/stable-release.md) - Stable release process
- [Distribution](release/distribution.md) - Where releases are published

### Technical Documentation
- [Architecture](architecture.md) - System architecture and design
- [Version Strategy](version-strategy.md) - Versioning and roadmap
- [CI/CD](cicd.md) - Continuous integration and deployment

## Maintainer Responsibilities

### 1. Code Review

**Guidelines:**
- Review PRs within 48 hours
- Check code quality and tests
- Ensure documentation is updated
- Verify CI passes
- Be constructive and helpful

### 2. Release Management

**Process:**
1. Update CHANGELOG.md
2. Bump version in code
3. Create and push tag
4. Monitor release workflow
5. Verify artifacts
6. Announce release

See [Release Overview](release/overview.md) for details.

### 3. Issue Triage

**Workflow:**
- Label new issues appropriately
- Ask for clarification if needed
- Assign to milestones
- Close duplicates
- Link related issues

### 4. Security

**Handling Security Issues:**
1. Acknowledge receipt within 24 hours
2. Assess severity
3. Create private fix
4. Coordinate disclosure
5. Release security update

See [Security Policy](../reference/SECURITY.md).

### 5. Community Management

**Responsibilities:**
- Answer questions in Discussions
- Welcome new contributors
- Maintain Code of Conduct
- Foster inclusive environment

## Tools and Access

### Required Access

Maintainers need access to:
- ✅ GitHub repository (write access)
- ✅ GitHub Actions secrets
- ✅ Docker registry (ghcr.io)
- ✅ Homebrew tap repository

### Tools

**Development:**
- Go 1.21+
- Git
- Task (task runner)
- golangci-lint

**Release:**
- GitHub CLI (gh)
- Docker
- GPG key (for signing)

## Workflows

### Weekly Tasks

- [ ] Review open PRs
- [ ] Triage new issues
- [ ] Check CI/CD health
- [ ] Monitor security alerts
- [ ] Update dependencies

### Monthly Tasks

- [ ] Review roadmap progress
- [ ] Plan next release
- [ ] Update documentation
- [ ] Check metrics
- [ ] Community engagement

### Release Cycle

- **Beta releases**: As needed for testing
- **Stable releases**: Monthly or when ready
- **Patch releases**: As needed for bugs
- **Security releases**: Immediate

## Decision Making

### Technical Decisions

**Process:**
1. Create RFC (Request for Comments) issue
2. Discuss with maintainers
3. Gather community feedback
4. Make decision
5. Document in architecture docs

### Breaking Changes

**Requirements:**
- Major version bump
- Migration guide
- Deprecation warnings (if possible)
- Community notification

## Communication

### Channels

- **GitHub Issues**: Bug reports, features
- **GitHub Discussions**: Questions, ideas
- **Pull Requests**: Code changes
- **Releases**: Announcements

### Response Times

- **Security issues**: 24 hours
- **Bug reports**: 48 hours
- **Feature requests**: 1 week
- **Pull requests**: 48 hours

## Onboarding New Maintainers

### Checklist

- [ ] Grant repository access
- [ ] Add to maintainers team
- [ ] Share access credentials
- [ ] Review maintainer guide
- [ ] Pair on first release
- [ ] Introduce to community

### Expectations

- Active participation
- Timely responses
- Professional conduct
- Continuous learning
- Community focus

## Resources

### Internal Documentation
- [Architecture](architecture.md)
- [Release Process](release/overview.md)
- [Version Strategy](version-strategy.md)
- [CI/CD](cicd.md)

### External Resources
- [Semantic Versioning](https://semver.org/)
- [Keep a Changelog](https://keepachangelog.com/)
- [GitHub Actions](https://docs.github.com/en/actions)

## Questions?

If you have questions about maintainer responsibilities:
- Review this guide
- Ask other maintainers
- Check GitHub documentation
- Create a discussion

## Maintainer List

Current maintainers:
- **Hakkim** - Lead Maintainer

---

**Thank you for maintaining Takwin!** 🙏
