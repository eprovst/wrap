package linewrap

import (
	"unicode"

	"github.com/Wraparound/wrap/pkg/ast"
)

// WrapLine breaks line into lines of correct length.
func WrapLine(line ast.Line, lineLength int) []ast.Line {
	head, tail := WrapLineOnce(line, lineLength)
	lines := []ast.Line{head}
	for len(tail) > 0 {
		head, tail = WrapLineOnce(tail, lineLength)
		lines = append(lines, head)
	}

	return lines
}

// WrapLineOnce breaks line into a line of correct length and the remainder.
func WrapLineOnce(line ast.Line, lineLength int) (ast.Line, ast.Line) {
	headLength := 0
	head := []ast.Cell{}

	for currentCell := 0; currentCell < len(line); currentCell++ {
		cell := line[currentCell]
		headLength += cell.Lenght()

		if headLength <= lineLength {
			head = append(head, cell)

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
				head = head[:len(head)-1]

				overflow = 1
				cellContent = []rune(cell.Content)

			} else {
				overflow = headLength - lineLength

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

			head = append(head, firstHalfOfCell)

			// Place the second half at the current position so it is reexamined, if not empty
			if secondHalfOfCell.Content != "" {
				line[currentCell] = secondHalfOfCell
				// This cell is still part of the tail
				currentCell--
			}

			return head, line[currentCell+1:]
		}
	}

	// Did not need to wrap line
	return line, ast.Line{}
}
