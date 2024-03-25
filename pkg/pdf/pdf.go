package pdf

import (
	"math"
	"strings"
	"time"

	"github.com/eprovst/wrap/pkg/ast"
	"github.com/eprovst/wrap/pkg/languages"
	"github.com/signintech/gopdf"
)

const (
	fontSize      int     = 12
	point         float64 = 1
	pica          float64 = 12 * point
	en            float64 = in / 10 // 10 pitch.
	em            float64 = float64(fontSize)
	in            float64 = 6 * pica
	vPageWidth    float64 = 8.5 * in
	vPageHeight   float64 = 11 * in
	vTopMargin    float64 = 1 * in
	vLeftMargin   float64 = 1.5 * in
	vRightMargin  float64 = 1 * in
	vBottomMargin float64 = 1 * in
	maxNumOfLines int     = 55
	maxNumOfChars int     = 60
)

var (
	pageWidth    float64
	pageHeight   float64
	topMargin    float64
	leftMargin   float64
	rightMargin  float64
	bottomMargin float64
)

// Production makes the export module add scene numbers and other
// production specific additions
var Production = true

// PageSize specifies the prefered page size
var PageSize = ""

var currentTheme = screenplay
var currentTranslation = languages.Default.Translation()

var linesOnPage int
var pageNumber int
var thisPDF = &gopdf.GoPdf{}

// buildPDF creates a PDF file structure from a script.
func buildPDF(script *ast.Script) (*gopdf.GoPdf, error) {
	// Handle custom page size
	if PageSize == "" {
		richPageSize := script.TitlePage["pagesize"]
		if len(richPageSize) != 0 {
			PageSize = richPageSize[0].String()
		}
	}

	// Set page size
	selectedPageSize := *gopdf.PageSizeLetter
	if strings.EqualFold(PageSize, "a4") {
		selectedPageSize = *gopdf.PageSizeA4
	}
	thisPDF.Start(gopdf.Config{PageSize: selectedPageSize})

	// Correct margins
	pageWidth = selectedPageSize.W
	pageHeight = selectedPageSize.H

	hOff := (pageWidth - vPageWidth) / 2
	vOff := (pageHeight - vPageHeight) / 2
	leftMargin = vLeftMargin + hOff
	rightMargin = vRightMargin + hOff
	topMargin = vTopMargin + vOff
	bottomMargin = vBottomMargin + vOff

	// Set language:
	currentTranslation = script.Language.Translation()

	// ADD FONTS
	loadFonts()

	// HANDLE METADATA:
	// Handle play theming
	richTheme := script.TitlePage["type"]
	if len(richTheme) != 0 {
		currentTheme = themeMap[strings.ToLower(richTheme[0].String())]
	}

	if currentTheme == nil {
		currentTheme = screenplay
	}

	// Handle custom CONT'D and MORE tags
	richContd := script.TitlePage["contdtag"]
	if len(richContd) != 0 {
		currentTranslation.Contd = richContd[0].String()
	}

	richMore := script.TitlePage["moretag"]
	if len(richMore) != 0 {
		currentTranslation.More = richMore[0].String()
	}

	// Clean up title
	title := ast.LinesToString(script.TitlePage["title"])
	title = strings.Replace(title, "\n", " ", -1)

	// Get the author(s)
	authors := script.TitlePage["authors"]
	if len(authors) == 0 || authors[0].Lenght() == 0 {
		authors = script.TitlePage["author"]
	}

	// Add PDF info
	thisPDF.SetInfo(gopdf.PdfInfo{
		Title:        title,
		Author:       ast.LinesToString(authors),
		Creator:      "Wrap",
		Producer:     "Wrap PDF",
		CreationDate: time.Now(),
	})

	// Minimize size on disk
	thisPDF.SetCompressLevel(2)

	// START PDF BUILDING:

	// Start building the title page
	addTitlePage(script)

	// Sectionize lines
	var sections []aSection

	for _, element := range script.Elements {
		sections = append(sections, sectionize(element))
	}

	// Start a new page
	newPage()

	for _, section := range sections {
		// TODO: Improve page breaking.
		// MAYBE TODO: Dual dialogue pagebreaking with (cont'd)

		switch section.Type {
		case ast.TPageBreak:
			if linesOnPage != 0 {
				newPage()
			}

		case ast.TBeginAct:
			if linesOnPage != 0 {
				newPage()
			}
			addLines(section.Lines)

		case ast.TScene:
			if linesOnPage+section.height()+2 > maxNumOfLines {
				newPage()
			}
			// TODO: More pagebreaking stuff...

			if Production {
				// Add scene numbers
				oldY := thisPDF.GetY()
				firstLineLeading := section.Lines[0].leading()
				thisPDF.Br(float64(firstLineLeading) * em)
				// Left
				thisPDF.SetX(leftMargin - 7.5*en)
				addCell(section.Metadata["scenenumber"].(ast.Cell))
				// Right
				thisPDF.SetX(pageWidth - rightMargin + 2.5*en)
				addCell(section.Metadata["scenenumber"].(ast.Cell))

				thisPDF.SetY(oldY) // Prepare for the slugline.
			}

			// Add the sceneheading itself.
			addLines(section.Lines)

		case ast.TDialogue:
			var lastCharacterLine styledLine
			for i, line := range section.Lines {
				addedLeading := line.leading()

				switch line.Type {
				case character:
					// Keep track of this line.
					lastCharacterLine = line
					// This won't always work, eg. dialogue of one line etc.
					if linesOnPage+5 > maxNumOfLines {
						newPage()
					}
					// TODO: More pagebreaking stuff...
					// Add the line and keep track of it.
					addLine(line)

				case parenthetical:
					fallthrough // TODO Should be "keep together" actually,
					// keep in mind it could be more than one page long though...

				case dialogue, lyrics:
					// Line might not fit on page (if it fits, the next line won't and will handle page breaking)
					if linesOnPage+addedLeading+1 >= maxNumOfLines &&
						// If it just fits we skip this special breaking, unless it's followed by dialogue stuff.
						!(linesOnPage+addedLeading+1 == maxNumOfLines && i+1 >= len(section.Lines)) {
						// First add the more tag.
						addLine(styledLine{
							Type: more,
							Content: []ast.Cell{ast.Cell{
								Content: currentTranslation.More,
							}},
						})
						// Now go to the next page.
						newPage()
						// Prepare a charactertag.
						tmpLine := lastCharacterLine
						// Only add (cont'd) tag when not yet present:
						if len(tmpLine.Content) > 0 && !strings.HasSuffix(strings.ToLower(strings.TrimSpace(
							tmpLine.Content[len(tmpLine.Content)-1].Content)), strings.ToLower(currentTranslation.Contd)) {
							tmpLine.Content = append(tmpLine.Content, ast.Cell{
								Content: " " + currentTranslation.Contd,
							})
						}
						// Add the charactertag and keep track of it.
						addLine(tmpLine)
					}

					// Add the line and keep track of it.
					addLine(line)
				}
			}

		case ast.TDualDialogue:
			leftY := thisPDF.GetY()
			leftLines := linesOnPage
			rightY := thisPDF.GetY()
			rightLines := linesOnPage

			for _, line := range section.Lines {
				addedLeading := line.leading()

				switch line.Type {
				case dualCharacterOne, dualDialogueOne, dualParentheticalOne, dualLyricsOne:
					thisPDF.SetY(leftY)

					// Line doesn't fit on page:
					if leftLines+addedLeading+1 > maxNumOfLines {
						newPage()

						leftLines = linesOnPage
						rightY = thisPDF.GetY()
						rightLines = linesOnPage
					}

					// Add the line. (styleLine() as we do not want to track it globally yet)
					styleLine(line)
					leftLines += addedLeading + 1

					leftY = thisPDF.GetY()

				case dualCharacterTwo, dualDialogueTwo, dualParentheticalTwo, dualLyricsTwo:
					thisPDF.SetY(rightY)

					// Line doesn't fit on page:
					if rightLines+addedLeading+1 > maxNumOfLines {
						newPage()

						leftY = thisPDF.GetY()
						leftLines = linesOnPage
						rightLines = linesOnPage
					}

					// Add the line. (styleLine() as we do not want to track it globally yet)
					styleLine(line)
					rightLines += addedLeading + 1

					rightY = thisPDF.GetY()
				}
			}

			// Update lines on page and Y
			linesOnPage = max(leftLines, rightLines)
			thisPDF.SetY(math.Max(leftY, rightY))

		default:
			addLines(section.Lines)
		}
	}

	return thisPDF, nil
}
