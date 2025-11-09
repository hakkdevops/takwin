# Installing Go on Windows

The Chocolatey installation failed due to permission issues. Here are several ways to install Go on Windows:

## Method 1: Run PowerShell as Administrator (Recommended)

1. **Close current PowerShell/Command Prompt**
2. **Right-click on PowerShell** and select "Run as Administrator"
3. **Run the Chocolatey command**:
   ```powershell
   choco install golang -y
   ```
4. **Restart your terminal** after installation
5. **Verify installation**:
   ```powershell
   go version
   ```

## Method 2: Direct Download from Go Website

1. **Visit**: https://golang.org/dl/
2. **Download**: `go1.21.5.windows-amd64.msi` (or latest version)
3. **Run the installer** and follow the prompts
4. **Restart your terminal**
5. **Verify installation**:
   ```powershell
   go version
   ```

## Method 3: Winget (Windows Package Manager)

If you have Windows 10/11 with winget:

```powershell
winget install GoLang.Go
```

## Method 4: Scoop Package Manager

If you have Scoop installed:

```powershell
scoop install go
```

## Method 5: Manual Installation

1. **Download**: https://golang.org/dl/go1.21.5.windows-amd64.zip
2. **Extract** to `C:\Go` (or your preferred location)
3. **Add to PATH**:
   - Open System Properties → Advanced → Environment Variables
   - Add `C:\Go\bin` to your PATH variable
4. **Restart terminal**
5. **Verify**:
   ```powershell
   go version
   ```

## After Installation

Once Go is installed, navigate to the takwin directory and run:

```powershell
# Navigate to project
cd takwin

# Initialize Go module
go mod tidy

# Build the project
go build -o takwin.exe main.go

# Run tests
go test ./...

# Test with examples
.\takwin.exe build --config examples\simple\build.toml
```

## Troubleshooting

### Go Command Not Found
- **Restart your terminal** after installation
- **Check PATH**: `echo $env:PATH` should include Go's bin directory
- **Reinstall** if PATH is not set correctly

### Permission Issues
- **Run as Administrator** when using package managers
- **Use direct download** if package managers fail
- **Check antivirus** - some antivirus software blocks installations

### Module Issues
- **Run `go mod tidy`** to fix dependency issues
- **Delete go.sum** and run `go mod download` if needed
- **Check internet connection** for downloading dependencies

## Verification Commands

After successful installation:

```powershell
# Check Go version
go version

# Check Go environment
go env

# Check GOPATH and GOROOT
go env GOPATH
go env GOROOT
```

Expected output:
```
go version go1.21.5 windows/amd64
```

## Next Steps

Once Go is installed:

1. **Navigate to takwin directory**
2. **Run setup commands**:
   ```powershell
   go mod tidy
   go build -o takwin.exe main.go
   go test ./...
   ```
3. **Test with examples**:
   ```powershell
   .\takwin.exe build --config examples\simple\build.toml
   ```
4. **Compare performance** with Python version

Choose the method that works best for your system. The direct download from the Go website is usually the most reliable option.
