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

func hasCaseInsensitivePrefixInSlice(line string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if hasCaseInsensitivePrefix(line, prefix) {
			return true
		}
	}

	return false
}

func hasCaseInsensitivePrefix(line string, prefix string) bool {
	return len(line) >= len(prefix) && strings.EqualFold(line[0:len(prefix)], prefix)
}

func hasSuffixInSlice(line string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(line, suffix) {
			return true
		}
	}

	return false
}

func isUppercase(line string) bool {
	for _, c := range []rune(line) {
		if unicode.IsLower(c) {
			return false
		}
	}

	return true
}

func containsLetter(line string) bool {
	for _, char := range line {
		if unicode.IsLetter(char) {
			return true
		}
	}

	return false
}
