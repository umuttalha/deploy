// Package ui renders clean, minimal CLI output in the style of Vercel CLI.
// All output uses plain fmt — no colors, no emojis, no lipgloss.
package ui

import (
	"fmt"
	"io"
	"strings"
)

// Header prints a single dim banner line: "> Name version · repo"
func Header(w io.Writer, name, version string) {
	fmt.Fprintf(w, "> %s %s · github.com/umuttalha/deploy\n", name, version)
}

// Info prints an informational line prefixed with ">".
func Info(w io.Writer, format string, a ...any) {
	fmt.Fprintf(w, "> %s\n", fmt.Sprintf(format, a...))
}

// Error prints an error line prefixed with "> Error:".
func Error(w io.Writer, msg string) {
	fmt.Fprintf(w, "> Error: %s\n", msg)
}

// Hint prints an indented hint line below an error.
func Hint(w io.Writer, msg string) {
	fmt.Fprintf(w, "  %s\n", msg)
}

// Fail prints an error with an optional hint.
func Fail(w io.Writer, msg, hint string) {
	Error(w, msg)
	if hint != "" {
		Hint(w, hint)
	}
}

// Setting prints a key/value pair as "  - Key: value".
func Setting(w io.Writer, key, val string) {
	fmt.Fprintf(w, "  - %s  %s\n", padRight(key+":", 12), val)
}

// Answered prints a completed prompt in Vercel style: "? Label value".
func Answered(w io.Writer, label, value string) {
	fmt.Fprintf(w, "? %s %s\n", label, value)
}

func padRight(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return s + strings.Repeat(" ", width-len(s))
}
