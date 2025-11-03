package main

import (
	"regexp"
	"strings"
)

func WordFrequency(text string) map[string]int {
	text = strings.ToLower(text)
	re := regexp.MustCompile(`[^\w\s]`)
	text = re.ReplaceAllString(text, "")
	words := strings.Fields(text)
	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}
	return freq
}
