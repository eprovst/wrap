package pdf

import (
	"math"

	"github.com/Wraparound/wrap/ast"
)

type aSection struct {
	Lines    []aLine
	Type     ast.ElementType
	Metadata map[string]interface{}
}

func (section aSection) height() int {
	height := 0
	for _, line := range section.Lines {
		height += line.leading() + 1
	}
	return height
}

func sectionize(element ast.Element) aSection {
	var lines []aLine
	metadata := map[string]interface{}{}

	switch element.(type) {
	case ast.Action:
		lines = append(lines, cellify(string(element.(ast.Action)), action)...)

	case ast.BeginAct:
		// Begin of act is on a new page.
		lines = append(lines, cellify(string(element.(ast.BeginAct)), act)...)

	case ast.EndAct:
		lines = append(lines, cellify(string(element.(ast.EndAct)), act)...)

	case ast.CenteredText:
		lines = append(lines, cellify(string(element.(ast.CenteredText)), centeredText)...)

	case ast.Lyrics:
		lines = append(lines, cellify(string(element.(ast.Lyrics)), looseLyrics)...)

	case ast.Scene:
		metadata["scenenumber"] = aCell{Content: element.(ast.Scene).SceneNumber}
		lines = append(lines, cellify(element.(ast.Scene).Slugline, slugLine)...)

	case ast.Transition:
		lines = append(lines, cellify(string(element.(ast.Transition)), transition)...)

	case ast.Dialogue:
		lines = append(lines, cellify(element.(ast.Dialogue).Character, character)...)
		for _, elem := range element.(ast.Dialogue).Lines {
			switch elem.(type) {
			case ast.Parenthetical:
				lines = append(lines, cellify(string(elem.(ast.Parenthetical)), parenthetical)...)

			case ast.Speech:
				lines = append(lines, cellify(string(elem.(ast.Speech)), dialogue)...)

			case ast.Lyrics:
				lines = append(lines, cellify(string(elem.(ast.Lyrics)), lyrics)...)
			}
		}

	case ast.DualDialogue:
		// First add the characters:
		lines = append(lines, cellify(element.(ast.DualDialogue).LCharacter, dualCharacterOne)...)
		lines = append(lines, cellify(element.(ast.DualDialogue).RCharacter, dualCharacterTwo)...)

		// Now add the lines in a intertwined fashion.
		leftDialogue := element.(ast.DualDialogue).LLines
		rightDialogue := element.(ast.DualDialogue).RLines
		for i := 0; i < int(math.Max(float64(len(leftDialogue)), float64(len(rightDialogue)))); i++ {
			if i < len(leftDialogue) {
				switch leftDialogue[i].(type) {
				case ast.Parenthetical:
					lines = append(lines, cellify(string(leftDialogue[i].(ast.Parenthetical)), dualParentheticalOne)...)

				case ast.Speech:
					lines = append(lines, cellify(string(leftDialogue[i].(ast.Speech)), dualDialogueOne)...)

				case ast.Lyrics:
					lines = append(lines, cellify(string(leftDialogue[i].(ast.Lyrics)), dualLyricsOne)...)
				}
			}

			if i < len(rightDialogue) {
				switch rightDialogue[i].(type) {
				case ast.Parenthetical:
					lines = append(lines, cellify(string(rightDialogue[i].(ast.Parenthetical)), dualParentheticalTwo)...)

				case ast.Speech:
					lines = append(lines, cellify(string(rightDialogue[i].(ast.Speech)), dualDialogueTwo)...)

				case ast.Lyrics:
					lines = append(lines, cellify(string(rightDialogue[i].(ast.Lyrics)), dualLyricsTwo)...)
				}
			}
		}
	}

	return aSection{
		Lines:    lines,
		Type:     ast.GetElementType(element),
		Metadata: metadata,
	}
}
