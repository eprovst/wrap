package pdf

import (
	"strings"

	"github.com/eprovst/wrap/pkg/ast"
)

// This file contains tools for theming.

type lineType byte

const (
	act lineType = iota
	centeredText
	action
	slugLine
	sceneNumber
	transition
	looseLyrics

	character
	parenthetical
	dialogue
	lyrics

	dualCharacterOne
	dualParentheticalOne
	dualDialogueOne
	dualLyricsOne

	dualCharacterTwo
	dualParentheticalTwo
	dualDialogueTwo
	dualLyricsTwo

	more

	titlePageTitle
	titlePageSubtitle
	titlePageCredit
	titlePageAuthor
	titlePageSource

	titlePageLeft
	titlePageRight
)

type aTheme map[lineType]lineStyle

type lineStyle struct {
	Indent          float64
	FirstLineOffset int
	LeadingBefore   int
	Leading         int
	LineLength      int
	AllCaps         bool
	Italics         bool
	Boldface        bool
	Underline       bool
	Centered        bool
	FlushRight      bool
}

var themeMap = map[string]aTheme{
	"screenplay": screenplay,
	"stageplay":  stageplay,
}

// Returns the leading in lines.
func (line styledLine) leading() int {
	// Get the theming for this line type.
	currentStyle := currentTheme[line.Type]

	if thisPDF.GetY() != topMargin {
		// First of kind only has different styling when explicitely set.
		if line.FirstOfSection {
			return currentStyle.LeadingBefore
		}
		return currentStyle.Leading
	}
	return 0
}

func styleLine(line styledLine) {
	// First handle the theming
	currentStyle := currentTheme[line.Type]

	// Positioning
	// No leading on top of page.
	thisPDF.Br(float64(line.leading()) * em)

	// Compute indentation
	indent := currentStyle.Indent
	if line.FirstOfSection {
		indent += float64(currentStyle.FirstLineOffset) * en
	}

	// Place the writehead:
	if currentStyle.Centered {
		thisPDF.SetX(leftMargin/2 + indent/2 + (pageWidth-rightMargin)/2 - float64(line.len())*en/2)
	} else if currentStyle.FlushRight {
		thisPDF.SetX(pageWidth - rightMargin - indent - float64(line.len())*en)
	} else {
		thisPDF.SetX(leftMargin + indent)
	}

	// Now write the cells.
	for _, cell := range line.Content {
		// Modify using theme
		if currentStyle.AllCaps {
			cell.Content = strings.ToUpper(cell.Content)
		}
		// We swap the current style.
		if currentStyle.Italics {
			cell.Italics = !cell.Italics
		}
		if currentStyle.Boldface {
			cell.Boldface = !cell.Boldface
		}
		if currentStyle.Underline {
			cell.Underline = !cell.Underline
		}
		addCell(cell)
	}

	// Next line
	thisPDF.Br(1 * em)
}

func addCell(cell ast.Cell) {
	if cell.Comment {
		return
	}

	fontStyle := ""

	if cell.Underline {
		fontStyle += "U"
	}
	if cell.Boldface {
		fontStyle += "B"
	}
	if cell.Italics {
		fontStyle += "I"
	}

	setStyledFont(fontStyle)

	thisPDF.Cell(nil, cell.Content)

	setDefaultFont()
}
