package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/arsham/figurine/figurine"
	"github.com/briandowns/spinner"
	"github.com/jwalton/go-supportscolor"
	"github.com/mattn/go-colorable"
)

const (
	ColorReset = "\033[0m"
	ColorGreen = "\033[32m"
	ColorBold  = "\033[1m"
)

// PrintStyledText prints a styled text to the terminal
func PrintStyledText(text string) error {
	// Check if the terminal supports colors
	if supportscolor.Stdout().SupportsColor {
		return figurine.Write(os.Stdout, text, "ANSI Regular.flf")
	}
	return nil
}

// StartSpinner prints a spinner to the terminal
func StartSpinner(message string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Color("magenta")
	s.Writer = colorable.NewColorableStdout() // Ensure colors are supported on Windows
	s.Suffix = " " + message
	fmt.Printf("%s%s%s ", ColorBold+ColorGreen, s.Suffix, ColorReset)
	s.Start()
	return s
}

// StopSpinner stops the spinner
func StopSpinner(s *spinner.Spinner) {
	s.Stop()
}
