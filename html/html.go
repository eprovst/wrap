package html

import (
	"html"
	"io"
	"strings"

	"github.com/Wraparound/wrap/ast"
)

// AddSceneNumbers makes the export module add scene numbers
var AddSceneNumbers = true

// URLToCSS contains the path to the Wraparound (or your) CSS.
var URLToCSS = "https://cdn.jsdelivr.net/gh/wraparound/css@1.0/"

var styles = map[string]string{
	"screenplay": "screen",
	"stageplay":  "stage",
}

// WriteHTML writes a script in HTML format to a *io.Writer
func WriteHTML(script *ast.Script, writer io.Writer) {
	richStyle := script.TitlePage["type"]
	var styleFile string
	if len(richStyle) != 0 {
		styleFile = styles[strings.ToLower(richStyle[0].String())]
	}

	if styleFile == "" {
		styleFile = "screen"
	}

	writeText(`<div class="`+styleFile+` play">`+"\n", writer)
	indent++

	for idx, element := range script.Elements {
		switch element := element.(type) {
		case ast.Action:
			writeText(`<div class="action">`+"\n", writer)
			indent++
			writeLines(element, writer)
			indent--
			writeText("</div>\n", writer)

		case ast.CenteredText:
			writeText(`<div class="action centered">`+"\n", writer)
			indent++
			writeLines(element, writer)
			indent--
			writeText("</div>\n", writer)

		case ast.Dialogue:
			writeString(`<div class="dialog">`+"\n", writer)
			indent++

			writeText(`<p class="character">`+"\n", writer)
			indent++
			writeLines(element.Character, writer)
			indent--
			writeText("</p>\n", writer)

			for _, elem := range element.Lines {
				switch elem := elem.(type) {
				case ast.Speech:
					writeText("<p>\n", writer)
					indent++
					writeLines(elem, writer)
					indent--
					writeText("</p>\n", writer)

				case ast.Parenthetical:
					writeText(`<p class="parenthetical">`+"\n", writer)
					indent++
					writeLines(elem, writer)
					indent--
					writeText("</p>\n", writer)

				case ast.Lyrics:
					writeText(`<p class="lyrics">`+"\n", writer)
					indent++
					writeLines(elem, writer)
					indent--
					writeText("</p>\n", writer)
				}
			}

			indent--
			writeString("</div>\n", writer)

		case ast.DualDialogue:
			writeString(`<div class="dual">`+"\n", writer)
			indent++
			writeString(`<div class="left">`+"\n", writer)
			indent++
			writeText(`<p class="character">`+"\n", writer)
			indent++
			writeLines(element.LCharacter, writer)
			indent--
			writeText("</p>\n", writer)

			for _, elem := range element.LLines {
				switch elem := elem.(type) {
				case ast.Speech:
					writeText(`<p>`+"\n", writer)
					indent++
					writeLines(elem, writer)
					indent--
					writeText("</p>\n", writer)

				case ast.Parenthetical:
					writeText(`<p class="parenthetical">`+"\n", writer)
					indent++
					writeLines(elem, writer)
					indent--
					writeText("</p>\n", writer)

				case ast.Lyrics:
					writeText(`<p class="lyrics">`+"\n", writer)
					indent++
					writeLines(elem, writer)
					indent--
					writeText("</p>\n", writer)
				}
			}

			indent--
			writeString("</div>\n", writer)

			writeString(`<div class="right">`+"\n", writer)
			indent++
			writeText(`<p class="character">`+"\n", writer)
			indent++
			writeLines(element.RCharacter, writer)
			indent--
			writeText("</p>\n", writer)

			for _, elem := range element.RLines {
				switch elem := elem.(type) {
				case ast.Speech:
					writeText("<p>\n", writer)
					indent++
					writeLines(elem, writer)
					indent--
					writeText("</p>\n", writer)

				case ast.Parenthetical:
					writeText(`<p class="parenthetical">`+"\n", writer)
					indent++
					writeLines(elem, writer)
					indent--
					writeText("</p>\n", writer)

				case ast.Lyrics:
					writeText(`<p class="lyrics">`+"\n", writer)
					indent++
					writeLines(elem, writer)
					indent--
					writeText("</p>\n", writer)
				}
			}

			indent--
			writeString("</div>\n", writer)
			indent--
			writeString("</div>\n", writer)

		case ast.Lyrics:
			writeText(`<div class="lyrics">`+"\n", writer)
			indent++
			writeLines(element, writer)
			indent--
			writeText("</div>\n", writer)

		case ast.PageBreak:
			writeString(`<div class="page-break"></div>`+"\n", writer)

		case ast.Scene:
			writeString(`<div class="slug">`+"\n", writer)
			indent++

			if AddSceneNumbers {
				writeString(`<span class="scnuml">`+element.SceneNumber+"</span>\n", writer)
			}

			writeLines(element.Slugline, writer)

			if AddSceneNumbers {
				writeString(`<span class="scnumr">`+element.SceneNumber+"</span>\n", writer)
			}

			indent--
			writeString("</div>\n", writer)

		case ast.Transition:
			writeText(`<div class="transition">`+"\n", writer)
			indent++
			writeLines(element, writer)
			indent--
			writeText("</div>\n", writer)

		case ast.BeginAct:
			// Acts start on a new page if it isn't the first element in the play.
			if idx != 0 {
				writeString(`<div class="page-break"></div>`+"\n", writer)
			}

			writeText(`<div class="act">`+"\n", writer)
			indent++
			writeLines(element, writer)
			indent--
			writeText("</div>\n", writer)

		case ast.EndAct:
			writeText(`<div class="act">`+"\n", writer)
			indent++
			writeLines(element, writer)
			indent--
			writeText("</div>\n", writer)
		}
	}

	indent--
	writeString("</div>\n", writer)
}

// WriteHTMLPage writes the script as html page to a *io.Writer.
func WriteHTMLPage(script *ast.Script, writer io.Writer) {
	richStyle := script.TitlePage["type"]
	var styleFile string
	if len(richStyle) != 0 {
		styleFile = styles[strings.ToLower(richStyle[0].String())]
	}

	if styleFile == "" {
		styleFile = "screen"
	}

	writeString("<!DOCTYPE html>\n", writer)
	writeString("<html>\n", writer)
	indent++

	writeString("<head>\n", writer)
	indent++
	writeString(`<meta charset="utf-8">`+"\n", writer)
	writeString("<title>"+strings.TrimSpace(ast.LinesToString(script.TitlePage["title"]))+"</title>\n", writer)

	for name, richcontent := range script.TitlePage {
		if name != "title" { // Already written.
			content := html.EscapeString(ast.LinesToString(richcontent))
			writeText(`<meta name="`+name+`" content="`+strings.TrimSpace(content)+"\">\n", writer)
		}
	}

	writeString(`<link rel="stylesheet" type="text/css" href="`+URLToCSS+styleFile+`.min.css">`+"\n", writer)
	writeString(`<link rel="stylesheet" type="text/css" href="`+URLToCSS+`play.min.css">`+"\n", writer)
	writeString(`<link rel="stylesheet" type="text/css" href="`+URLToCSS+`page.min.css">`+"\n", writer)

	indent--
	writeString("</head>\n", writer)

	writeString("<body>\n", writer)
	indent++
	WriteHTML(script, writer)
	indent--
	writeString("</body>\n", writer)

	indent--
	writeString("</html>\n", writer)
}
