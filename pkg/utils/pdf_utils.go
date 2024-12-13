package utils

import (
	"bytes"

	"github.com/ledongthuc/pdf"
)

// ExtractTextFromPDF extracts text from a PDF file
func ExtractTextFromPDF(file string) (text []string, err error) {
	f, r, err := pdf.Open(file)
	if err != nil {
		LogErrorAndExit(err)
	}

	defer f.Close()

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		LogErrorAndExit(err)
	}

	buf.ReadFrom(b)
	text = append(text, buf.String())
	return text, nil
}

// RemoveCommonWords removes common words from a list of keywords
func RemoveCommonWords(keywords []string) (wordList []string) {
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

	var result []string

	for _, word := range keywords {
		if _, exists := commonWords[word]; !exists {
			result = append(result, word)
		}
	}

	return result
}
