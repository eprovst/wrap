package languages

import "strings"

// This file defines general structures to give a flexibel translation system

// Language represents a supported language
type Language byte

// Translation represents a specific transaltion of Wrap
type Translation struct {
	Language Language

	SceneTags          []string
	StageplaySceneTags []string
	TransitionTags     []string
	BeginActTags       []string
	EndActTags         []string

	More  string
	Contd string
}

// All available languages
const (
	English Language = iota
	French
	German
	Dutch
)

// Default language is English
var Default = English

// GetLanguage allows to convert aliases to a Language defaults to Default
func GetLanguage(alias string) Language {
	switch strings.ToLower(alias) {
	case
		"english",
		"american",
		"australian":

		return English

	case
		"french",
		"francais", // ç is difficult to type on some/most keyboard layouts...
		"français":

		return French

	case
		"german",
		"deutsch":

		return German

	case
		"dutch",
		"flemish",
		"nederlands",
		"vlaams":

		return Dutch

	default:
		return Default
	}
}

// String allows us to convert a Language to it's standard name, panics when language unknown
func (lang Language) String() string {
	switch lang {
	case English:
		return "English"

	case French:
		return "French"

	case German:
		return "German"

	case Dutch:
		return "Dutch"

	default:
		panic("unknown language")
	}
}

// Translation gives the translation for the language, panics when no translation available
func (lang Language) Translation() Translation {
	switch lang {
	case English:
		return EnglishTranslation

	case French:
		return FrenchTranslation

	case German:
		return GermanTranslation

	case Dutch:
		return DutchTranslation

	default:
		panic("unknown language")
	}
}
