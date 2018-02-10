package parser

import (
	"strings"
	"unicode"
)

func hasPrefixInSlice(line string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(line, prefix) {
			return true
		}
	}

	return false
}

func hasSuffixInSlice(line string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(line, suffix) {
			return true
		}
	}

	return false
}

func containsLetter(line string) bool {
	for _, char := range line {
		if unicode.IsLetter(char) {
			return true
		}
	}

	return false
}
