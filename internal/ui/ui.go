// Package ui renders styled CLI output (status lines, headers, key/value
// pairs, errors). All styles auto-disable on non-TTY output via lipgloss.
package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)
	infoStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
	warnStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("11"))
	dangerStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true)
	dimStyle     = lipgloss.NewStyle().Faint(true)
	labelStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("14")).Bold(true)
)

// Header prints a single-line branded banner followed by a blank line.
func Header(w io.Writer, name, version string) {
	fmt.Fprintln(w, dimStyle.Render(fmt.Sprintf("%s %s · github.com/umuttalha/deploy", name, version)))
	fmt.Fprintln(w)
}

// OK prints a green ✓ line.
func OK(w io.Writer, format string, a ...any) {
	fmt.Fprintln(w, successStyle.Render("✓"), fmt.Sprintf(format, a...))
}

// Step prints a cyan → line for ongoing/informational steps.
func Step(w io.Writer, format string, a ...any) {
	fmt.Fprintln(w, infoStyle.Render("→"), fmt.Sprintf(format, a...))
}

// Warn prints a yellow ⚠ line.
func Warn(w io.Writer, format string, a ...any) {
	fmt.Fprintln(w, warnStyle.Render("⚠"), fmt.Sprintf(format, a...))
}

// Fail prints a red ✗ line. If hint is non-empty, it follows on a dim line.
func Fail(w io.Writer, msg, hint string) {
	fmt.Fprintln(w, dangerStyle.Render("✗"), msg)
	if hint != "" {
		fmt.Fprintln(w, "  "+dimStyle.Render(hint))
	}
}

// KV prints a key/value pair with the key column padded to 10 chars.
func KV(w io.Writer, key, val string) {
	fmt.Fprintf(w, "  %s  %s\n", labelStyle.Render(padRight(key, 10)), val)
}

// Dim renders text in dim/faint style.
func Dim(s string) string { return dimStyle.Render(s) }

func padRight(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return s + strings.Repeat(" ", width-len(s))
}
