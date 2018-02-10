package parser

import "strings"

// Categorises the line and does some cleaning
func categoriser(line string) categorisedLine {
	lineType := getLineType(line)
	isForced := false

	switch lineType {
	case emptyLine:
		line = ""

	// beginAct is handled later on
	case forcedBeginAct:
		line = normaliseLine(line)
		line = strings.TrimPrefix(line, "%")
		line = strings.TrimSpace(line)
		isForced = true
		lineType = beginAct

	case centeredText:
		line = normaliseLine(line)
		line = strings.TrimPrefix(line, ">")
		line = strings.TrimSuffix(line, "<")
		line = strings.TrimSpace(line)

	// character is handled later on
	case forcedCharacter:
		line = normaliseLine(line)
		line = strings.TrimPrefix(line, "@")
		line = strings.TrimSpace(line)
		isForced = true
		lineType = character

	// characterTwo is handled later on
	case forcedCharacterTwo:
		line = normaliseLine(line)
		line = strings.TrimPrefix(line, "@")
		line = strings.TrimSuffix(line, "^")
		line = strings.TrimSpace(line)
		isForced = true
		lineType = characterTwo

	// endAct is handled later on
	case forcedEndAct:
		line = normaliseLine(line)
		line = strings.TrimPrefix(line, "%!")
		line = strings.TrimSpace(line)
		isForced = true
		lineType = endAct

	case forcedAction:
		line = normaliseLine(line)
		line = strings.TrimPrefix(line, "!")
		isForced = true
		lineType = action

	case lyrics:
		line = normaliseLine(line)
		line = strings.TrimPrefix(line, "~")
		line = strings.TrimSpace(line)

	case pageBreak:
		// No need for content
		line = ""

	case parenthetical:
		// A parenthetical is detected by the dialogue grouper.
		lineType = other

	// sceneTag is handled later on
	case forcedSceneTag:
		line = normaliseLine(line)
		line = strings.TrimPrefix(line, ".")
		line = strings.TrimSpace(line)
		isForced = true
		lineType = sceneTag

	// transitionTag is handled later on
	case forcedTransitionTag:
		line = normaliseLine(line)
		line = strings.TrimPrefix(line, ">")
		line = strings.TrimSpace(line)
		isForced = true
		lineType = transitionTag

	case section:
		line = normaliseLine(line)

	case synopse:
		line = normaliseLine(line)
		line = strings.TrimPrefix(line, "=")
		line = strings.TrimSpace(line)

		// Else do nothing, the line needs no cleaning.
	}

	return categorisedLine{lineType, isForced, line}
}
