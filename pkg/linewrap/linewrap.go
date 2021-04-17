package linewrap

import (
	"unicode"

	"github.com/Wraparound/wrap/pkg/ast"
)

// WrapLine breaks line into lines of correct length.
func WrapLine(line ast.Line, lineLength int) []ast.Line {
	currentLineLength := 0
	currentLineContent := []ast.Cell{}
	lines := []ast.Line{}

	for currentCell := 0; currentCell < len(line); currentCell++ {
		cell := line[currentCell]
		currentLineLength += cell.Lenght()

		if currentLineLength <= lineLength {
			currentLineContent = append(currentLineContent, cell)

		} else {
			// Select breaking parameters
			var (
				overflow    int
				cellContent []rune
			)

			if cell.Lenght() == 1 {
				// Very peculiar edge case, backtrack one cell and break that one
				currentCell--
				cell = line[currentCell]
				currentLineContent = currentLineContent[:len(currentLineContent)-1]

				overflow = 1
				cellContent = []rune(cell.Content)

			} else {
				overflow = currentLineLength - lineLength

				// We use a rune slice to be able to find 'nonbreakingspaces'.
				cellContent = []rune(cell.Content)
			}

			// Now split the cell:
			// Let's look for possible break points:
			breakoffset := overflow
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
			if len(cellContent)-breakoffset >= 1 &&
				unicode.IsLetter(cellContent[len(cellContent)-breakoffset-1]) &&
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

			lines = append(lines, currentLineContent)

			// Prepare next line aggregation
			currentLineLength = 0
			currentLineContent = []ast.Cell{}

			// Place the second half at the current position so it is reexamined, if not empty
			if secondHalfOfCell.Content != "" {
				line[currentCell] = secondHalfOfCell
				currentCell-- // Reexamine
			}
		}
	}

	// Add last line:
	lines = append(lines, currentLineContent)

	return lines
}
