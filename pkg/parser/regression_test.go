package parser

import (
	"testing"

	"github.com/eprovst/wrap/pkg/ast"
)

// Bug found and example given by Paul W. Rankin
func TestDialogueAfterNote(t *testing.T) {
	input := `ALFRED
Tea sir?

[[ something
noteworthy
indeed ]]

BRUCE
Bring me the bat tea.`

	output := scriptFromElements([]ast.Element{
		ast.Dialogue{
			Character: textHandler([]string{"ALFRED"}),
			Lines:     []ast.Element{ast.Speech(textHandler([]string{"Tea sir?"}))},
		},
		ast.Note(textHandler([]string{"[[ something", "noteworthy", "indeed ]]"})),
		ast.Dialogue{
			Character: textHandler([]string{"BRUCE"}),
			Lines:     []ast.Element{ast.Speech(textHandler([]string{"Bring me the bat tea."}))},
		},
	})

	assertMatch(t, input, output)
}
