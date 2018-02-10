package pdf

import (
	"html"
	"strings"
	"unicode/utf8"
)

type aLine struct {
	Content        []aCell
	Type           lineType
	FirstOfSection bool
}

func (line aLine) len() int {
	lineLength := 0
	for _, cell := range line.Content {
		lineLength += cell.len()
	}
	return lineLength
}

func (line aLine) isEmpty() bool {
	for _, cell := range line.Content {
		if len(strings.TrimSpace(cell.Content)) > 0 {
			return false
		}
	}
	return true
}

type aCell struct {
	Content   string
	Boldface  bool
	Italics   bool
	Underline bool
}

func (cell aCell) len() int {
	return utf8.RuneCountInString(cell.Content)
}

/* Prepares a section for insertion into the PDF.
Linelength is the maximum linelenght in characters (we work with a monospace font). */
func cellify(text string, style lineType) []aLine {
	// Remove comments.
	text = removeBetweenTags(text, "ins")

	// Now sart building the line list.
	lines := []aLine{}

	for _, line := range strings.Split(text, "<br>\n") {
		// Split into tokens, even index will be content, uneven tags.
		tmpSplit := strings.Split(line, "<")
		var tokens []string

		// Now split at '>'.
		for _, part := range tmpSplit {
			tokens = append(tokens, strings.Split(part, ">")...)
		}

		var (
			currentlyBold      bool
			currentlyItalic    bool
			currentlyUnderline bool
		)

		line := aLine{
			Type: style,
		}

		for i, token := range tokens {
			// Even, thus content: add
			if i%2 == 0 {
				// Unless there is no content, of cource...
				if token != "" {
					line.Content = append(line.Content, aCell{
						Content:   html.UnescapeString(token), // Normalise cell
						Boldface:  currentlyBold,
						Italics:   currentlyItalic,
						Underline: currentlyUnderline,
					})
				}
			} else {
				// Now i is uneven thus a tag:
				switch token {
				case "b":
					currentlyBold = true
				case "/b":
					currentlyBold = false
				case "i":
					currentlyItalic = true
				case "/i":
					currentlyItalic = false
				case "u":
					currentlyUnderline = true
				case "/u":
					currentlyUnderline = false
				}
			}
		}

		lines = append(lines, wordwrap(line)...)
	}

	// If the last break was only a break, then...
	if lines[len(lines)-1].isEmpty() {
		// ...remove the last line.
		lines = lines[:len(lines)-1]
	}

	// Lable the first line as first of section
	if len(lines) > 0 {
		lines[0].FirstOfSection = true
	}

	return lines
}
