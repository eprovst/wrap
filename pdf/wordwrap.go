package pdf

import (
	"unicode"

	"github.com/Wraparound/wrap/ast"
)

// Breaks line into lines of correct lenght.
func wordwrap(line styledLine) []styledLine {
	lineLenght := currentTheme[line.Type].LineLenght

	// If the line lenght is undefined, change it to a default.
	if lineLenght == 0 {
		lineLenght = maxNumOfChars
	}

	currentLineLenght := 0
	currentLineContent := []ast.Cell{}
	lines := []styledLine{}

	for currentCell := 0; currentCell < len(line.Content); currentCell++ {
		cell := line.Content[currentCell]
		currentLineLenght += cell.Lenght()

		if currentLineLenght <= lineLenght {
			currentLineContent = append(currentLineContent, cell)

		} else {
			overflow := currentLineLenght - lineLenght
			breakoffset := overflow
			// We use a rune slice to be able to find 'nonbreakingspaces'.
			cellContent := []rune(cell.Content)

			// Now split the cell:
			// Let's look for possible break points:
			j := overflow
			for foundbreak := false; !foundbreak &&
				j < len(cellContent); j++ {

				// Is potential breakpoint?
				if unicode.IsSpace(cellContent[len(cellContent)-j]) ||
					cellContent[len(cellContent)-j-1] == '-' {

					breakoffset = j
					foundbreak = true
				}
			}

			// Now break the line:
			firstHalfOfCell := cell
			hyphenInsterted := false
			if unicode.IsLetter(cellContent[len(cellContent)-breakoffset-1]) &&
				unicode.IsLetter(cellContent[len(cellContent)-breakoffset]) {

				hyphenInsterted = true
				firstHalfOfCell.Content = string(cellContent[:len(cellContent)-breakoffset-1]) + "-"

			} else {
				firstHalfOfCell.Content = removeSpaceAfter(string(cellContent[:len(cellContent)-breakoffset]))
			}

			secondHalfOfCell := cell

			if hyphenInsterted {
				secondHalfOfCell.Content = removeSpaceBefore(string(cellContent[len(cellContent)-breakoffset-1:]))

			} else {
				secondHalfOfCell.Content = removeSpaceBefore(string(cellContent[len(cellContent)-breakoffset:]))
			}

			// Add the first half to the current line and that line to the final list
			currentLineContent = append(currentLineContent, firstHalfOfCell)

			lines = append(lines, styledLine{
				Content: currentLineContent,
				Type:    line.Type,
			})

			// Prepare next line aggregation
			currentLineLenght = 0
			currentLineContent = []ast.Cell{}

			// Place the second half at the current position so it is reexamined, if not empty
			if secondHalfOfCell.Content != "" {
				line.Content[currentCell] = secondHalfOfCell
				currentCell-- // Reexamine
			}
		}
	}

	// Add last line:
	lines = append(lines, styledLine{
		Content: currentLineContent,
		Type:    line.Type,
	})

	return lines
}
