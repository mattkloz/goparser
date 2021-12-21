package parse

import (
	"strings"
)

func Chunks(s string, chunkSize int) []string {

	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}

func ParseItem(description string) []string {
	noSpaceString := strings.ReplaceAll(description, " ", "")

	a := []string{}

	s := noSpaceString
	n := 65000
	cs := Chunks(s, n)

	// If string is > 65000 chars, it needs to be chunked.
	for _, c := range cs {
		a = append(a, c)
	}
	return a
}
