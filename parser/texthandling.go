package parser

import (
	"strings"

	"github.com/Wraparound/wrap/ast"
)

type ipKind byte
type ipType byte

const (
	bold ipKind = iota
	italic
	underline
	note
	backslash
	linebreak
	emptyline
)

const (
	vague ipType = iota
	start
	end
)

type insertionPoint struct {
	Kind      ipKind
	Type      ipType
	Point     int
	Activated bool
	Escaped   bool
}

/*
 * Notes on emphasis in Fountain/Markdown:
 - A * or _ is only seen as start of emphasis if there is a end.
 - Escaping always works.
*/

func textHandler(lines []string) []ast.Line {
	insertPs := []insertionPoint{}
	nowComment := false

	for _, line := range lines {
		// Details of amount of spaces are handled by Parser()
		if line == "" {
			nowComment = false

			insertPs = append(insertPs,
				insertionPoint{
					Kind:      emptyline,
					Type:      start,
					Point:     0,
					Activated: true,
				},
			)

			// Line is done
			continue
		}

		// Replace tab by four spaces
		line = strings.Replace(line, "\t", strings.Repeat(" ", 4), -1)

		// Search for insert points.
		for i := 0; i < len(line); i++ {
			// Is it escaped?
			escaped := false

			if i > 0 && line[i-1] == '\\' {
				escaped = true
			}

			switch line[i] {
			case '*':
				// If not escaped: bold? Need to be able too look forward.
				if !escaped && i+1 < len(line) && line[i+1] == '*' {
					// Vague? (no surrounding spaces)
					if (i != 0 && line[i-1] != ' ') && (i+2 < len(line) && line[i+2] != ' ') {
						insertPs = append(insertPs,
							insertionPoint{
								Kind:      bold,
								Type:      vague,
								Point:     i,
								Activated: false,
								Escaped:   escaped,
							},
						)

						// Start? -> No space after.
					} else if i+2 < len(line) && line[i+2] != ' ' {
						insertPs = append(insertPs,
							insertionPoint{
								Kind:      bold,
								Type:      start,
								Point:     i,
								Activated: false,
								Escaped:   escaped,
							},
						)

						// End? -> No space before.
					} else if i != 0 && line[i-1] != ' ' {
						insertPs = append(insertPs,
							insertionPoint{
								Kind:      bold,
								Type:      end,
								Point:     i,
								Activated: false,
								Escaped:   escaped,
							},
						)
					}

					// Else: surrounded by spaces: literal
					// Now skip the next *
					i++

				} else {
					// No it's italics or escaped (thus a single *)
					// Vague? (no surrounding spaces)
					if (i != 0 && line[i-1] != ' ') && (i+1 < len(line) && line[i+1] != ' ') {
						insertPs = append(insertPs,
							insertionPoint{
								Kind:      italic,
								Type:      vague,
								Point:     i,
								Activated: false,
								Escaped:   escaped,
							},
						)

						// Start? -> No space after.
					} else if i+1 < len(line) && line[i+1] != ' ' {
						insertPs = append(insertPs,
							insertionPoint{
								Kind:      italic,
								Type:      start,
								Point:     i,
								Activated: false,
								Escaped:   escaped,
							},
						)

						// End? -> No space before.
					} else if i != 0 && line[i-1] != ' ' {
						insertPs = append(insertPs,
							insertionPoint{
								Kind:      italic,
								Type:      end,
								Point:     i,
								Activated: false,
								Escaped:   escaped,
							},
						)
					}
				}

				// Else: surrounded by spaces: literal

			case '_':
				// Underline
				// Vague? (no surrounding spaces)
				if (i != 0 && line[i-1] != ' ') && (i+1 < len(line) && line[i+1] != ' ') {
					insertPs = append(insertPs,
						insertionPoint{
							Kind:      underline,
							Type:      vague,
							Point:     i,
							Activated: false,
							Escaped:   escaped,
						},
					)

					// Start? -> No space after.
				} else if i+1 < len(line) && line[i+1] != ' ' {
					insertPs = append(insertPs,
						insertionPoint{
							Kind:      underline,
							Type:      start,
							Point:     i,
							Activated: false,
							Escaped:   escaped,
						},
					)

					// End? -> No space before.
				} else if i != 0 && line[i-1] != ' ' {
					insertPs = append(insertPs,
						insertionPoint{
							Kind:      underline,
							Type:      end,
							Point:     i,
							Activated: false,
							Escaped:   escaped,
						},
					)
				}

				// Else: surrounded by spaces: literal

			case '[':
				// Not yet in comment and start comment? Need to be able too look forward.
				if !nowComment && i+1 < len(line) && line[i+1] == '[' {
					nowComment = true
					insertPs = append(insertPs,
						insertionPoint{
							Kind:      note,
							Type:      start,
							Point:     i,
							Activated: true,
							Escaped:   false, // Notes can't be escaped.
						},
					)
				}

			case ']':
				// Not yet in comment and end comment? Need to be able too look forward.
				if nowComment && i+1 < len(line) && line[i+1] == ']' {
					nowComment = false
					insertPs = append(insertPs,
						insertionPoint{
							Kind:      note,
							Type:      end,
							Point:     i,
							Activated: true,
							Escaped:   false, // Notes can't be escaped.
						},
					)
				}
			case '\\':
				// Handle escaped backslashes.
				if escaped {
					insertPs = append(insertPs,
						insertionPoint{
							Kind:      backslash,
							Type:      vague,
							Point:     i,
							Activated: false,
							Escaped:   true,
						},
					)
				}
				// Else: skip
			}
		}

		insertPs = append(insertPs,
			insertionPoint{
				Kind:      linebreak,
				Type:      start,
				Point:     len(line),
				Activated: true,
			},
		)
	}

	// Activate insert points.
	// No need to check comments, they are already activated when needed above.
	// This also handles bold as italics
	nowBold := false
	startPointBold := 0
	nowItalic := false
	startPointItalic := 0
	nowPotentiallyItalic := false // Keep track of double ** that start italics
	startPointPotentiallyItalic := 0
	nowUnderline := false
	startPointUnderline := 0

	for location := 0; location < len(insertPs); location++ {
		// An insertpoint can only be activated when it's not escaped.
		if !insertPs[location].Escaped {
			switch insertPs[location].Type {
			case start:
				switch insertPs[location].Kind {
				// case linebreak don't do anything

				case emptyline:
					nowBold = false
					nowItalic = false
					nowPotentiallyItalic = false
					nowUnderline = false

				case bold:
					nowBold = true
					startPointBold = location
					nowPotentiallyItalic = true // Might be italics...
					startPointPotentiallyItalic = location

				case italic:
					nowItalic = true
					startPointItalic = location

				case underline:
					nowUnderline = true
					startPointUnderline = location
				}

			case end:
				switch insertPs[location].Kind {
				case bold:
					if nowBold {
						nowBold = false
						nowPotentiallyItalic = false // Nope it really is bold.
						insertPs[startPointBold].Activated = true
						insertPs[location].Activated = true

					} else if nowItalic {
						// This could also be the end of italics...
						// Not potential italics, as those were bold.
						nowItalic = false
						insertPs[startPointItalic].Activated = true
						insertPs[location].Kind = italic
						insertPs[location].Activated = true
					}

				case italic:
					if nowItalic {
						nowItalic = false
						nowPotentiallyItalic = false // This isn't possible anymore...
						insertPs[startPointItalic].Activated = true
						insertPs[location].Activated = true

					} else if nowPotentiallyItalic {
						nowPotentiallyItalic = false
						nowBold = false // Nope it apparently was italics
						insertPs[startPointPotentiallyItalic].Activated = true
						insertPs[startPointPotentiallyItalic].Kind = italic
						insertPs[location].Activated = true
					}

				case underline:
					if nowUnderline {
						nowUnderline = false
						insertPs[startPointUnderline].Activated = true
						insertPs[location].Activated = true
					}
				}

			case vague:
				switch insertPs[location].Kind {
				case bold:
					if nowBold {
						nowBold = false
						nowPotentiallyItalic = false // Not possible anymore
						insertPs[startPointBold].Activated = true
						insertPs[location].Type = end
						insertPs[location].Activated = true

					} else if nowItalic {
						nowItalic = false
						// We now have to split this in two, so first make this point
						// An active italics end.
						insertPs[startPointItalic].Activated = true
						insertPs[location].Type = end
						insertPs[location].Kind = italic
						insertPs[location].Activated = true

						// Now add the next character as a potential italics start.
						insertPsTail := insertPs[location+1:]
						insertPs = append(insertPs[:location+1],
							insertionPoint{
								Kind:      italic,
								Type:      start,
								Point:     insertPs[location].Point + 1,
								Activated: false,
								Escaped:   false,
							},
						)
						insertPs = append(insertPs, insertPsTail...)

					} else {
						nowBold = true
						startPointBold = location
						nowPotentiallyItalic = true // Might be italics...
						startPointPotentiallyItalic = location
						insertPs[location].Type = start
					}

				case italic:
					if nowItalic {
						nowItalic = false
						nowPotentiallyItalic = false // This isn't possible anymore...
						insertPs[startPointItalic].Activated = true
						insertPs[location].Type = end
						insertPs[location].Activated = true

					} else if nowPotentiallyItalic && // And they are not next to each other
						insertPs[location].Point-insertPs[startPointPotentiallyItalic].Point != 2 {

						nowPotentiallyItalic = false
						nowBold = false // Nope it apparently was italics
						insertPs[startPointPotentiallyItalic].Activated = true
						insertPs[startPointPotentiallyItalic].Kind = italic
						insertPs[location].Type = end
						insertPs[location].Activated = true

					} else {
						nowItalic = true
						insertPs[location].Type = start
						startPointItalic = location
					}

				case underline:
					if nowUnderline {
						nowUnderline = false
						insertPs[startPointUnderline].Activated = true
						insertPs[location].Type = end
						insertPs[location].Activated = true

					} else {
						nowUnderline = true
						insertPs[location].Type = start
						startPointUnderline = location
					}
				}
			}
		}
	}

	// Start writing
	endResult := []ast.Line{}
	currentLine := ast.Line{}
	currentCell := ast.Cell{}

	// Shifting isn't an issue, as we don't change the read string.
	lastWritePoint := 0

	currentlyBold := false
	currentlyItalic := false
	currentlyUnderline := false
	currentlyComment := false

	line := 0
	for _, insPoint := range insertPs {
		if insPoint.Escaped {
			// Add everything appart from the (first) backslash.
			currentCell.Content += lines[line][lastWritePoint : insPoint.Point-1]
			// Update lastWritePoint so that the next write opperation will include the
			// character which has been escaped.
			lastWritePoint = insPoint.Point

		} else if insPoint.Activated {
			currentCell.Content += lines[line][lastWritePoint:insPoint.Point]

			switch insPoint.Kind {
			case bold:
				lastWritePoint = insPoint.Point + 2

				switch insPoint.Type {
				case start:
					currentlyBold = true

				case end:
					currentlyBold = false
				}

			case italic:
				lastWritePoint = insPoint.Point + 1

				switch insPoint.Type {
				case start:
					currentlyItalic = true

				case end:
					currentlyItalic = false
				}

			case underline:
				lastWritePoint = insPoint.Point + 1

				switch insPoint.Type {
				case start:
					currentlyUnderline = true

				case end:
					currentlyUnderline = false
				}

			case note:
				lastWritePoint = insPoint.Point + 2

				switch insPoint.Type {
				case start:
					currentlyComment = true

				case end:
					currentlyComment = false
				}

			case emptyline:
				currentlyBold = false
				currentlyItalic = false
				currentlyUnderline = false
				currentlyComment = false
				fallthrough

			case linebreak:
				// Add cell to line and add line to result.
				if !currentCell.Empty() {
					currentLine = append(currentLine, currentCell)
				}

				endResult = append(endResult, currentLine)

				// Reset for next line
				currentLine = ast.Line{}

				currentCell = ast.Cell{
					Boldface:  currentlyBold,
					Italics:   currentlyItalic,
					Underline: currentlyUnderline,
					Comment:   currentlyComment,
				}

				line++
				lastWritePoint = 0
			}

			if !currentCell.Empty() {
				currentLine = append(currentLine, currentCell)
			}

			currentCell = ast.Cell{
				Boldface:  currentlyBold,
				Italics:   currentlyItalic,
				Underline: currentlyUnderline,
				Comment:   currentlyComment,
			}
		}
	}

	return endResult
}
