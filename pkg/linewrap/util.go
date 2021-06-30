package linewrap

import (
	"unicode"
)

func removeSpaceBefore(s []rune) []rune {
	// Can't remove from something empty...
	if len(s) == 0 {
		return []rune{}
	}

	i := 0
	currentlyspace := true
	for ; i < len(s) && currentlyspace; i++ {
		currentlyspace = unicode.IsSpace(s[i])
	}

	// It's all spaces
	if currentlyspace {
		return []rune{}
	}

	// Now i is the point after where the first nonspace is, thus
	return s[i-1:]
}

func removeSpaceAfter(s []rune) []rune {
	// Can't remove from something empty...
	if len(s) == 0 {
		return s
	}

	s = reverse(s)
	s = removeSpaceBefore(s)
	s = reverse(s)
	return s
}

func reverse(s []rune) []rune {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
