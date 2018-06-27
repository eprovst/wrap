package parser

import (
	"strings"
	"unicode"
)

/* Here we have functions to identify the kind of line we're dealing with
these also implement the translations in the Wrap specification. */

// UseWrapExtensions is defined in parser.go
// Translation specific items are defined in language.go

/* IMPORTANT: THE is* FUNCTIONS ARE RATHER ORDER DEPENDENT */

// Utils
func normaliseLine(line string) string {
	return strings.TrimSpace(line)
}

// TitlePage identification is fully embedded inside parser.go

func isForcedAction(line string) bool {
	line = normaliseLine(line)
	return strings.HasPrefix(line, "!")
}

func isEmptyLine(line string) bool {
	// No normalisation as spaces are important for identification...
	// An empty line contains no or one space
	return line == "" || line == " "
}

func isSceneTag(line string) bool {
	line = normaliseLine(line)

	// Case doesn't matter so make everything lower case.
	if isForcedAction(line) || isForcedSceneTag(line) {
		return false

	} else if hasCaseInsensitivePrefixInSlice(line, translation.SceneTags) {
		return true

	} else if UseWrapExtensions && hasCaseInsensitivePrefixInSlice(line, translation.StageplaySceneTags) {
		// Only supported in Wrap
		return true
	}

	return false
}

func isForcedSceneTag(line string) bool {
	line = normaliseLine(line)

	if strings.HasPrefix(line, ".") {
		secChar := rune(line[1])

		if unicode.IsLetter(secChar) || unicode.IsNumber(secChar) {
			return true
		}
	}

	return false
}

func isTransitionTag(line string) bool {
	line = normaliseLine(line)

	if isForcedAction(line) || isForcedTransitionTag(line) {
		return false

	} else if hasSuffixInSlice(line, translation.TransitionTags) {
		return true
	}

	return false
}

func isForcedTransitionTag(line string) bool {
	line = normaliseLine(line)
	return strings.HasPrefix(line, ">") && !strings.HasSuffix("<", line)
}

func isBeginAct(line string) bool {
	line = normaliseLine(line)

	// This element only exists in Wrap
	if isForcedAction(line) || !UseWrapExtensions || isForcedBeginAct(line) {
		return false

	} else if hasPrefixInSlice(line, translation.BeginActTags) {
		isUppercase(line)
	}

	return false
}

func isForcedBeginAct(line string) bool {
	line = normaliseLine(line)
	return strings.HasPrefix(line, "%") && !strings.HasPrefix(line, "%!") && UseWrapExtensions
}

func isEndAct(line string) bool {
	line = normaliseLine(line)

	// This element only exists in Wrap
	if isForcedAction(line) || !UseWrapExtensions || isForcedEndAct(line) {
		return false

	} else if hasPrefixInSlice(line, translation.EndActTags) {
		return isUppercase(line)
	}

	return false
}

func isForcedEndAct(line string) bool {
	line = normaliseLine(line)
	return strings.HasPrefix(line, "%!") && UseWrapExtensions
}

func isUppercaseIgnoringCharacterExtensions(line string) bool {
	// NOTE: We do not check if the brackets actually close
	// we would need a second loop for that and the spec doesn't
	// say the brackets *have* to be balanced...

	inCharacterExtension := false

	for _, c := range []rune(line) {
		if c == '(' {
			inCharacterExtension = true

		} else if c == ')' {
			inCharacterExtension = false

		} else if !inCharacterExtension && unicode.IsLower(c) {
			return false
		}
	}

	return true
}

func isCharacter(line string) bool {
	line = normaliseLine(line)

	if strings.HasSuffix(line, "^") || isForcedCharacter(line) {
		// Might be a characterTwo or is a forced character
		return false

	} else if UseWrapExtensions {
		// In Wrap a character name may also start with Mc or Mac.
		line = strings.TrimPrefix(line, "Mc")
		line = strings.TrimPrefix(line, "Mac")
	}

	return isUppercaseIgnoringCharacterExtensions(line) && containsLetter(line)
}

func isForcedCharacter(line string) bool {
	line = normaliseLine(line)
	return strings.HasPrefix(line, "@") && !strings.HasSuffix(line, "^")
}

func isCharacterTwo(line string) bool {
	line = normaliseLine(line)

	if strings.HasSuffix(line, "^") {
		// Without the ^ at the end it should be a normal character.
		line = strings.TrimSuffix(line, "^")
		return isCharacter(line)
	}

	return false
}

func isForcedCharacterTwo(line string) bool {
	line = normaliseLine(line)
	return strings.HasPrefix(line, "@") && strings.HasSuffix(line, "^")
}

func isParenthetical(line string) bool {
	line = normaliseLine(line)

	// We do not have to check for a forced action as this wouldn't start with '('
	return strings.HasPrefix(line, "(") && strings.HasSuffix(line, ")")
}

func isLyrics(line string) bool {
	line = normaliseLine(line)

	// We do not have to check for a forced action as this wouldn't start with '~'
	return strings.HasPrefix(line, "~")
}

func isCenteredText(line string) bool {
	// NOTE: Line cleanup is for parser.go, thus removing space within '> ... <' too.
	line = normaliseLine(line)

	// We do not have to check for a forced normal action as this wouldn't start with '>'
	return strings.HasPrefix(line, ">") && strings.HasSuffix(line, "<")
}

func isNoteOnOwnLine(line string) bool {
	line = normaliseLine(line)

	// We do not have to check for a forced normal action as this wouldn't start with '[['
	return strings.HasPrefix(line, "[[") && strings.HasSuffix(line, "]]")
}

func isPageBreak(line string) bool {
	line = normaliseLine(line)

	if len(line) < 3 {
		// We need at least three '='s
		return false
	}

	// We do not have to check for a forced normal action as this wouldn't start with '='
	for _, c := range line {
		if c != '=' {
			return false
		}
	}

	return true
}

// NOTE: boneyard and notes are for parser.go and texthandler.go respectively.

func isSection(line string) bool {
	line = normaliseLine(line)

	// We do not have to check for a forced normal action as this wouldn't start with '#'
	return strings.HasPrefix(line, "#")
}

func isSynopse(line string) bool {
	line = normaliseLine(line)

	// We do not have to check for a forced normal action as this wouldn't start with '='
	if strings.HasPrefix(line, "=") && !isPageBreak(line) {
		//           it shouldn't be a pagebreak ^
		return true
	}

	return false
}

// The Parser() will check if it is end/start/inline etc.

func isBoneyard(line string) bool {
	return strings.Contains(line, "/*") || strings.Contains(line, "*/")
}

/* NOW COMBINE */

type lineCat byte

const (
	emptyLine lineCat = iota
	other
	beginAct
	centeredText
	character
	characterTwo
	endAct
	forcedAction
	forcedBeginAct
	forcedCharacter
	forcedCharacterTwo
	forcedEndAct
	forcedSceneTag
	forcedTransitionTag
	noteOnOwnLine
	boneyard
	lyrics
	pageBreak
	parenthetical
	sceneTag
	section
	synopse
	transitionTag
	action = other
)

/* Gives the type of line in lowercase, if no test succesfull:
"other", note: a forced action -> forcedAction */

// THE ORDER IS IMPORTANT!
func getLineType(line string) lineCat {
	if isEmptyLine(line) {
		return emptyLine

	} else if isBoneyard(line) {
		return boneyard

	} else if isCenteredText(line) {
		return centeredText

	} else if isNoteOnOwnLine(line) {
		return noteOnOwnLine

	} else if isForcedAction(line) {
		return forcedAction

	} else if isForcedBeginAct(line) {
		return forcedBeginAct

	} else if isForcedCharacter(line) {
		return forcedCharacter

	} else if isForcedCharacterTwo(line) {
		return forcedCharacterTwo

	} else if isForcedEndAct(line) {
		return forcedEndAct

	} else if isForcedSceneTag(line) {
		return forcedSceneTag

	} else if isForcedTransitionTag(line) {
		return forcedTransitionTag

	} else if isPageBreak(line) {
		return pageBreak

	} else if isSection(line) {
		return section

	} else if isSynopse(line) {
		return synopse

	} else if isLyrics(line) {
		return lyrics

	} else if isBeginAct(line) {
		return beginAct

	} else if isEndAct(line) {
		return endAct

	} else if isSceneTag(line) {
		return sceneTag

	} else if isTransitionTag(line) {
		return transitionTag

	} else if isParenthetical(line) {
		return parenthetical

	} else if isCharacterTwo(line) {
		return characterTwo

	} else if isCharacter(line) {
		return character
	}

	return other // == action
}

func isIgnoredLineType(lineType lineCat) bool {
	switch lineType {
	case boneyard, synopse, section, pageBreak, emptyLine:
		return true

	default:
		return false
	}
}
