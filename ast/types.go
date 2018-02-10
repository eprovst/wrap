package ast

import "github.com/Wraparound/wrap/languages"

/* This file defines all the types used by the parser
to generate a syntax tree. */

/* Text formating is done through HTML tags, this includes linebreaks.
Fountain notes are enclosed by <ins></ins>.
The parser adds  both a <br> tag and newlines into the strings. */

// Script contains the entire script.
type Script struct {
	Language  languages.Language
	TitlePage map[string]string
	Elements  []Element
}

// Element represents a part of the script.
type Element interface{}

// Scene header type.
type Scene struct {
	Slugline    string
	SceneNumber string
}

// BeginAct type
type BeginAct string

// EndAct type
type EndAct string

// Action type.
type Action string

// Dialogue type.
type Dialogue struct {
	Character string
	Lines     []Element
}

// DualDialogue type.
type DualDialogue struct {
	LCharacter string
	LLines     []Element

	RCharacter string
	RLines     []Element
}

// The following three can be used with Dialogue, Lyrics can also be standalone.

// Parenthetical type
type Parenthetical string

// Speech type
type Speech string

// Lyrics type.
type Lyrics string

// Transition type.
type Transition string

// CenteredText type.
type CenteredText string

// PageBreak type.
type PageBreak struct{}

// Section type.
type Section struct {
	Level uint8
	Line  string
}

// Synopse type.
type Synopse string

// Note represents a note which isn't part of another element.
type Note string
