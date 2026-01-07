# pkg/prog - Standard Output Handler for Progress Tracking

## Overview

**`pkg/prog` 是负责处理标准输出 (stdout) 的包** (This package is responsible for handling standard output)

This package provides the primary interface for rendering progress information to standard output in the tdl application. It wraps the `github.com/jedib0t/go-pretty/v6/progress` library with custom configuration and styling.

## Key Responsibilities

1. **Standard Output Management**: Creates and configures `progress.Writer` instances that write directly to `os.Stdout`
2. **Progress Bar Display**: Renders download/upload progress bars and statistics to the terminal
3. **System Metrics**: Displays CPU, memory, and goroutine usage alongside progress information
4. **Styling & Formatting**: Applies consistent colors, widths, and formats across all progress displays

## Architecture

```
pkg/prog (Standard Output Handler)
├── prog.go      - Main progress writer creation and configuration
└── tracker.go   - Individual tracker management and system metrics
```

## Standard Output Flow

```
Application Operation (download/upload/export)
    ↓
prog.New() creates progress.Writer
    ↓
progress.NewWriter() internally sets output to os.Stdout
    ↓
pw.Render() starts goroutine that writes to stdout
    ↓
Progress bars and metrics displayed in terminal
```

## Usage Examples

### Creating a Progress Writer

```go
// Create a progress writer that outputs to stdout
pw := prog.New(progress.FormatNumber)

// Start rendering to stdout in background
go pw.Render()

// Add a tracker for a download operation
tracker := prog.AppendTracker(pw, progress.FormatNumber, "Downloading file.mp4", 1024000)

// Enable system performance metrics display
prog.EnablePS(ctx, pw)

// Wait for all progress to complete
prog.Wait(ctx, pw)
```

### Used Throughout Application

- **app/dl/dl.go**: Download progress tracking
- **app/up/up.go**: Upload progress tracking  
- **app/chat/export.go**: Message export progress
- **app/chat/users.go**: User enumeration progress
- **app/forward/forward.go**: Message forwarding progress

## Technical Details

### Default Output Destination

The `progress.NewWriter()` function from `github.com/jedib0t/go-pretty/v6/progress` creates a writer that outputs to `os.Stdout` by default. This behavior can be changed using the `SetOutputWriter(io.Writer)` method if needed, but the current implementation uses the default stdout.

### Terminal Width Detection

The package automatically detects terminal width using `github.com/kopoli/go-terminal-size` and adjusts the progress bar dimensions accordingly to fit the user's terminal.

### Concurrent Safety

The progress writer handles concurrent updates safely, allowing multiple goroutines to update different trackers simultaneously while maintaining correct output to stdout.

## Related Packages

- **`github.com/jedib0t/go-pretty/v6/progress`**: The underlying library that handles stdout rendering
- **`pkg/ps`**: Provides system performance statistics displayed via EnablePS
- **`github.com/fatih/color`**: Used for colorizing output messages
- **`github.com/kopoli/go-terminal-size`**: Used for detecting terminal dimensions

## Other Standard Output Usage

While `pkg/prog` is the primary package for structured progress output to stdout, other parts of the codebase also write to stdout using `fmt.Print*` functions for:
- Simple status messages
- JSON output (e.g., `app/chat/ls.go`, `app/chat/export.go`)
- Table displays (e.g., `app/extension/extension.go`)
- QR code display (e.g., `app/login/qr.go`)

However, for all **progress tracking and long-running operation visualization**, `pkg/prog` is the designated handler.
