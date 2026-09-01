package utils

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/ledongthuc/pdf"
)

// isValidWord checks if a word is meaningful and should be kept
func isValidWord(word string) bool {
	// Trim non-letter characters
	word = strings.TrimFunc(word, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	// Reject words that are too short or invalid
	if utf8.RuneCountInString(word) < 3 {
		return false
	}

	// Check the ratio of letters to other characters
	letters, numbers := 0, 0
	for _, r := range word {
		switch {
		case unicode.IsLetter(r):
			letters++
		case unicode.IsNumber(r):
			numbers++
		}
	}

	// Require at least 2 letters and <50% numeric characters
	return letters >= 2 && float64(numbers)/float64(len(word)) < 0.5
}

// ExtractTextFromPDF extracts text from a PDF file and splits it into cleaned words
func ExtractTextFromPDF(file string) ([]string, error) {
	f, r, err := pdf.Open(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open PDF file: %w", err)
	}
	defer f.Close()

	var buffer bytes.Buffer

	// Extract text from each page
	for pageIndex := 0; pageIndex < r.NumPage(); pageIndex++ {
		page := r.Page(pageIndex)
		if page.V.IsNull() {
			continue
		}
		buffer.WriteString(fmt.Sprintf("%v", page.Content()))
	}

	// Unicode normalization for consistent encoding
	text := strings.ToValidUTF8(buffer.String(), "")

	// Remove non-alphanumeric characters except spaces and dashes
	reg := regexp.MustCompile(`[^\w\s-]`)
	cleanedText := reg.ReplaceAllString(text, " ")

	// Split text into words
	words := strings.Fields(cleanedText)
	var validWords []string

	// Validate words
	for _, word := range words {
		if isValidWord(word) {
			validWords = append(validWords, word)
		}
	}

	return validWords, nil
}

// RemoveCommonWords filters out common words and deduplicates the input
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

	seenWords := make(map[string]struct{})
	var result []string

	// Filter out common words and duplicates
	for _, word := range keywords {
		lowerWord := strings.ToLower(word)
		if _, isCommon := commonWords[lowerWord]; !isCommon {
			if _, seen := seenWords[lowerWord]; !seen {
				seenWords[lowerWord] = struct{}{}
				result = append(result, lowerWord)
			}
		}
	}

	return result
}

// cleanToken applies additional token cleanup rules
func cleanToken(token string) string {
	// Remove hexadecimal prefixes (e.g., "0x1234")
	token = regexp.MustCompile(`^(0x|0\d+)`).ReplaceAllString(token, "")

	// Trim non-letter characters
	token = strings.TrimFunc(token, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	return token
}
