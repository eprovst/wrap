package parser

import (
	"testing"

	"github.com/Wraparound/wrap/pkg/ast"
)

/* FOUNTAIN MODE TESTS */

/* Many of the tests in this section are based on the ones by Nima Yousefi and
   John August, what follows is the license notice of their tests.

    Copyright (c) 2013 Nima Yousefi & John August

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to
    deal in the Software without restriction, including without limitation the
    rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
    sell copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in
    all copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
    FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
    IN THE SOFTWARE.
*/

func TestBoneyard(t *testing.T) {
	UseWrapExtensions = false

	input := `A line of action.

/* This is an inline Boneyard. */

/*
This is a
multi-line
Boneyard.
*/

This is an /* internal */ boneyard.`

	output := scriptFromElements([]ast.Element{
		ast.Action(textHandler([]string{"A line of action.", ""})), // Why aren't these merged?
		ast.Action(textHandler([]string{"This is an  boneyard."})),
	})

	assertMatch(t, input, output)
}

func TestCenteredText(t *testing.T) {
	UseWrapExtensions = false

	input := `> Centered <

> Not centered

>No space<

> Lots of Space <`

	output := scriptFromElements([]ast.Element{
		ast.CenteredText(textHandler([]string{"Centered"})),
		ast.Transition(textHandler([]string{"Not centered"})),
		ast.CenteredText(textHandler([]string{"No space", "",
			"Lots of Space"})),
	})

	assertMatch(t, input, output)
}

func TestDialogue(t *testing.T) {
	UseWrapExtensions = false

	input := `ADAM
I like to write.

EVE (O.S.)
What kind of writing?

ADAM
(nervous)
Screenwriting.

EVE
ME TOO!

EVE (cont'd)
I think screenwriting is the best.

ADAM
That's great.
I was really worried you wouldn't
like screenwriting.

EVE
Oh, Adam.
(kissing him)

ADAM
I love you.

EVE ^
You're so stupid.

R2D2
Bleep Boop.
(*This is a valid character cue.*)

23 (O.S.)
Character name must include a letter

ADAM
That was really weird.
I can be weird to. Here's
a space in the dialogue
block
  
for absolutely no good
reason.

		EVE
	  (feeling cocky)
Well, then I'll just indent!`

	output := scriptFromElements([]ast.Element{
		ast.Dialogue{
			Character: textHandler([]string{"ADAM"}),
			Lines: []ast.Element{
				ast.Speech(textHandler([]string{"I like to write."})),
			},
		},
		ast.Dialogue{
			Character: textHandler([]string{"EVE (O.S.)"}),
			Lines: []ast.Element{
				ast.Speech(textHandler([]string{"What kind of writing?"})),
			},
		},
		ast.Dialogue{
			Character: textHandler([]string{"ADAM"}),
			Lines: []ast.Element{
				ast.Parenthetical(textHandler([]string{"(nervous)"})),
				ast.Speech(textHandler([]string{"Screenwriting."})),
			},
		},
		ast.Dialogue{
			Character: textHandler([]string{"EVE"}),
			Lines: []ast.Element{
				ast.Speech(textHandler([]string{"ME TOO!"})),
			},
		},
		ast.Dialogue{
			Character: textHandler([]string{"EVE (cont'd)"}),
			Lines: []ast.Element{
				ast.Speech(textHandler([]string{
					"I think screenwriting is the best."})),
			},
		},
		ast.Dialogue{
			Character: textHandler([]string{"ADAM"}),
			Lines: []ast.Element{
				ast.Speech(textHandler([]string{"That's great.",
					"I was really worried you wouldn't",
					"like screenwriting."})),
			},
		},
		ast.Dialogue{
			Character: textHandler([]string{"EVE"}),
			Lines: []ast.Element{
				ast.Speech(textHandler([]string{"Oh, Adam."})),
				ast.Parenthetical(textHandler([]string{"(kissing him)"})),
			},
		},
		ast.DualDialogue{
			LCharacter: textHandler([]string{"ADAM"}),
			LLines: []ast.Element{
				ast.Speech(textHandler([]string{"I love you."})),
			},
			RCharacter: textHandler([]string{"EVE"}),
			RLines: []ast.Element{
				ast.Speech(textHandler([]string{"You're so stupid."})),
			},
		},
		ast.Dialogue{
			Character: textHandler([]string{"R2D2"}),
			Lines: []ast.Element{
				ast.Speech(textHandler([]string{"Bleep Boop."})),
				ast.Parenthetical(textHandler([]string{
					"(*This is a valid character cue.*)"})),
			},
		},
		ast.Action(textHandler([]string{"23 (O.S.)",
			"Character name must include a letter"})),
		ast.Dialogue{
			Character: textHandler([]string{"ADAM"}),
			Lines: []ast.Element{
				ast.Speech(textHandler([]string{"That was really weird.",
					"I can be weird to. Here's", "a space in the dialogue",
					"block", "", "for absolutely no good", "reason."})),
			},
		},
		ast.Dialogue{
			Character: textHandler([]string{"EVE"}),
			Lines: []ast.Element{
				ast.Parenthetical(textHandler([]string{"(feeling cocky)"})),
				ast.Speech(textHandler([]string{
					"Well, then I'll just indent!"})),
			},
		},
	})

	assertMatch(t, input, output)
}

func TestDualDialogue(t *testing.T) {
	UseWrapExtensions = false

	input := `ADAM
Yes.

EVE ^
No.`

	output := scriptFromElements([]ast.Element{
		ast.DualDialogue{
			LCharacter: textHandler([]string{"ADAM"}),
			LLines: []ast.Element{
				ast.Speech(textHandler([]string{"Yes."})),
			},
			RCharacter: textHandler([]string{"EVE"}),
			RLines: []ast.Element{
				ast.Speech(textHandler([]string{"No."})),
			},
		},
	})

	assertMatch(t, input, output)
}

func TestForced(t *testing.T) {
	UseWrapExtensions = false

	input := `!BANG
BANG
BANG

@McDUCK
I'm vegan.

SINGER
~These are the songs
~That I sing.`

	output := scriptFromElements([]ast.Element{
		ast.Action(textHandler([]string{"BANG", "BANG", "BANG"})),
		ast.Dialogue{
			Character: textHandler([]string{"McDUCK"}),
			Lines: []ast.Element{
				ast.Speech(textHandler([]string{"I'm vegan."})),
			},
		},
		ast.Dialogue{
			Character: textHandler([]string{"SINGER"}),
			Lines: []ast.Element{
				ast.Lyrics(textHandler([]string{"These are the songs",
					"That I sing."})),
			},
		},
	})

	assertMatch(t, input, output)
}

func TestIndenting(t *testing.T) {
	UseWrapExtensions = false

	input := `                CUT TO:

INT. GARAGE - DAY

BRICK and STEEL get into Mom's PORSCHE, Steel at the wheel.  They
pause for a beat, the gravity of the situation catching up with
them.

            BRICK
    This is everybody we've ever put away.

            STEEL
        (starting the engine)
    So much for retirement!

They speed off. To destiny!`

	output := scriptFromElements([]ast.Element{
		ast.Transition(textHandler([]string{"CUT TO:"})),
		ast.Scene{
			Slugline:    textHandler([]string{"INT. GARAGE - DAY"}),
			SceneNumber: "1",
		},
		ast.Action(textHandler([]string{
			"BRICK and STEEL get into Mom's PORSCHE, Steel at the wheel.  They",
			"pause for a beat, the gravity of the situation catching up with",
			"them."})),
		ast.Dialogue{
			Character: textHandler([]string{"BRICK"}),
			Lines: []ast.Element{
				ast.Speech(textHandler([]string{
					"This is everybody we've ever put away."})),
			},
		},
		ast.Dialogue{
			Character: textHandler([]string{"STEEL"}),
			Lines: []ast.Element{
				ast.Parenthetical(textHandler([]string{
					"(starting the engine)"})),
				ast.Speech(textHandler([]string{"So much for retirement!"})),
			},
		},
		ast.Action(textHandler([]string{"They speed off. To destiny!"})),
	})

	assertMatch(t, input, output)
}

func TestMultilineAction(t *testing.T) {
	UseWrapExtensions = false

	input := `A normal line of action.

This is a multi-line block of Action.
Made famous by the great
WALTER HILL
but used by many other writers.

A normal line of action.`

	output := scriptFromElements([]ast.Element{
		ast.Action(textHandler([]string{
			"A normal line of action.",
			"",
			"This is a multi-line block of Action.",
			"Made famous by the great",
			"WALTER HILL",
			"but used by many other writers.",
			"",
			"A normal line of action."})),
	})

	assertMatch(t, input, output)
}

func TestNoteSections(t *testing.T) {
	UseWrapExtensions = false

	input := `A line.

[[A note.]]

[[This line spans
  multiple lines.]]

This is an [[internal]] note.`

	output := scriptFromElements([]ast.Element{
		ast.Action(textHandler([]string{"A line."})),
		ast.Note(textHandler([]string{"[[A note.]]"})),
		ast.Note(textHandler([]string{"[[This line spans",
			"  multiple lines.]]"})),
		ast.Action(textHandler([]string{"This is an [[internal]] note."})),
	})

	assertMatch(t, input, output)
}

func TestPageBreaks(t *testing.T) {
	UseWrapExtensions = false

	input := `Page one.

===

Page two.

=========

Page three.`

	output := scriptFromElements([]ast.Element{
		ast.Action(textHandler([]string{"Page one."})),
		ast.PageBreak{},
		ast.Action(textHandler([]string{"Page two."})),
		ast.PageBreak{},
		ast.Action(textHandler([]string{"Page three."})),
	})

	assertMatch(t, input, output)
}

func TestSceneHeaders(t *testing.T) {
	UseWrapExtensions = false

	input := `INT. HOUSE - DAY

EXT. HOUSE - DAY

INT HOUSE DAY

EXT HOUSE DAY

INT/EXT. HOUSE - DAY

INT./EXT. HOUSE - DAY

I/E. HOUSE - DAY

I./E. HOUSE - DAY

I/E HOUSE DAY

INT - HOUSE - DAY

EXT - HOUSE - DAY

EST. HOUSE - DAY

EST HOUSE DAY

INT. HOUSE - DAY (1979)

int. house - day

.KITCHEN

... not a scene header.

INT. HOUSE - DAY
EXT. HOUSE - DAY

ESTABLISHING`

	output := scriptFromElements([]ast.Element{
		ast.Scene{
			Slugline:    textHandler([]string{"INT. HOUSE - DAY"}),
			SceneNumber: "1",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"EXT. HOUSE - DAY"}),
			SceneNumber: "2",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"INT HOUSE DAY"}),
			SceneNumber: "3",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"EXT HOUSE DAY"}),
			SceneNumber: "4",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"INT/EXT. HOUSE - DAY"}),
			SceneNumber: "5",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"INT./EXT. HOUSE - DAY"}),
			SceneNumber: "6",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"I/E. HOUSE - DAY"}),
			SceneNumber: "7",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"I./E. HOUSE - DAY"}),
			SceneNumber: "8",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"I/E HOUSE DAY"}),
			SceneNumber: "9",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"INT - HOUSE - DAY"}),
			SceneNumber: "10",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"EXT - HOUSE - DAY"}),
			SceneNumber: "11",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"EST. HOUSE - DAY"}),
			SceneNumber: "12",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"EST HOUSE DAY"}),
			SceneNumber: "13",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"INT. HOUSE - DAY (1979)"}),
			SceneNumber: "14",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"int. house - day"}),
			SceneNumber: "15",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"KITCHEN"}),
			SceneNumber: "16",
		},
		ast.Action(textHandler([]string{
			"... not a scene header.",
			"",
			"INT. HOUSE - DAY",
			"EXT. HOUSE - DAY",
			"",
			"ESTABLISHING",
		})),
	})

	assertMatch(t, input, output)
}

func TestSceneNumbers(t *testing.T) {
	UseWrapExtensions = false

	input := `INT. HOUSE - DAY #1#

INT. HOUSE - DAY #1A#

INT. HOUSE - DAY #1a#

INT. HOUSE - DAY #A1#

INT. HOUSE - DAY #I-1-A#

INT. HOUSE - DAY #1.#

INT. HOUSE - DAY - FLASHBACK (1944) #110A#

` // <-- Has to be followed by an empty line

	output := scriptFromElements([]ast.Element{
		ast.Scene{
			Slugline:    textHandler([]string{"INT. HOUSE - DAY"}),
			SceneNumber: "1",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"INT. HOUSE - DAY"}),
			SceneNumber: "1A",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"INT. HOUSE - DAY"}),
			SceneNumber: "1a",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"INT. HOUSE - DAY"}),
			SceneNumber: "A1",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"INT. HOUSE - DAY"}),
			SceneNumber: "I-1-A",
		},
		ast.Scene{
			Slugline:    textHandler([]string{"INT. HOUSE - DAY"}),
			SceneNumber: "1.",
		},
		ast.Scene{
			Slugline: textHandler([]string{
				"INT. HOUSE - DAY - FLASHBACK (1944)",
			}),
			SceneNumber: "110A",
		},
	})

	assertMatch(t, input, output)
}

func TestSections(t *testing.T) {
	UseWrapExtensions = false
	input := `# Act One

#Act Two

#ACT THREE

# ACT FOUR

# 5

## Level 2

### Level 3`

	output := scriptFromElements([]ast.Element{
		ast.Section{
			Line:  textHandler([]string{"Act One"}),
			Level: 1,
		},
		ast.Section{
			Line:  textHandler([]string{"Act Two"}),
			Level: 1,
		},
		ast.Section{
			Line:  textHandler([]string{"ACT THREE"}),
			Level: 1,
		},
		ast.Section{
			Line:  textHandler([]string{"ACT FOUR"}),
			Level: 1,
		},
		ast.Section{
			Line:  textHandler([]string{"5"}),
			Level: 1,
		},
		ast.Section{
			Line:  textHandler([]string{"Level 2"}),
			Level: 2,
		},
		ast.Section{
			Line:  textHandler([]string{"Level 3"}),
			Level: 3,
		},
	})

	assertMatch(t, input, output)
}

func TestComplexSection(t *testing.T) {
	UseWrapExtensions = false

	input := `# Act 1
= Synopsis 1

INT. LOCATION - DAY

# Act 2
## Sequence 1`

	output := scriptFromElements([]ast.Element{
		ast.Section{
			Line:  textHandler([]string{"Act 1"}),
			Level: 1,
		},
		ast.Synopsis(textHandler([]string{"Synopsis 1"})),
		ast.Scene{
			Slugline:    textHandler([]string{"INT. LOCATION - DAY"}),
			SceneNumber: "1",
		},
		ast.Section{
			Line:  textHandler([]string{"Act 2"}),
			Level: 1,
		},
		ast.Section{
			Line:  textHandler([]string{"Sequence 1"}),
			Level: 2,
		},
	})

	assertMatch(t, input, output)
}

func TestSynopsis(t *testing.T) {
	UseWrapExtensions = false

	input := `# Act One

= This is the first act.

# Act Two

=This is another act.

# Act Three

=THE FINAL ACT`

	output := scriptFromElements([]ast.Element{
		ast.Section{
			Line:  textHandler([]string{"Act One"}),
			Level: 1,
		},
		ast.Synopsis(textHandler([]string{"This is the first act."})),
		ast.Section{
			Line:  textHandler([]string{"Act Two"}),
			Level: 1,
		},
		ast.Synopsis(textHandler([]string{"This is another act."})),
		ast.Section{
			Line:  textHandler([]string{"Act Three"}),
			Level: 1,
		},
		ast.Synopsis(textHandler([]string{"THE FINAL ACT"})),
	})

	assertMatch(t, input, output)
}

func TestTransitions(t *testing.T) {
	UseWrapExtensions = false

	input := `

CUT TO:

SMASH CUT TO:

FADE TO BLACK.

> NOT A STANDARD TRANSITION:

CUT TO:  

GO TO:

TO:`

	output := scriptFromElements([]ast.Element{
		ast.Transition(textHandler([]string{"CUT TO:"})),
		ast.Transition(textHandler([]string{"SMASH CUT TO:"})),
		ast.Action(textHandler([]string{"FADE TO BLACK."})),
		ast.Transition(textHandler([]string{"NOT A STANDARD TRANSITION:"})),
		ast.Action(textHandler([]string{"CUT TO:"})),
		ast.Transition(textHandler([]string{"GO TO:"})),
		ast.Action(textHandler([]string{"TO:"})),
	})

	assertMatch(t, input, output)
}
