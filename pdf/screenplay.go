package pdf

// This file contains the definitions for a screenplay.

var screenplay = aTheme{
	action: lineStyle{
		LeadingBefore: 1,
		LineLenght:    60,
	},
	act: lineStyle{
		LeadingBefore: 2,
		LineLenght:    35,
		Indent:        25 * en,
	},
	slugLine: lineStyle{
		LineLenght:    60,
		LeadingBefore: 2,
		AllCaps:       true,
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
		LeadingBefore: 1,
		LineLenght:    15,
		Indent:        45 * en,
	},

	character: lineStyle{
		LeadingBefore: 1,
		Indent:        22 * en,
		LineLenght:    38,
	},
	parenthetical: lineStyle{
		FirstLineIndent: 16 * en,
		Indent:          17 * en,
		LineLenght:      24,
	},
	dialogue: lineStyle{
		Indent:     10 * en,
		LineLenght: 40,
	},
	lyrics: lineStyle{
		Indent:     10 * en,
		LineLenght: 40,
		Italics:    true,
	},

	dualCharacterOne: lineStyle{
		LeadingBefore: 1,
		Indent:        12 * en,
		LineLenght:    17,
	},
	dualParentheticalOne: lineStyle{
		FirstLineIndent: 2 * en,
		Indent:          3 * en,
		LineLenght:      17,
	},
	dualDialogueOne: lineStyle{
		LineLenght: 29,
	},
	dualLyricsOne: lineStyle{
		LineLenght: 29,
		Italics:    true,
	},

	dualCharacterTwo: lineStyle{
		LeadingBefore: 1,
		Indent:        43 * en,
		LineLenght:    17,
	},
	dualParentheticalTwo: lineStyle{
		FirstLineIndent: 33 * en,
		Indent:          34 * en,
		LineLenght:      17,
	},
	dualDialogueTwo: lineStyle{
		Indent:     31 * en,
		LineLenght: 29,
	},
	dualLyricsTwo: lineStyle{
		Indent:     31 * en,
		LineLenght: 29,
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
		LineLenght:    60,
	},
	titlePageLeft: lineStyle{
		LeadingBefore: 2,
		LineLenght:    60,
	},
}
