# Standard Output Handling Package Analysis

## Question
Identify which package is handling standard output

## Answer

The primary package handling standard output is: **`github.com/jedib0t/go-pretty/v6/progress`**

## Detailed Analysis

### 1. Primary Output Package

#### `github.com/jedib0t/go-pretty/v6/progress`
This is the core package for handling standard output, used to display progress bars and download status.

**Default Behavior:**
- When creating a `progress.Writer`, if no output target is set via `SetOutputWriter()`, it defaults to outputting to `os.Stdout`
- Source location: `github.com/jedib0t/go-pretty/v6@v6.5.0/progress/progress.go:311`

```go
// if no output writer has been set, output to STDOUT
if p.outputWriter == nil {
    p.outputWriter = os.Stdout
}
```

**Usage in the Project:**
- Created in `pkg/prog/prog.go`'s `New()` function
- Used by the following modules:
  - `app/dl/dl.go` - Download progress display
  - `app/up/up.go` - Upload progress display
  - `app/forward/forward.go` - Forward progress display
  - `app/chat/users.go` - User export progress
  - `app/chat/export.go` - Chat export progress

### 2. Auxiliary Output Packages

#### `github.com/fatih/color`
Used for colored output and terminal text formatting.

**Use Cases:**
- Error message display (red)
- Success message display (green)
- Warning message display (yellow)
- Information prompts

**Main Usage Locations:**
- `main.go` - Error output
- `cmd/login.go` - Login-related prompts
- `cmd/version.go` - Version information display
- Status messages in various app sub-packages

#### `fmt` Standard Package
Used for basic formatted output.

**Use Cases:**
- Simple text output
- Table rendering results (via `go-pretty/table`)
- General information printing

### 3. Package Dependency Chain

```
main.go
  └── cmd/root.go (uses fatih/color for error output)
       └── app/dl/dl.go (uses progress.Writer for download progress)
            └── pkg/prog/prog.go (creates and configures progress.Writer)
                 └── github.com/jedib0t/go-pretty/v6/progress (defaults to os.Stdout)
```

### 4. Extension Command Output

For extension commands, standard output is handled as follows:

**Location:** `cmd/extension.go:148`
```go
cmd.AddCommand(NewExtensionCmd(em, e, os.Stdin, os.Stdout, os.Stderr))
```

**Location:** `pkg/extensions/manager.go:90-92`
```go
cmd.Stdin = stdin
cmd.Stdout = stdout
cmd.Stderr = stderr
```

Extension commands directly use the passed `os.Stdout` as standard output.

### 5. Summary

**Main Standard Output Handling Packages:**
1. **`github.com/jedib0t/go-pretty/v6/progress`** - Progress bars and status tracking (most important)
2. **`github.com/fatih/color`** - Colored text output
3. **`fmt`** - Standard formatted output
4. **Direct use of `os.Stdout`** - Extension commands

All of these ultimately output to `os.Stdout` (standard output) or `os.Stderr` (standard error).

## Related Files

### Core Files
- `pkg/prog/prog.go` - Progress writer creation and configuration
- `pkg/prog/tracker.go` - Progress tracker implementation

### Usage Files
- `app/dl/progress.go` - Download progress implementation
- `app/up/progress.go` - Upload progress implementation
- `app/forward/progress.go` - Forward progress implementation

### Configuration Files
- `go.mod` - Dependency declaration (line 26): `github.com/jedib0t/go-pretty/v6 v6.5.0`
