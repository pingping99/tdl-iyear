// Package prog is responsible for handling standard output (stdout) for progress tracking and display.
//
// This file contains utilities for managing individual progress trackers and pinned system information
// that are displayed to standard output via the progress.Writer.
package prog

import (
	"context"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/progress"

	"github.com/iyear/tdl/pkg/ps"
)

// AppendTracker creates and appends a new progress tracker to the given progress writer.
// The tracker will be displayed on standard output when the writer renders.
//
// Parameters:
//   - pw: The progress writer that outputs to stdout
//   - formatter: Units formatter for the tracker
//   - message: Display message for the tracker
//   - total: Total units to track
//
// Returns:
//   - *progress.Tracker: The created tracker that will be rendered to stdout
func AppendTracker(pw progress.Writer, formatter progress.UnitsFormatter, message string, total int64) *progress.Tracker {
	units := progress.UnitsBytes
	units.Formatter = formatter

	tracker := progress.Tracker{
		Message: message,
		Total:   total,
		Units:   units,
	}

	pw.AppendTracker(&tracker)

	return &tracker
}

// EnablePS enables pinned messages with ps info: cpu, memory, goroutines.
// These system performance metrics are continuously updated and displayed at the top
// of the progress output on standard output.
//
// Parameters:
//   - ctx: Context for cancellation
//   - pw: The progress writer that outputs to stdout
func EnablePS(ctx context.Context, pw progress.Writer) {
	go func() {
		f := func() { pw.SetPinnedMessages(strings.Join(ps.Humanize(ctx), " ")) }
		f()

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				pw.SetPinnedMessages()
				return
			case <-ticker.C:
				f()
			}
		}
	}()
}
