package html

import (
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
	styleFile := styles[strings.ToLower(script.TitlePage["type"])]

	if styleFile == "" {
		styleFile = "screen"
	}

	writeText(`<div class="`+styleFile+` play">`+"\n", writer)
	indent++

	for _, element := range script.Elements {
		switch element := element.(type) {
		case ast.Action:
			writeText(`<div class="action">`+string(element)+"</div>\n", writer)

		case ast.CenteredText:
			writeText(`<div class="action centered">`+string(element)+"</div>\n", writer)

		case ast.Dialogue:
			writeString(`<div class="dialog">`+"\n", writer)
			indent++

			writeText(`<p class="character">`+element.Character+"</p>\n", writer)

			for _, elem := range element.Lines {
				switch elem := elem.(type) {
				case ast.Speech:
					writeText(`<p>`+string(elem)+"</p>\n", writer)

				case ast.Parenthetical:
					writeText(`<p class="parenthetical">`+string(elem)+"</p>\n", writer)

				case ast.Lyrics:
					writeText(`<p class="lyrics">`+string(elem)+"</p>\n", writer)
				}
			}

			indent--
			writeString("</div>\n", writer)

		case ast.DualDialogue:
			writeString(`<div class="dual">`+"\n", writer)
			indent++
			writeString(`<div class="left">`+"\n", writer)
			indent++
			writeText(`<p class="character">`+element.LCharacter+"</p>\n", writer)

			for _, elem := range element.LLines {
				switch elem := elem.(type) {
				case ast.Speech:
					writeText(`<p>`+string(elem)+"</p>\n", writer)

				case ast.Parenthetical:
					writeText(`<p class="parenthetical">`+string(elem)+"</p>\n", writer)

				case ast.Lyrics:
					writeText(`<p class="lyrics">`+string(elem)+"</p>\n", writer)
				}
			}

			indent--
			writeString("</div>\n", writer)

			writeString(`<div class="right">`+"\n", writer)
			indent++
			writeText(`<p class="character">`+element.RCharacter+"</p>\n", writer)

			for _, elem := range element.RLines {
				switch elem := elem.(type) {
				case ast.Speech:
					writeText(`<p>`+string(elem)+"</p>\n", writer)

				case ast.Parenthetical:
					writeText(`<p class="parenthetical">`+string(elem)+"</p>\n", writer)

				case ast.Lyrics:
					writeText(`<p class="lyrics">`+string(elem)+"</p>\n", writer)
				}
			}

			indent--
			writeString("</div>\n", writer)
			indent--
			writeString("</div>\n", writer)

		case ast.Lyrics:
			writeText(`<div class="lyrics">`+string(element)+"</div>\n", writer)

		case ast.PageBreak:
			writeString(`<div class="page-break"></div>`+"\n", writer)

		case ast.Scene:
			writeString(`<div class="slug">`+"\n", writer)
			indent++

			if AddSceneNumbers {
				writeString(`<span class="scnuml">`+element.SceneNumber+"</span>\n", writer)
			}

			writeText(element.Slugline+"\n", writer)

			if AddSceneNumbers {
				writeString(`<span class="scnumr">`+element.SceneNumber+"</span>\n", writer)
			}

			indent--
			writeString("</div>\n", writer)

		case ast.Transition:
			writeText(`<div class="transition">`+string(element)+"</div>\n", writer)

		case ast.BeginAct:
			// Acts start on a new page
			writeString(`<div class="page-break"></div>`+"\n", writer)
			writeText(`<div class="act">`+string(element)+"</div>\n", writer)

		case ast.EndAct:
			writeText(`<div class="act">`+string(element)+"</div>\n", writer)
		}
	}

	indent--
	writeString("</div>\n", writer)
}

// WriteHTMLPage writes the script as html page to a *io.Writer.
func WriteHTMLPage(script *ast.Script, writer io.Writer) {
	styleFile := styles[strings.ToLower(script.TitlePage["type"])]
	if styleFile == "" {
		styleFile = "screen"
	}

	writeString("<!DOCTYPE html>\n", writer)
	writeString("<html>\n", writer)
	indent++

	writeString("<head>\n", writer)
	indent++
	writeString(`<meta charset="utf-8">`+"\n", writer)
	writeString("<title>"+strings.TrimSpace(removeHTMLTags(script.TitlePage["title"]))+"</title>\n", writer)

	for name, content := range script.TitlePage {
		if name != "title" { // Already written.
			content = strings.Replace(content, "<br>\n", "\n", -1)
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
