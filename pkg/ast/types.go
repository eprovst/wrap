package ast

import (
	"bytes"
	"strings"
	"unicode/utf8"

	"github.com/Wraparound/wrap/pkg/languages"
)

/* This file defines all the types used by the parser
to generate a syntax tree. */

// Check if two lists of Elements are equal
func areElementSequencesEqual(a []Element, b []Element) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !a[i].Equals(b[i]) {
			return false
		}
	}

	return true
}

/* Rich text */

// Check if two lists of Lines are equal
func areLinesEqual(a []Line, b []Line) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !a[i].Equals(b[i]) {
			return false
		}
	}

	return true
}

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

// Equals checks if both cells are identical
func (line Line) Equals(other Line) bool {
	if len(line) != len(other) {
		return false
	}

	for i := range line {
		if line[i] != other[i] {
			return false
		}
	}

	return true
}

// StringBuffer returns a buffer containing the contents of the line
func (line Line) StringBuffer() *bytes.Buffer {
	buffer := bytes.NewBufferString("")

	for _, cell := range line {
		buffer.WriteString(cell.Content)
	}

	return buffer
}

// String returns a plaintext version of the line
func (line Line) String() string {
	return line.StringBuffer().String()
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
// 0 if the cell is a comment
func (cell Cell) Lenght() int {
	// Comments have no lenght
	if cell.Comment {
		return 0
	}

	return utf8.RuneCountInString(cell.Content)
}

// Empty checks if the cell has no contents
func (cell Cell) Empty() bool {
	return cell.Content == ""
}

// Equals checks if both cells are identical
func (cell Cell) Equals(other Cell) bool {
	return cell == other
}

/* Script sections */

// Script contains the entire script.
type Script struct {
	Language  languages.Language
	TitlePage map[string][]Line
	Elements  []Element
}

// Equals checks if two Scripts are identical
func (script *Script) Equals(other *Script) bool {
	if script.Language != other.Language ||
		len(script.TitlePage) != len(other.TitlePage) ||
		len(script.Elements) != len(other.Elements) {
		return false
	}

	for k := range script.TitlePage {
		if !areLinesEqual(script.TitlePage[k], other.TitlePage[k]) {
			return false
		}
	}

	return areElementSequencesEqual(script.Elements, other.Elements)
}

// Element represents a part of the script.
type Element interface {
	Equals(Element) bool
}

// Scene header type.
type Scene struct {
	Slugline    []Line
	SceneNumber string
}

// Equals checks if an Element is the same Scene
func (scene Scene) Equals(other Element) bool {
	switch other := other.(type) {
	case Scene:
		return areLinesEqual(scene.Slugline, other.Slugline) &&
			scene.SceneNumber == other.SceneNumber
	default:
		return false
	}
}

// BeginAct type
type BeginAct []Line

// Equals checks if an Element is the same BeginAct
func (beginact BeginAct) Equals(other Element) bool {
	switch other := other.(type) {
	case BeginAct:
		return areLinesEqual(beginact, other)
	default:
		return false
	}
}

// EndAct type
type EndAct []Line

// Equals checks if an Element is the same EndAct
func (endact EndAct) Equals(other Element) bool {
	switch other := other.(type) {
	case EndAct:
		return areLinesEqual(endact, other)
	default:
		return false
	}
}

// Action type.
type Action []Line

// Equals checks if an Element is the same Action
func (action Action) Equals(other Element) bool {
	switch other := other.(type) {
	case Action:
		return areLinesEqual(action, other)
	default:
		return false
	}
}

// Dialogue type.
type Dialogue struct {
	Character []Line
	Lines     []Element
}

// Equals checks if an Element is the same Dialogue
func (dialogue Dialogue) Equals(other Element) bool {
	switch other := other.(type) {
	case Dialogue:
		return areLinesEqual(dialogue.Character, other.Character) &&
			areElementSequencesEqual(dialogue.Lines, other.Lines)
	default:
		return false
	}
}

// DualDialogue type.
type DualDialogue struct {
	LCharacter []Line
	LLines     []Element

	RCharacter []Line
	RLines     []Element
}

// Equals checks if an Element is the same DualDialogue
func (dualdialogue DualDialogue) Equals(other Element) bool {
	switch other := other.(type) {
	case DualDialogue:
		return areLinesEqual(dualdialogue.LCharacter, other.LCharacter) &&
			areElementSequencesEqual(dualdialogue.LLines, other.LLines) &&
			areLinesEqual(dualdialogue.RCharacter, other.RCharacter) &&
			areElementSequencesEqual(dualdialogue.RLines, other.RLines)
	default:
		return false
	}
}

// The following three can be used with Dialogue, Lyrics can also be standalone.

// Parenthetical type
type Parenthetical []Line

// Equals checks if an Element is the same Parenthetical
func (parenthetical Parenthetical) Equals(other Element) bool {
	switch other := other.(type) {
	case Parenthetical:
		return areLinesEqual(parenthetical, other)
	default:
		return false
	}
}

// Speech type
type Speech []Line

// Equals checks if an Element is the same Speech
func (speech Speech) Equals(other Element) bool {
	switch other := other.(type) {
	case Speech:
		return areLinesEqual(speech, other)
	default:
		return false
	}
}

// Lyrics type.
type Lyrics []Line

// Equals checks if an Element is the same Lyrics
func (lyrics Lyrics) Equals(other Element) bool {
	switch other := other.(type) {
	case Lyrics:
		return areLinesEqual(lyrics, other)
	default:
		return false
	}
}

// Transition type.
type Transition []Line

// Equals checks if an Element is the same Transition
func (transition Transition) Equals(other Element) bool {
	switch other := other.(type) {
	case Transition:
		return areLinesEqual(transition, other)
	default:
		return false
	}
}

// CenteredText type.
type CenteredText []Line

// Equals checks if an Element is the same CenteredText
func (centeredtext CenteredText) Equals(other Element) bool {
	switch other := other.(type) {
	case CenteredText:
		return areLinesEqual(centeredtext, other)
	default:
		return false
	}
}

// PageBreak type.
type PageBreak struct{}

// Equals checks if an Element is also a PageBreak
func (pagebreak PageBreak) Equals(other Element) bool {
	switch other.(type) {
	case PageBreak:
		return true
	default:
		return false
	}
}

// Section type.
type Section struct {
	Level uint8
	Line  []Line
}

// Equals checks if an Element is the same Section
func (section Section) Equals(other Element) bool {
	switch other := other.(type) {
	case Section:
		return section.Level == other.Level &&
			areLinesEqual(section.Line, other.Line)
	default:
		return false
	}
}

// Synopsis type.
type Synopsis []Line

// Equals checks if an Element is the same Synopsis
func (synopsis Synopsis) Equals(other Element) bool {
	switch other := other.(type) {
	case Synopsis:
		return areLinesEqual(synopsis, other)
	default:
		return false
	}
}

// Note represents a note which isn't part of another element.
type Note []Line

// Equals checks if an Element is the same Note
func (note Note) Equals(other Element) bool {
	switch other := other.(type) {
	case Note:
		return areLinesEqual(note, other)
	default:
		return false
	}
}
