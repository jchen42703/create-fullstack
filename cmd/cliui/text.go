package cliui

import (
	"fmt"

	"github.com/cli/cli/v2/pkg/iostreams"
)

// Handles printing using iostreams.ColorScheme()
type ColorUi struct {
	IoStreams   *iostreams.IOStreams
	colorScheme *iostreams.ColorScheme
}

func NewColorUi() *ColorUi {
	ioStreams := iostreams.System()

	return &ColorUi{
		IoStreams:   ioStreams,
		colorScheme: ioStreams.ColorScheme(),
	}
}

func (ui *ColorUi) Error(t string) {
	fmt.Fprint(ui.IoStreams.ErrOut, ui.colorScheme.Red(t))
}

func (ui *ColorUi) Errorf(t string, args ...interface{}) {
	fmt.Fprint(ui.IoStreams.ErrOut, ui.colorScheme.Redf(t, args...))
}

func (ui *ColorUi) Warn(t string) {
	fmt.Fprint(ui.IoStreams.ErrOut, ui.colorScheme.Yellow(t))
}

func (ui *ColorUi) Warnf(t string, args ...interface{}) {
	fmt.Fprint(ui.IoStreams.ErrOut, ui.colorScheme.Yellowf(t, args...))
}

func (ui *ColorUi) Gray(t string) {
	fmt.Fprint(ui.IoStreams.ErrOut, ui.colorScheme.Gray(t))
}

func (ui *ColorUi) Grayln(t string) {
	fmt.Fprintln(ui.IoStreams.Out, ui.colorScheme.Gray(t))
}

func (ui *ColorUi) Grayf(t string, args ...interface{}) {
	fmt.Fprint(ui.IoStreams.Out, ui.colorScheme.Grayf(t, args...))
}

func (ui *ColorUi) Log(t string) {
	fmt.Fprint(ui.IoStreams.Out, t)
}

func (ui *ColorUi) Logf(t string, args ...interface{}) {
	fmt.Fprintf(ui.IoStreams.Out, t, args...)
}

func (ui *ColorUi) Bold(t string) {
	fmt.Fprint(ui.IoStreams.Out, ui.colorScheme.Bold(t))
}

func (ui *ColorUi) Success(t string) {
	fmt.Fprint(ui.IoStreams.Out, ui.colorScheme.SuccessIcon()+" "+ui.colorScheme.Green(t))
}

func (ui *ColorUi) Successf(t string, args ...interface{}) {
	fmt.Fprint(ui.IoStreams.Out, ui.colorScheme.SuccessIcon()+" "+ui.colorScheme.Greenf(t, args...))
}
