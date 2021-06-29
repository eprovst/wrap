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
	// TODO: Allow break with a hyphen on relevant marker.
	// TODO: Ignore zero width characters.

	lineText := []rune(line.String())

	// If the line is short enough, return
	if len(lineText) <= lineLength {
		return line, ast.Line{}
	}

	// Else look for possible break point
	bidx := lineLength // bidx is the first index to be included in the next line.
	for ; bidx > 0; bidx-- {
		// Is potential breakpoint?
		precC := lineText[bidx-1]
		thisC := lineText[bidx]
		if (unicode.IsSpace(thisC) && thisC != '\u00A0') || // U+00A0 is NBSP
			(unicode.IsLetter(precC) && !unicode.IsLetter(thisC) && !unicode.IsPunct(thisC)) ||
			(!unicode.IsLetter(precC) && unicode.IsLetter(thisC)) {
			break
		}
	}

	// No break found, TODO: at least try something?
	if bidx == 0 {
		return line, ast.Line{}
	}

	// Now break the line at the found breakpoint
	headLength := 0
	head := []ast.Cell{}
	tail := []ast.Cell{}

	for currentCell := 0; currentCell < len(line); currentCell++ {
		cell := line[currentCell]

		// If breakpoint in this cell: break cell
		if headLength+cell.Lenght() > bidx {
			cidx := bidx - headLength

			fh := cell
			sh := cell

			fh.Content = removeSpaceAfter(string(cell.Content[:cidx]))
			sh.Content = removeSpaceBefore(string(cell.Content[cidx:]))

			head = append(head, fh)
			tail = append(ast.Line{sh}, line[currentCell+1:]...)

			break
		}

		head = append(head, cell)
		headLength += cell.Lenght()
	}

	return head, tail
}
