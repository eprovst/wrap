package pdf

// This file contains the definitions for a stageplay.

var stageplay = aTheme{
	action: lineStyle{
		Indent:        12.5 * en,
		LineLenght:    33,
		LeadingBefore: 1,
	},
	act: lineStyle{
		LeadingBefore: 0,
		LineLenght:    35,
		Indent:        25 * en,
		AllCaps:       true,
		Underline:     true,
	},
	slugLine: lineStyle{
		Indent:        25 * en,
		LineLenght:    35,
		LeadingBefore: 1,
		Underline:     true,
	},
	centeredText: lineStyle{
		LeadingBefore: 1,
		LineLenght:    60,
		Centered:      true,
	},
	looseLyrics: lineStyle{
		LeadingBefore: 1,
		LineLenght:    60,
		Italics:       true,
	},
	transition: lineStyle{
		Indent:        25 * en,
		LineLenght:    35,
		LeadingBefore: 1,
		AllCaps:       true,
	},

	character: lineStyle{
		Indent:        25 * en,
		LineLenght:    35,
		LeadingBefore: 1,
	},
	parenthetical: lineStyle{
		Indent:     12.5 * en,
		LineLenght: 33,
	},
	dialogue: lineStyle{
		LineLenght: 60,
	},
	lyrics: lineStyle{
		LineLenght: 60,
		Italics:    true,
	},

	dualCharacterOne: lineStyle{
		LeadingBefore: 1,
		Indent:        8 * en,
		LineLenght:    20,
	},
	dualParentheticalOne: lineStyle{
		Indent:     3.5 * en,
		LineLenght: 18,
	},
	dualDialogueOne: lineStyle{
		LineLenght: 28,
	},
	dualLyricsOne: lineStyle{
		LineLenght: 28,
		Italics:    true,
	},

	dualCharacterTwo: lineStyle{
		LeadingBefore: 1,
		Indent:        40 * en,
		LineLenght:    20,
	},
	dualParentheticalTwo: lineStyle{
		Indent:     35.5 * en,
		LineLenght: 18,
	},
	dualDialogueTwo: lineStyle{
		Indent:     32 * en,
		LineLenght: 28,
	},
	dualLyricsTwo: lineStyle{
		Indent:     32 * en,
		LineLenght: 28,
		Italics:    true,
	},

	more: lineStyle{
		Indent: 25 * en,
	},

	titlePageTitle: lineStyle{
		LeadingBefore: 15,
		Indent:        25 * en,
		LineLenght:    35,
		Underline:     true,
		AllCaps:       true,
	},
	titlePageSubtitle: lineStyle{
		LeadingBefore: 1,
		Indent:        25 * en,
		LineLenght:    35,
	},
	titlePageCredit: lineStyle{
		LeadingBefore: 1,
		Indent:        25 * en,
		LineLenght:    35,
	},
	titlePageAuthor: lineStyle{
		LeadingBefore: 1,
		Indent:        25 * en,
		LineLenght:    35,
	},
	titlePageSource: lineStyle{
		LeadingBefore: 1,
		Indent:        25 * en,
		LineLenght:    35,
	},

	titlePageRight: lineStyle{
		LeadingBefore: 2,
		FlushRight:    true,
		LineLenght:    30,
	},
	titlePageLeft: lineStyle{
		LeadingBefore: 2,
		LineLenght:    30,
	},
}
