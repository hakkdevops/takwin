# Troubleshooting

Common issues and solutions for Takwin users.

## Installation Issues

### Binary Not Found

**Problem:** `takwin: command not found`

**Solution:**
```bash
# Ensure binary is in PATH
echo $PATH

# Add to PATH (Linux/macOS)
export PATH=$PATH:/usr/local/bin

# Or move binary to PATH location
sudo mv takwin /usr/local/bin/
```

### Permission Denied

**Problem:** `Permission denied` when running takwin

**Solution:**
```bash
# Make binary executable
chmod +x takwin

# Or run with sudo for installation
sudo mv takwin /usr/local/bin/
```

## Build Issues

### Compiler Not Found

**Problem:** `gcc: command not found` or `clang: command not found`

**Solution:**

=== "Linux"
    ```bash
    # Install GCC
    sudo apt-get install build-essential
    
    # Or Clang
    sudo apt-get install clang
    ```

=== "macOS"
    ```bash
    # Install Xcode Command Line Tools
    xcode-select --install
    
    # Or install via Homebrew
    brew install gcc
    ```

=== "Windows"
    ```powershell
    # Install MinGW
    choco install mingw
    
    # Or install MSVC via Visual Studio
    ```

### Source Files Not Found

**Problem:** `no source files found for target`

**Solution:**
1. Check glob patterns in `build.toml`
2. Verify file paths are correct
3. Use absolute paths or relative to project root

```toml
[[targets]]
name = "myapp"
sources = ["src/*.cpp"]  # Check this pattern
```

### Build Fails Silently

**Problem:** Build completes but no output

**Solution:**
```bash
# Run with verbose output
takwin build -v

# Check output directory
ls -la build/
```

## Configuration Issues

### Invalid TOML Syntax

**Problem:** `failed to parse build.toml`

**Solution:**
- Check TOML syntax
- Validate with online TOML validator
- Common issues:
  - Missing quotes around strings
  - Incorrect array syntax
  - Duplicate keys

```toml
# ❌ Wrong
[project]
name = hello  # Missing quotes

# ✅ Correct
[project]
name = "hello"
```

### Target Not Found

**Problem:** `target 'xyz' not found`

**Solution:**
```bash
# List available targets
takwin list-targets

# Check target name in build.toml
[[targets]]
name = "correct_name"  # Use this name
```

## Performance Issues

### Slow Build Times

**Solutions:**
1. Use optimization flags
2. Enable parallel builds (when available)
3. Use incremental builds (when available)

```toml
[build]
optimization = "O2"  # Use appropriate level
```

### High Memory Usage

**Solutions:**
1. Build targets individually
2. Reduce parallel jobs
3. Close other applications

## Platform-Specific Issues

### Windows: Path Issues

**Problem:** Paths with spaces or special characters

**Solution:**
```toml
# Use forward slashes
sources = ["C:/Program Files/project/src/*.cpp"]

# Or escape backslashes
sources = ["C:\\Program Files\\project\\src\\*.cpp"]
```

### macOS: Code Signing

**Problem:** "Developer cannot be verified"

**Solution:**
```bash
# Remove quarantine attribute
xattr -d com.apple.quarantine takwin

# Or allow in System Preferences > Security & Privacy
```

### Linux: Library Not Found

**Problem:** `error while loading shared libraries`

**Solution:**
```bash
# Update library cache
sudo ldconfig

# Or set LD_LIBRARY_PATH
export LD_LIBRARY_PATH=/path/to/libs:$LD_LIBRARY_PATH
```

## Getting Help

If you can't find a solution:

1. **Check Documentation**
   - [Installation Guide](installation.md)
   - [Configuration Reference](configuration.md)
   - [CLI Reference](cli.md)

2. **Search Issues**
   - [GitHub Issues](https://github.com/hakkdevops/takwin/issues)
   - Search for similar problems

3. **Ask for Help**
   - [GitHub Discussions](https://github.com/hakkdevops/takwin/discussions)
   - Provide:
     - Takwin version (`takwin --version`)
     - Operating system
     - Compiler version
     - `build.toml` content
     - Error messages

4. **Report a Bug**
   - [Create an Issue](https://github.com/hakkdevops/takwin/issues/new)
   - Include reproduction steps
   - Attach relevant files

## Debug Mode

Enable verbose output for debugging:

```bash
# Verbose build
takwin build -v

# Very verbose (if available)
takwin build -vv
```

## Common Error Messages

### "no targets defined"

**Cause:** Empty or missing `[[targets]]` section

**Fix:** Add at least one target in `build.toml`

### "invalid type 'xyz'"

**Cause:** Invalid target type

**Fix:** Use: `executable`, `static_library`, or `shared_library`

### "project name is required"

**Cause:** Missing `[project]` section or `name` field

**Fix:** Add project configuration:
```toml
[project]
name = "myproject"
version = "1.0.0"
```

## Still Having Issues?

Contact us:
- **GitHub Issues**: https://github.com/hakkdevops/takwin/issues
- **Discussions**: https://github.com/hakkdevops/takwin/discussions
- **Email**: support@takwin.dev (if available)
