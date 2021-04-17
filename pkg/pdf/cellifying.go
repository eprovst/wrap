package pdf

import (
	"strings"

	"github.com/Wraparound/wrap/pkg/ast"
	"github.com/Wraparound/wrap/pkg/linewrap"
)

type styledLine struct {
	Content        []ast.Cell
	Type           lineType
	FirstOfSection bool
}

func (line styledLine) len() int {
	lineLength := 0
	for _, cell := range line.Content {
		lineLength += cell.Lenght()
	}
	return lineLength
}

func (line styledLine) isEmpty() bool {
	for _, cell := range line.Content {
		if len(strings.TrimSpace(cell.Content)) > 0 {
			return false
		}
	}
	return true
}

/* Prepares a section for insertion into the PDF.
Linelength is the maximum linelenght in characters (we work with a monospace font). */
func cellify(text []ast.Line, style lineType) []styledLine {
	// Now sart building the line list.
	lines := []styledLine{}

	for _, line := range text {
		styledline := styledLine{
			Content: line,
			Type:    style,
		}

		lines = append(lines, wrapline(styledline)...)
	}

	// If the last break was only a break, then...
	if len(lines) > 0 && lines[len(lines)-1].isEmpty() {
		// ...remove the last line.
		lines = lines[:len(lines)-1]
	}

	// Lable the first line as first of section
	if len(lines) > 0 {
		lines[0].FirstOfSection = true
	}

	return lines
}

// Breaks line into lines of correct lenght.
func wrapline(line styledLine) []styledLine {
	lineType := line.Type
	lineLength := currentTheme[lineType].LineLength

	// If the line length is undefined or invalid, change it to a default.
	if lineLength <= 1 {
		lineLength = maxNumOfChars
	}

	lines := linewrap.WrapLine(line.Content, lineLength)

	styledLines := make([]styledLine, len(lines))
	for i, line := range lines {
		styledLines[i] = styledLine{
			Content: line,
			Type:    lineType,
		}
	}

	return styledLines
}
