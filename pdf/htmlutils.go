package pdf

import (
	"bytes"
	"strings"
)

// Removes tag(s) and it/their content
func removeBetweenTags(s, tag string) string {
	sftag := strings.Split(s, "<"+tag+">")
	var slist []string
	// Now split at the closing tag.
	for _, part := range sftag {
		slist = append(slist, strings.Split(part, "</"+tag+">")...)
	}

	// The first string isn't inside a tag.
	result := bytes.NewBufferString(slist[0])
	// Every uneven index is inside the tag (so we start at two).
	for i := 2; i < len(slist); i += 2 {
		result.WriteString(slist[i])
	}

	return result.String()
}

// Removes only the tag(s) themselves.
func removeTags(s, tag string) string {
	s = strings.Replace(s, "<"+tag+">", "", -1)
	return strings.Replace(s, "</"+tag+">", "", -1)
}

// Remove all HTML tags
func removeAllTags(s string) string {
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
