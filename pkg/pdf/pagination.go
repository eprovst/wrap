package pdf

import (
	"strconv"

	"github.com/Wraparound/wrap/pkg/ast"
)

func addLines(lines []styledLine) {
	// Add each line individualy
	for _, line := range lines {
		if !(thisPDF.GetY() == topMargin && line.isEmpty()) {
			// ^^ skip empty lines on the beginning of the page.
			addLine(line)
		}
	}
}

func addLine(line styledLine) {
	addedLeading := line.leading()

	// Line doesn't fit on page:
	if linesOnPage+addedLeading+1 > maxNumOfLines {
		newPage()
	}

	// Add the line and keep track of it.
	styleLine(line)
	linesOnPage += 1 + addedLeading
}

func newPage() {
	thisPDF.AddPage()
	pageNumber++

	// Add a pagenumber.
	if pageNumber > 1 {
		pageNumStr := strconv.Itoa(pageNumber) + "."
		thisPDF.SetX(pageWidth - rightMargin - float64(len(pageNumStr))*en)
		thisPDF.SetY(topMargin/2 - 1*em)
		thisPDF.Cell(nil, pageNumStr)
	}

	thisPDF.SetY(topMargin)
	linesOnPage = 0
}

func addTitlePage(script *ast.Script) {
	// Is there a title?
	_, exists := script.TitlePage["title"]

	// No title: no title page
	if !exists {
		return
	}

	// First get the author(s)
	authors := script.TitlePage["authors"]
	if len(authors) == 0 || authors[0].Lenght() == 0 {
		authors = script.TitlePage["author"]
	}

	// Now start building the titlepage itself
	thisPDF.AddPage()

	topPart := []styledLine{}
	topPart = append(topPart, cellify(script.TitlePage["title"], titlePageTitle)...)
	topPart = append(topPart, cellify(script.TitlePage["subtitle"], titlePageSubtitle)...)
	topPart = append(topPart, cellify(script.TitlePage["credit"], titlePageCredit)...)
	topPart = append(topPart, cellify(authors, titlePageAuthor)...)
	topPart = append(topPart, cellify(script.TitlePage["source"], titlePageSource)...)

	// Now add the top part
	addLines(topPart)

	// Build the lower half of the page
	rightLowerPart := []styledLine{}
	rightLowerPart = append(rightLowerPart, cellify(script.TitlePage["draft date"], titlePageRight)...)
	rightLowerPart = append(rightLowerPart, cellify(script.TitlePage["notes"], titlePageRight)...)
	rightLowerPart = append(rightLowerPart, cellify(script.TitlePage["contact"], titlePageRight)...)
	thisPDF.SetY(pageHeight - float64(getHeight(rightLowerPart))*em - bottomMargin)
	addLines(rightLowerPart)

	leftLowerPart := []styledLine{}
	leftLowerPart = append(leftLowerPart, cellify(script.TitlePage["copyright"], titlePageLeft)...)
	thisPDF.SetY(pageHeight - float64(getHeight(leftLowerPart))*em - bottomMargin)
	addLines(leftLowerPart)
}
