package pdf

// This file contains the definitions for a stageplay.

var stageplay = aTheme{
	action: lineStyle{
		LeadingBefore: 1,
		Indent:        10 * en,
		LineLenght:    50,
		Italics:       true,
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
		LeadingBefore: 2,
		Underline:     true,
	},
	centeredText: lineStyle{
		LeadingBefore: 1,
		Indent:        10 * en,
		LineLenght:    50,
		Italics:       true,
		Centered:      true,
	},
	looseLyrics: lineStyle{
		LeadingBefore: 1,
		LineLenght:    60,
		AllCaps:       true,
		Italics:       true,
	},
	transition: lineStyle{
		LeadingBefore: 1,
		FlushRight:    true,
	},

	character: lineStyle{
		LeadingBefore: 1,
	},
	parenthetical: lineStyle{
		FirstLineIndent: 14 * en,
		Indent:          15 * en,
		LineLenght:      45,
	},
	dialogue: lineStyle{
		Indent:     5 * en,
		LineLenght: 55,
	},
	lyrics: lineStyle{
		Indent:     5 * en,
		LineLenght: 55,
		AllCaps:    true,
		Italics:    true,
	},

	dualCharacterOne: lineStyle{
		LeadingBefore: 1,
		LineLenght:    28,
	},
	dualParentheticalOne: lineStyle{
		FirstLineIndent: 9 * en,
		Indent:          10 * en,
		LineLenght:      18,
	},
	dualDialogueOne: lineStyle{
		Indent:     5 * en,
		LineLenght: 23,
	},
	dualLyricsOne: lineStyle{
		Indent:     5 * en,
		LineLenght: 23,
		AllCaps:    true,
		Italics:    true,
	},

	dualCharacterTwo: lineStyle{
		LeadingBefore: 1,
		Indent:        32 * en,
		LineLenght:    28,
	},
	dualParentheticalTwo: lineStyle{
		FirstLineIndent: 41 * en,
		Indent:          42 * en,
		LineLenght:      17,
	},
	dualDialogueTwo: lineStyle{
		Indent:     37 * en,
		LineLenght: 28,
	},
	dualLyricsTwo: lineStyle{
		Indent:     37 * en,
		LineLenght: 28,
		AllCaps:    true,
		Italics:    true,
	},

	more: lineStyle{
		Indent: 20 * en,
	},

	titlePageTitle: lineStyle{
		LeadingBefore: 20,
		Indent:        rightMargin - leftMargin,
		Centered:      true,
	},
	titlePageSubtitle: lineStyle{
		Indent:   rightMargin - leftMargin,
		Centered: true,
	},
	titlePageCredit: lineStyle{
		LeadingBefore: 1,
		Indent:        rightMargin - leftMargin,
		Centered:      true,
	},
	titlePageAuthor: lineStyle{
		LeadingBefore: 1,
		Indent:        rightMargin - leftMargin,
		Centered:      true,
	},
	titlePageSource: lineStyle{
		LeadingBefore: 2,
		Indent:        rightMargin - leftMargin,
		Centered:      true,
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
