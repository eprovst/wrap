package html

import (
	"bytes"
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

// ToHTML returns the script in html format.
func ToHTML(script *ast.Script) string {
	styleFile := styles[strings.ToLower(script.TitlePage["type"])]

	if styleFile == "" {
		styleFile = "screen"
	}

	htm := bytes.NewBufferString(`<div class="` + styleFile + ` play">` + "\n")

	for _, element := range script.Elements {
		switch element.(type) {
		case ast.Action:
			elem := string(element.(ast.Action))
			htm.WriteString(`  <div class="action">` + string(elem) + "</div>\n")

		case ast.CenteredText:
			elem := string(element.(ast.CenteredText))
			htm.WriteString(`  <div class="action centered">` + string(elem) + "</div>\n")

		case ast.Dialogue:
			htm.WriteString(`  <div class="dialog">` + "\n")
			element := element.(ast.Dialogue)
			character := string(element.Character)
			htm.WriteString(`    <p class="character">` + character + "</p>\n")

			for _, elem := range element.Lines {
				switch elem.(type) {
				case ast.Speech:
					el := string(elem.(ast.Speech))
					htm.WriteString(`    <p>` + string(el) + "</p>\n")

				case ast.Parenthetical:
					el := string(elem.(ast.Parenthetical))
					htm.WriteString(`    <p class="parenthetical">` + string(el) + "</p>\n")

				case ast.Lyrics:
					el := string(elem.(ast.Lyrics))
					htm.WriteString(`    <p class="lyrics">` + string(el) + "</p>\n")
				}
			}
			htm.WriteString("  </div>\n")

		case ast.DualDialogue:
			htm.WriteString(`  <div class="dual">` + "\n")
			element := element.(ast.DualDialogue)

			htm.WriteString(`    <div class="left">` + "\n")
			character := string(element.LCharacter)
			htm.WriteString(`      <p class="character">` + character + "</p>\n")

			for _, elem := range element.LLines {
				switch elem.(type) {
				case ast.Speech:
					el := string(elem.(ast.Speech))
					htm.WriteString(`      <p>` + string(el) + "</p>\n")

				case ast.Parenthetical:
					el := string(elem.(ast.Parenthetical))
					htm.WriteString(`      <p class="parenthetical">` + string(el) + "</p>\n")

				case ast.Lyrics:
					el := string(elem.(ast.Lyrics))
					htm.WriteString(`      <p class="lyrics">` + string(el) + "</p>\n")
				}
			}
			htm.WriteString("    </div>\n")

			htm.WriteString(`    <div class="right">` + "\n")
			character = string(element.RCharacter)
			htm.WriteString(`      <p class="character">` + character + "</p>\n")

			for _, elem := range element.RLines {
				switch elem.(type) {
				case ast.Speech:
					el := string(elem.(ast.Speech))
					htm.WriteString(`      <p>` + string(el) + "</p>\n")

				case ast.Parenthetical:
					el := string(elem.(ast.Parenthetical))
					htm.WriteString(`      <p class="parenthetical">` + string(el) + "</p>\n")

				case ast.Lyrics:
					el := string(elem.(ast.Lyrics))
					htm.WriteString(`      <p class="lyrics">` + string(el) + "</p>\n")
				}
			}
			htm.WriteString("    </div>\n  </div>\n")

		case ast.Lyrics:
			elem := string(element.(ast.Lyrics))
			htm.WriteString(`  <div class="lyrics">` + string(elem) + "</div>\n")

		case ast.PageBreak:
			htm.WriteString(`  <div class="page-break"></div>` + "\n")

		case ast.Scene:
			elem := element.(ast.Scene)
			htm.WriteString(`  <div class="slug">`)

			if AddSceneNumbers {
				htm.WriteString(`<span class="scnuml">` + elem.SceneNumber + "</span>")
			}

			htm.WriteString(elem.Slugline)

			if AddSceneNumbers {
				htm.WriteString(`<span class="scnumr">` + elem.SceneNumber + "</span>")
			}

			htm.WriteString("</div>\n")

		case ast.Transition:
			elem := string(element.(ast.Transition))
			htm.WriteString(`  <div class="transition">` + string(elem) + "</div>\n")

		case ast.BeginAct:
			// Acts start on a new page
			htm.WriteString(`  <div class="page-break"></div>` + "\n")

			elem := string(element.(ast.BeginAct))
			htm.WriteString(`  <div class="act">` + string(elem) + "</div>\n")

		case ast.EndAct:
			elem := string(element.(ast.EndAct))
			htm.WriteString(`  <div class="act">` + string(elem) + "</div>\n")
		}
	}

	htm.WriteString("</div>\n")
	return htm.String()
}

// ToHTMLPage returns the script as html page.
func ToHTMLPage(script *ast.Script) string {
	styleFile := styles[strings.ToLower(script.TitlePage["type"])]
	if styleFile == "" {
		styleFile = "screen"
	}

	htm := bytes.NewBufferString("<!DOCTYPE html>\n")
	htm.WriteString("<html>\n")
	htm.WriteString("  <head>\n")
	htm.WriteString(`    <meta charset="utf-8">` + "\n")
	htm.WriteString("    <title>" + strings.TrimSpace(removeHTMLTags(script.TitlePage["title"])) + "</title>\n")

	for name, content := range script.TitlePage {
		if name != "title" { // Already written.
			content = strings.Replace(content, "<br>\n", "\n", -1)
			htm.WriteString(`    <meta name="` + name + `" content="` + strings.TrimSpace(content) + "\">\n")
		}
	}

	htm.WriteString(`    <link rel="stylesheet" type="text/css" href="` + URLToCSS + styleFile + `.min.css">` + "\n")
	htm.WriteString(`    <link rel="stylesheet" type="text/css" href="` + URLToCSS + `play.min.css">` + "\n")
	htm.WriteString(`    <link rel="stylesheet" type="text/css" href="` + URLToCSS + `page.min.css">` + "\n")

	htm.WriteString("  </head>\n")
	htm.WriteString("  <body>\n")

	htm.WriteString(ToHTML(script))

	htm.WriteString("  </body>\n")
	htm.WriteString("</html>\n")
	return htm.String()
}
