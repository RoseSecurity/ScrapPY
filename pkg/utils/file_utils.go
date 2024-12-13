package utils

import (
	"fmt"
	"os"
)

// WriteToFile writes the keywords to a specified output file
func WriteToFile(filename string, keywords []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer file.Close()

	// Write keywords to file
	for _, keyword := range keywords {
		if _, err := file.WriteString(keyword + "\n"); err != nil {
			return fmt.Errorf("error writing to file: %v", err)
		}
	}

	return nil
}
