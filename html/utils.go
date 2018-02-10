package html

import (
	"bytes"
	"strings"
)

// Removes tag(s) and it/their content
func removeHTMLTags(s string) string {
	sftag := strings.Split(s, "<")
	var slist []string
	// Now split at the closing bracket.
	for _, part := range sftag {
		slist = append(slist, strings.Split(part, ">")...)
	}

	// The first string is content.
	result := bytes.NewBufferString(slist[0])
	// Every even index is not a tag (so we start at two).
	for i := 2; i < len(slist); i += 2 {
		result.WriteString(slist[i])
	}

	return result.String()
}
