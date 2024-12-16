package utils

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/ledongthuc/pdf"
)

// isValidWord checks if a word is meaningful and should be kept
func isValidWord(word string) bool {
	// Remove any leading/trailing non-letter characters
	word = strings.TrimFunc(word, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	// Reject if word is empty after trimming
	if len(word) < 3 {
		return false
	}

	// Reject if word contains too many numbers or special characters
	numCount := 0
	letterCount := 0
	for _, r := range word {
		if unicode.IsNumber(r) {
			numCount++
		}
		if unicode.IsLetter(r) {
			letterCount++
		}
	}

	// Require at least 2 letters and less than 50% numbers
	return letterCount >= 2 && float64(numCount)/float64(len(word)) < 0.5
}

// ExtractTextFromPDF extracts text from a PDF file and splits it into cleaned words
func ExtractTextFromPDF(file string) ([]string, error) {
	f, r, err := pdf.Open(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open PDF file: %w", err)
	}
	defer f.Close()

	// Get plain text from the PDF reader
	b, err := r.GetPlainText()
	if err != nil {
		return nil, fmt.Errorf("failed to extract plain text from PDF: %w", err)
	}

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(b); err != nil {
		return nil, fmt.Errorf("failed to read text from buffer: %w", err)
	}

	text := buf.String()

	// Comprehensive text cleaning
	// Remove special characters, preserve letters, numbers, and spaces
	reg := regexp.MustCompile(`[^a-zA-Z0-9\s-]`)
	cleanedText := reg.ReplaceAllString(text, " ")

	// Split into words and filter
	words := strings.Fields(cleanedText)
	var validWords []string
	for _, word := range words {
		// Additional filtering for meaningful words
		if isValidWord(word) {
			validWords = append(validWords, word)
		}
	}

	return validWords, nil
}

// RemoveCommonWords removes common words and deduplicates entries
func RemoveCommonWords(keywords []string) []string {
	commonWords := map[string]struct{}{
		"and": {}, "the": {}, "at": {}, "there": {}, "some": {}, "my": {}, "of": {}, "be": {},
		"use": {}, "her": {}, "than": {}, "this": {}, "an": {}, "would": {}, "first": {}, "a": {},
		"have": {}, "each": {}, "to": {}, "from": {}, "which": {}, "like": {}, "been": {}, "in": {},
		"or": {}, "she": {}, "him": {}, "is": {}, "one": {}, "do": {}, "into": {}, "who": {}, "you": {},
		"had": {}, "how": {}, "that": {}, "by": {}, "their": {}, "has": {}, "its": {}, "it": {}, "if": {},
		"he": {}, "but": {}, "was": {}, "not": {}, "up": {}, "more": {}, "for": {}, "are": {}, "were": {},
		"as": {}, "we": {}, "with": {}, "when": {}, "then": {}, "no": {}, "come": {}, "his": {}, "your": {},
		"them": {}, "way": {}, "they": {}, "can": {}, "these": {}, "could": {}, "may": {}, "I": {},
		"said": {}, "so": {},
	}

	seenWords := make(map[string]struct{}, len(keywords))
	var result []string

	for _, word := range keywords {
		lowerWord := strings.ToLower(word)

		// Check if word is not a common word
		if _, isCommon := commonWords[lowerWord]; !isCommon {
			// Check if word has not been seen before
			if _, isSeen := seenWords[lowerWord]; !isSeen {
				seenWords[lowerWord] = struct{}{} // Mark word as seen
				result = append(result, word)
			}
		}
	}

	return result
}

// Additional utility to clean problematic tokens
func cleanToken(token string) string {
	// Remove hexadecimal and numeric prefixes
	token = regexp.MustCompile(`^(0x|0\d+)`).ReplaceAllString(token, "")

	// Trim non-letter characters from start and end
	token = strings.TrimFunc(token, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	return token
}
