package html

import (
	"io"
	"strings"
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

	_, err := out.Write([]byte(strings.Repeat("  ", int(indent)) + line))

	if err != nil {
		panic(err)
	}
}

func writeLines(lines []string, out io.Writer) {
	for i, line := range lines {
		if i != len(lines)-1 {
			writeString(line+"<br>\n", out)

		} else {
			writeString(line, out)
		}
	}
}
