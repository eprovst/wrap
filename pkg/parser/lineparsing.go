package parser

import (
	"bytes"
	"strings"
)

// In this file we define functions to parse lines a bit further.

func parseSceneHeading(line string) (slugline, scenenumber string) {
	vals := strings.Split(line, "#")

	if len(vals) != 3 || len(vals[2]) != 0 {
		return line, ""
	}

	return strings.TrimRight(vals[0], " \t"), vals[1]
}

func parseSection(line string) (sect string, level uint8) {
	buffer := bytes.NewBufferString(line)

	for {
		chr, _, err := buffer.ReadRune()

		if err != nil {
			return "", level
		}

		switch chr {
		case '#':
			level++

		default:
			buffer.UnreadRune()
			line = buffer.String()
			line = strings.TrimSpace(line)
			return line, level
		}
	}
}
