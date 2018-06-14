package ast

import (
	"strings"
	"unicode/utf8"

	"github.com/Wraparound/wrap/languages"
)

/* This file defines all the types used by the parser
to generate a syntax tree. */

/* Rich text */

// Line a line of richtext cells
type Line []Cell

// Lenght is the amount of runes in the line
func (line Line) Lenght() int {
	lineLength := 0

	for _, cell := range line {
		lineLength += cell.Lenght()
	}

	return lineLength
}

// Empty checks if the line is empty
func (line Line) Empty() bool {
	for _, cell := range line {
		if len(strings.TrimSpace(cell.Content)) > 0 {
			return false
		}
	}

	return true
}

// String returns a plaintext version of the line
func (line Line) String() string {
	builder := strings.Builder{}

	for _, cell := range line {
		builder.WriteString(cell.Content)
	}

	return builder.String()
}

// Cell contains richtext
type Cell struct {
	Content   string
	Boldface  bool
	Italics   bool
	Underline bool
	Comment   bool
}

// Lenght of a Cell is the amount of runes in it
func (cell Cell) Lenght() int {
	return utf8.RuneCountInString(cell.Content)
}

// Empty checks if the cell has no contents
func (cell Cell) Empty() bool {
	return cell.Content == ""
}

/* Script sections */

// Script contains the entire script.
type Script struct {
	Language  languages.Language
	TitlePage map[string][]Line
	Elements  []Element
}

// Element represents a part of the script.
type Element interface{}

// Scene header type.
type Scene struct {
	Slugline    []Line
	SceneNumber string
}

// BeginAct type
type BeginAct []Line

// EndAct type
type EndAct []Line

// Action type.
type Action []Line

// Dialogue type.
type Dialogue struct {
	Character []Line
	Lines     []Element
}

// DualDialogue type.
type DualDialogue struct {
	LCharacter []Line
	LLines     []Element

	RCharacter []Line
	RLines     []Element
}

// The following three can be used with Dialogue, Lyrics can also be standalone.

// Parenthetical type
type Parenthetical []Line

// Speech type
type Speech []Line

// Lyrics type.
type Lyrics []Line

// Transition type.
type Transition []Line

// CenteredText type.
type CenteredText []Line

// PageBreak type.
type PageBreak struct{}

// Section type.
type Section struct {
	Level uint8
	Line  []Line
}

// Synopse type.
type Synopse []Line

// Note represents a note which isn't part of another element.
type Note []Line
