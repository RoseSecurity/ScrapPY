package utils

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/barasher/go-exiftool"
	"github.com/olekukonko/tablewriter"
)

// ANSI escape codes for bold and blue text
var (
	blue = "\033[1;34m"
	norm = "\x1b[0m"
)

type wordCount struct {
	word  string
	count int
}

// WordFrequency returns the top N most frequently used keywords from the input slice of strings
func WordFrequency(input []string, topN int) []string {
	// Create a map to store word frequencies
	frequency := make(map[string]int)

	// Iterate over each string in the input slice
	for _, sentence := range input {
		// Split the sentence into words and normalize them to lowercase
		words := strings.Fields(strings.ToLower(sentence))
		for _, word := range words {
			// Increment the count for each word in the map
			frequency[word]++
		}
	}

	// Convert the map to a slice of wordCount structs for sorting
	counts := make([]wordCount, 0, len(frequency))
	for word, count := range frequency {
		counts = append(counts, wordCount{word, count})
	}

	// Sort the slice by count in descending order, then by word alphabetically
	sort.Slice(counts, func(i, j int) bool {
		if counts[i].count == counts[j].count {
			return counts[i].word < counts[j].word
		}
		return counts[i].count > counts[j].count
	})

	// Collect the top N words
	result := make([]string, 0, topN)
	for i := 0; i < topN && i < len(counts); i++ {
		result = append(result, counts[i].word)
	}

	return result
}

// PrintMetadata uses exiftool to print PDF metadata
func PrintMetadata(file string) {
	et, err := exiftool.NewExiftool()
	if err != nil {
		LogErrorAndExit(err)
		return
	}
	defer et.Close()

	fileInfos := et.ExtractMetadata(file)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Key", "Value"})
	table.SetBorder(false)

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			LogError(fileInfo.Err)
			continue
		}

		for k, v := range fileInfo.Fields {
			table.Append([]string{k, fmt.Sprintf("%v", v)})
		}
	}

	table.Render()
}
