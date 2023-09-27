package linewrap

import (
	"unicode"

	"github.com/eprovst/wrap/pkg/ast"
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
		if (unicode.In(precC, unicode.Pd) && !unicode.In(thisC, unicode.Pd)) ||
			(unicode.IsSpace(thisC) && thisC != '\u00A0') { // U+00A0 is NBSP
			break
		}
	}

	// No break found, try a less ideal point
	if bidx == 0 {
		bidx = lineLength
		for ; bidx > 0; bidx-- {
			// Is second rate breakpoint?
			precC := lineText[bidx-1]
			thisC := lineText[bidx]
			if (unicode.IsLetter(precC) && !unicode.IsLetter(thisC)) ||
				(!unicode.IsLetter(precC) && unicode.IsLetter(thisC)) {
				break
			}
		}
	}

	// Still no break found, try an even less ideal point
	if bidx == 0 {
		bidx = lineLength + 1
		for ; bidx > 0 && bidx < len(lineText); bidx++ {
			// Is third rate breakpoint?
			precC := lineText[bidx-1]
			thisC := lineText[bidx]
			if (unicode.IsLetter(precC) && !unicode.IsLetter(thisC)) ||
				(!unicode.IsLetter(precC) && unicode.IsLetter(thisC)) {
				break
			}
		}
	}

	// Still no found: give up.
	if bidx >= len(lineText) {
		return line, ast.Line{}
	}

	// Now break the line at the found breakpoint
	headLength := 0
	head := []ast.Cell{}
	tail := []ast.Cell{}

	for currentCell := 0; currentCell < len(line); currentCell++ {
		cell := line[currentCell]
		cellContent := []rune(cell.Content)

		// If breakpoint in this cell: break cell
		if headLength+len(cellContent) > bidx {
			cidx := bidx - headLength

			fh := cell
			sh := cell

			fh.Content = string(removeSpaceAfter(cellContent[:cidx]))
			sh.Content = string(removeSpaceBefore(cellContent[cidx:]))

			head = append(head, fh)
			if len(sh.Content) == 0 {
				tail = line[currentCell+1:]
			} else {
				tail = append(ast.Line{sh}, line[currentCell+1:]...)
			}

			break
		}

		head = append(head, cell)
		headLength += len(cellContent)
	}

	return head, tail
}
