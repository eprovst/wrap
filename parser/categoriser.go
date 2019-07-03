package parser

import "strings"

// Categorises the line and does some cleaning
func categoriser(line string) categorisedLine {
	lineType := getLineType(line)
	isForced := false

	switch lineType {
	case emptyLine:
		line = ""

	case beginAct:
		line = normaliseLine(line)

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

	case character:
		line = normaliseLine(line)

	case forcedCharacter:
		line = normaliseLine(line)
		line = strings.TrimPrefix(line, "@")
		line = strings.TrimSpace(line)
		isForced = true
		lineType = character

	case characterTwo:
		line = normaliseLine(line)

	case forcedCharacterTwo:
		line = normaliseLine(line)
		line = strings.TrimPrefix(line, "@")
		line = strings.TrimSuffix(line, "^")
		line = strings.TrimSpace(line)
		isForced = true
		lineType = characterTwo

	case endAct:
		line = normaliseLine(line)

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
		line = normaliseLine(line)
		// A parenthetical is detected by the dialogue grouper.
		lineType = other

	case sceneTag:
		line = normaliseLine(line)

	case forcedSceneTag:
		line = normaliseLine(line)
		line = strings.TrimPrefix(line, ".")
		line = strings.TrimSpace(line)
		isForced = true
		lineType = sceneTag

	case transitionTag:
		line = normaliseLine(line)

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
