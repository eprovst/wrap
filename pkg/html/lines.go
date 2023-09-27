package html

import (
	"html"
	"io"
	"strings"

	"github.com/eprovst/wrap/pkg/ast"
)

// Indent keeps track of the current indent level.
var indent uint8

func writeText(text string, out io.Writer) {
	for _, line := range strings.SplitAfter(text, "\n") {
		writeString(line, out)
	}
}

func writeString(line string, out io.Writer) {
	if len(line) == 0 {
		return
	}

	out.Write([]byte(strings.Repeat("  ", int(indent)) + line))
}

func writeLines(lines []ast.Line, out io.Writer) {
	for i, line := range lines {
		writeLine(line, out)

		if i != len(lines)-1 || line.Empty() {
			out.Write([]byte("<br>"))
		}

		out.Write([]byte("\n"))
	}
}

func writeLine(line ast.Line, out io.Writer) {
	out.Write([]byte(strings.Repeat("  ", int(indent))))

	for _, cell := range line {
		writeCell(cell, out)
	}
}

func writeCell(cell ast.Cell, out io.Writer) {
	if cell.Comment {
		out.Write([]byte("<ins>"))
	}
	if cell.Boldface {
		out.Write([]byte("<b>"))
	}
	if cell.Italics {
		out.Write([]byte("<i>"))
	}
	if cell.Underline {
		out.Write([]byte("<u>"))
	}

	contents := cell.Content

	// Save multiple spaces.
	contents = html.EscapeString(contents)
	contents = strings.Replace(contents, "  ", "&nbsp; ", -1)
	contents = strings.Replace(contents, " &nbsp;", "&nbsp;&nbsp;", -1)
	contents = strings.Replace(contents, "  ", "&nbsp; ", -1)

	out.Write([]byte(contents))

	if cell.Underline {
		out.Write([]byte("</u>"))
	}
	if cell.Italics {
		out.Write([]byte("</i>"))
	}
	if cell.Boldface {
		out.Write([]byte("</b>"))
	}
	if cell.Comment {
		out.Write([]byte("</ins>"))
	}
}
