package utils

import (
	"os"

	"github.com/arsham/figurine/figurine"
	"github.com/jwalton/go-supportscolor"
)

// PrintStyledText prints a styled text to the terminal
func PrintStyledText(text string) error {
	// Check if the terminal supports colors
	if supportscolor.Stdout().SupportsColor {
		return figurine.Write(os.Stdout, text, "ANSI Regular.flf")
	}
	return nil
}
