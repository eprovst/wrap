package parser

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/Wraparound/wrap/ast"
	"github.com/Wraparound/wrap/languages"
)

// UseWrapExtensions is a flag to show wheter
// we are using the Wrap extentions or not.
// The user of this library can change the flag.
var UseWrapExtensions = false

// language contains the language used.
var language = languages.Default
var translation languages.Translation

type categorisedLine struct {
	// See identifier.go
	Category lineCat
	Forced   bool
	Line     string
}

// Parser parses an io.Reader
func Parser(input io.Reader) (*ast.Script, error) {
	// Wrap the input into a bufio.Reader for easier handling.
	scanner := bufio.NewScanner(input)

	// A variable to contain the current line and potential error.
	var line string
	var err error

	// Create a map to contain the titlepage.
	titlp := make(map[string][]ast.Line)

	// Read the first line.
	if scanner.Scan() {
		line = scanner.Text()

	} else if err = scanner.Err(); err != nil {
		// The input is empty so give them an empty script.
		return nil, err

	} else {
		return &ast.Script{}, nil
	}

	// A variable to remember if the last line was blank.
	var lastLineBlank bool

	/* As we work on the text scanned during the previous run a loop
	we need to keep if it was succesful back then, otherwise we won't read
	the last line. */
	var lastScanSuccessful = true

	/* If begin metadata, start PageTitle handling.
	Here unnecessary spacing is handled through regex.
	NOTE: regex is no time concern here as a title page is quite small:
	no noticable difference.*/
	if regexp.MustCompile("^(\\S+|\\S.+\\S)\\:(\\s.+)?$").MatchString(line) {
		lastLineBlank = false

		inlineValue := regexp.MustCompile("^(\\S+|\\S.+\\S)(?:\\:\\s+)(\\S+|\\S.+\\S)(?:\\s*)$")
		multiLineKey := regexp.MustCompile("^(\\S+|\\S.+\\S)(?:\\:\\s*)$")
		multiLineValue := regexp.MustCompile("^(?:   |\t)(?:\\s*)(\\S+|\\S.+\\S)(?:\\s*)$")
		//                                         ^ tripple space.

		onlyWhiteSpace := regexp.MustCompile("^\\s*$")

		/* PARSE THE TITLE PAGE */
	titlePageAgreggator:
		for lastScanSuccessful {
			if inlineValue.MatchString(line) {
				vls := inlineValue.FindStringSubmatch(line)
				titlp[strings.ToLower(vls[1])] = textHandler([]string{vls[2]})
				lastLineBlank = false

			} else if multiLineKey.MatchString(line) {
				key := strings.ToLower(multiLineKey.FindStringSubmatch(line)[1])
				var value []string

				lastScanSuccessful = scanner.Scan()
				line = scanner.Text()

				for lastScanSuccessful && (onlyWhiteSpace.MatchString(line) || multiLineValue.MatchString(line)) {
					if onlyWhiteSpace.MatchString(line) {
						if lastLineBlank {
							// Two blank lines -> end TitlePage
							titlp[key] = textHandler(value)

							// Prepare a line
							lastScanSuccessful = scanner.Scan()
							line = scanner.Text()
							break titlePageAgreggator
						}

						value = append(value, "")
						lastLineBlank = true

					} else {
						value = append(value, multiLineValue.FindStringSubmatch(line)[1])
						lastLineBlank = false
					}

					lastScanSuccessful = scanner.Scan()
					line = scanner.Text()
				}

				titlp[key] = textHandler(value)

				// Skip default scan as we already scanned.
				continue

			} else if onlyWhiteSpace.MatchString(line) {
				if lastLineBlank {
					// Prepare a line for the next part of the parser.
					lastScanSuccessful = scanner.Scan()
					line = scanner.Text()
					break

				} else {
					// Maybe the start of the end... of the TitlePage.
					lastLineBlank = true
				}

			} else {
				/* The current line is invalid, the Fountain spec
				doesn't define what to do now, but we'll just act
				as if it's the end of the Title Page and give the
				line to the next part of the program. */
				break
			}

			lastScanSuccessful = scanner.Scan()
			line = scanner.Text()
		}
	}

	// All metadata is collected, now set the used language.
	richLanguage := titlp["language"]

	if UseWrapExtensions && len(richLanguage) != 0 {
		language = languages.GetLanguage(richLanguage[0].String())

	} else {
		/* If we're not using the Wrap extensions the language
		is made the default, the global could be changed already. */
		language = languages.Default
	}

	// Get the translation of the current language
	translation = language.Translation()

	/* PARSE THE SCRIPT */

	// First we read, categorise and clean the forced lines.
	// The first line was already scanned by the previous step.
	var lines []categorisedLine

	for lastScanSuccessful {

		lines = append(lines, categoriser(line))
		lastScanSuccessful = scanner.Scan()
		line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// In this step we check rules concerning lines around the current line and do the final cleanup.
	for i := 1; i < len(lines); i++ {
		// A slightly more complex for loop as we remove lines.
		if !lines[i].Forced {
			switch lines[i].Category {
			// Has to be surrounded by empty lines
			case beginAct, endAct, transitionTag, sceneTag:
				if i > 0 && isIgnoredLineType(lines[i-1].Category) &&
					i+1 < len(lines) && isIgnoredLineType(lines[i+1].Category) {
					// Cleanup of current line
					lines[i].Line = normaliseLine(lines[i].Line)

				} else {
					lines[i].Category = action
				}

			// Has to be followed by an empty line.
			case character, characterTwo:
				if i > 0 && isIgnoredLineType(lines[i-1].Category) &&
					!(i+1 < len(lines) && isIgnoredLineType(lines[i+1].Category)) {
					// Cleanup of current line
					lines[i].Line = normaliseLine(lines[i].Line)

					if lines[i].Category == characterTwo {
						// Extra cleanup if a character two
						lines[i].Line = strings.TrimSuffix(lines[i].Line, "^")
						lines[i].Line = strings.TrimSpace(lines[i].Line)
					}

				} else {
					lines[i].Category = action
				}

			case synopse, section, noteOnOwnLine:
				// These gobble their surrounding lines.
				gobbledLines := 0

				if i > 0 && lines[i-1].Category == emptyLine {
					lines = append(lines[:i-1], lines[i:]...)
					i-- // Correct for removal of last line
					gobbledLines++
				}

				if i+1 < len(lines) && lines[i+1].Category == emptyLine {
					lines = append(lines[:i+1], lines[i+2:]...)
					gobbledLines++
				}

				if lines[i].Category == noteOnOwnLine && gobbledLines != 2 {
					// If it isn't a true note on one line it didn't gobble it's surrounding lines,
					// thus it's an action.
					lines[i].Category = action
				}
			}
		}
	}

	// Now we group the lines, interpret their formating and put them into elements
	// A list to contain the elements of our script.
	var elems []ast.Element
	sceneNumber := 0

	for i := 0; i < len(lines); i++ {
		// To make dualdialogue work
		mergeDialogue := false

		switch lines[i].Category {
		// Empty lines are ignored for as long they aren't part of an action element.
		case sceneTag:
			slugline, scenenumber := parseSceneHeading(lines[i].Line)

			// Auto add scenenumber.
			if scenenumber == "" {
				sceneNumber++
				scenenumber = strconv.Itoa(sceneNumber)
			}

			elems = append(elems, ast.Scene{Slugline: textHandler([]string{slugline}), SceneNumber: scenenumber})

		case boneyard:
			// No write as we just throw it away. And we start at the current line as boneyard can be inline.
			if !strings.Contains(lines[i].Line, "/*") {
				// Only contains an end marker, let's call that an action.
				lines[i].Category = action

				// Now redo this line:
				i--
				continue
			}

			lineBeginning := strings.SplitN(lines[i].Line, "/*", 2)[0]
			for _ = i; i < len(lines); i++ {
				if lines[i].Category == boneyard && strings.Contains(lines[i].Line, "*/") {
					break
					// No index correction here, as the line is still part of the boneyard.
				}
			}

			combinedLine := lineBeginning + strings.SplitN(lines[i].Line, "*/", 2)[1]
			lines[i] = categoriser(combinedLine)

			// Now revisit this line
			i--

		case action:
			contents := []string{lines[i].Line}

		actionAggregation:
			for i++; i < len(lines); i++ {
				switch lines[i].Category {
				case action:
					contents = append(contents, lines[i].Line)
				case emptyLine:
					// No checking of "double empty" as 'end of scope' is handled by textHandler().
					contents = append(contents, "")
				default:
					i-- // Revisit this line.
					break actionAggregation
				}
			}
			elems = append(elems, ast.Action(textHandler(contents)))

		case characterTwo:
			// Slightly awkward but it works...
			if len(elems) == 0 {
				// Yeah there's nothing before.
				lines[i].Category = character

			} else {
				switch elems[len(elems)-1].(type) {
				case ast.Dialogue:
					mergeDialogue = true

				default:
					// There isn't any dialogue to append to, ignore it.
					lines[i].Category = character
				}
			}

			fallthrough

		case character:
			charact := lines[i].Line
			var dialog []ast.Element

		dialogueAggregation:
			for i++; i < len(lines); i++ {
				switch lines[i].Category {
				case other:
					if isParenthetical(lines[i].Line) {
						lines := textHandler([]string{strings.TrimSpace(lines[i].Line)})
						dialog = append(dialog, ast.Parenthetical(lines))

					} else {
						contents := []string{strings.TrimSpace(lines[i].Line)}
						for i++; i < len(lines); i++ {
							if lines[i].Category == other && !isParenthetical(lines[i].Line) {
								contents = append(contents, strings.TrimSpace(lines[i].Line))

							} else {
								i-- // Revisit this line.
								break
							}
						}

						dialog = append(dialog, ast.Speech(textHandler(contents)))
					}

				case lyrics:
					contents := []string{lines[i].Line}
					for i++; i < len(lines); i++ {
						if lines[i].Category == lyrics {
							contents = append(contents, lines[i].Line)

						} else {
							i-- // Revisit this line.
							break
						}
					}

					dialog = append(dialog, ast.Lyrics(textHandler(contents)))

				default:
					break dialogueAggregation
				}
			}

			if mergeDialogue {
				lastDia := elems[len(elems)-1].(ast.Dialogue)
				elems[len(elems)-1] = ast.DualDialogue{
					LCharacter: lastDia.Character,
					LLines:     lastDia.Lines,
					RCharacter: textHandler([]string{charact}),
					RLines:     dialog,
				}

				// Clean up
				mergeDialogue = false

			} else {
				elems = append(elems, ast.Dialogue{
					Character: textHandler([]string{charact}),
					Lines:     dialog,
				})
			}

		case lyrics:
			contents := []string{lines[i].Line}
			for i++; i < len(lines); i++ {
				if lines[i].Category == lyrics {
					contents = append(contents, lines[i].Line)

				} else {
					i-- // Revisit this line.
					break
				}
			}

			elems = append(elems, ast.Lyrics(textHandler(contents)))

		case transitionTag:
			elems = append(elems, ast.Transition(textHandler([]string{lines[i].Line})))

		case centeredText:
			contents := []string{lines[i].Line}

		centeredTextAggregation:
			for i++; i < len(lines); i++ {
				switch lines[i].Category {
				case centeredText:
					contents = append(contents, lines[i].Line)

				case emptyLine:
					contents = append(contents, "")

					if lines[i-1].Category == emptyLine {
						// End of scope.
						break centeredTextAggregation
					}

				default:
					i-- // Revisit this line.
					break centeredTextAggregation
				}
			}

			elems = append(elems, ast.CenteredText(textHandler(contents)))

		case pageBreak:
			elems = append(elems, ast.PageBreak{})

		case section:
			descr, level := parseSection(lines[i].Line)
			elems = append(elems, ast.Section{
				Line:  textHandler([]string{descr}),
				Level: level,
			})

		case synopse:
			elems = append(elems, ast.Synopse(textHandler([]string{lines[i].Line})))

		case beginAct:
			elems = append(elems, ast.BeginAct(textHandler([]string{lines[i].Line})))

		case endAct:
			elems = append(elems, ast.EndAct(textHandler([]string{lines[i].Line})))

		case noteOnOwnLine:
			// We leave the <ins></ins>, as CSS/theming could/should depend on it.
			elems = append(elems, ast.Note(textHandler([]string{lines[i].Line})))
		}
	}

	// Finaly bring everything together in an ast.Script
	tree := ast.Script{
		Language:  language,
		TitlePage: titlp,
		Elements:  elems,
	}

	return &tree, nil
}
