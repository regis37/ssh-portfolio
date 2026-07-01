package ui

import (
	"bytes"
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

// forcedRenderer returns a lipgloss renderer in TrueColor mode so that
// ANSI color codes are always emitted regardless of the test environment.
// WithUnsafe() is required to suppress the isTTY check that would otherwise
// strip all ANSI codes when writing to a bytes.Buffer (non-TTY).
func forcedRenderer() *lipgloss.Renderer {
	return lipgloss.NewRenderer(
		&bytes.Buffer{},
		termenv.WithProfile(termenv.TrueColor),
		termenv.WithUnsafe(),
	)
}

func TestThemeToggleProducesDifferentOutput(t *testing.T) {
	r := forcedRenderer()

	dark := New(r)
	dark.width = 120
	dark.height = 30

	light := New(r)
	light.width = 120
	light.height = 30
	light.toggleTheme() // switch to light

	darkOut := dark.View()
	lightOut := light.View()

	// Both outputs must contain ANSI escape codes.
	if !strings.Contains(darkOut, "\x1b[") {
		t.Error("dark theme: no ANSI codes in output — renderer is in Ascii/no-color mode")
	}
	if !strings.Contains(lightOut, "\x1b[") {
		t.Error("light theme: no ANSI codes in output — renderer is in Ascii/no-color mode")
	}

	// The outputs must differ (different fg/bg color codes).
	if darkOut == lightOut {
		t.Fatal("FAIL: dark and light View() outputs are IDENTICAL — theme toggle has no effect on render")
	}

	t.Log("PASS: dark and light View() outputs differ")

	// Show first difference for diagnosis.
	for i := 0; i < len(darkOut) && i < len(lightOut); i++ {
		if darkOut[i] != lightOut[i] {
			lo := max(0, i-10)
			hiD := min(i+60, len(darkOut))
			hiL := min(i+60, len(lightOut))
			t.Logf("First diff at byte %d:", i)
			t.Logf("  dark:  %q", darkOut[lo:hiD])
			t.Logf("  light: %q", lightOut[lo:hiL])
			break
		}
	}
}
