# Documentation Setup Guide

This guide explains how to work with the Takwin documentation.

## Prerequisites

- Python 3.x
- pip (Python package manager)

## Installation

### 1. Install MkDocs and Dependencies

```bash
# Install from requirements.txt
pip install -r requirements.txt

# Or install individually
pip install mkdocs mkdocs-material mkdocs-git-revision-date-localized-plugin mkdocs-minify-plugin
```

### 2. Verify Installation

```bash
mkdocs --version
```

## Local Development

### Serve Documentation Locally

```bash
# Using Task
task docs-serve

# Or directly
mkdocs serve
```

Visit: http://localhost:8000

The site will auto-reload when you make changes.

### Build Documentation

```bash
# Using Task
task docs-build

# Or directly
mkdocs build
```

Output will be in `site/` directory.

## Deployment

### Deploy to GitHub Pages

```bash
# Using Task
task docs-deploy

# Or directly
mkdocs gh-deploy --force
```

This will:
1. Build the documentation
2. Push to `gh-pages` branch
3. Make it available at: https://hakkim.github.io/takwin

### Automatic Deployment

Documentation is automatically deployed when:
- Changes are pushed to `main` branch
- Changes are made to `docs/` or `mkdocs.yml`

See `.github/workflows/docs.yml` for the workflow.

## Documentation Structure

```
docs/
├── index.md                 # Home page
├── installation.md          # Installation guide
├── quickstart.md           # Quick start tutorial
├── configuration.md        # Configuration reference
├── cli.md                  # CLI reference
├── examples.md             # Usage examples
├── architecture.md         # Architecture documentation
├── CONTRIBUTING.md         # Contributing guide
├── CHANGELOG.md            # Changelog
├── SECURITY.md             # Security policy
├── LICENSE.md              # License
├── stylesheets/
│   └── extra.css          # Custom CSS
└── javascripts/
    └── extra.js           # Custom JavaScript
```

## Writing Documentation

### Markdown Extensions

MkDocs Material supports many Markdown extensions:

#### Code Blocks

\`\`\`python
def hello():
    print("Hello, World!")
\`\`\`

#### Admonitions

```markdown
!!! note
    This is a note.

!!! warning
    This is a warning.

!!! tip
    This is a tip.
```

#### Tabs

```markdown
=== "Linux"
    Linux instructions here

=== "macOS"
    macOS instructions here

=== "Windows"
    Windows instructions here
```

#### Tables

```markdown
| Column 1 | Column 2 |
|----------|----------|
| Value 1  | Value 2  |
```

#### Mermaid Diagrams

\`\`\`mermaid
graph TD
    A[Start] --> B[Process]
    B --> C[End]
\`\`\`

### Best Practices

1. **Use descriptive headings**
   - Clear hierarchy (H1 > H2 > H3)
   - Descriptive titles

2. **Include code examples**
   - Show real-world usage
   - Include expected output

3. **Add navigation**
   - Link related pages
   - Use "Next steps" sections

4. **Keep it concise**
   - Short paragraphs
   - Bullet points for lists
   - Clear, simple language

5. **Test locally**
   - Always preview changes
   - Check all links work
   - Verify code examples

## Configuration

### mkdocs.yml

Main configuration file:

```yaml
site_name: Takwin
theme:
  name: material
nav:
  - Home: index.md
  - Getting Started:
    - Installation: installation.md
```

### Theme Customization

Edit `docs/stylesheets/extra.css` for custom styles.

Edit `docs/javascripts/extra.js` for custom JavaScript.

## Troubleshooting

### Port Already in Use

```bash
# Use different port
mkdocs serve -a localhost:8001
```

### Build Errors

```bash
# Clean and rebuild
rm -rf site/
mkdocs build --strict
```

### Plugin Errors

```bash
# Reinstall dependencies
pip install --upgrade -r requirements.txt
```

## Resources

- [MkDocs Documentation](https://www.mkdocs.org/)
- [Material for MkDocs](https://squidfunk.github.io/mkdocs-material/)
- [Markdown Guide](https://www.markdownguide.org/)

## Support

For documentation issues:
- Check this guide
- Review MkDocs documentation
- Ask in GitHub Discussions
