package parser

import (
	"bytes"
	"regexp"
	"strings"
)

// In this file we define functions to parse lines a bit further.

var sceneNumberRegex = regexp.MustCompile("^(.+?)(?:\\s*)(?:#(.+)#)?$")

func parseSceneHeading(line string) (slugline, scenenumber string) {
	// TODO: Find way to do this without regex...
	vals := sceneNumberRegex.FindStringSubmatch(line)
	return vals[1], vals[2]
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
