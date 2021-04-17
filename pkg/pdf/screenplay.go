package pdf

// This file contains the definitions for a screenplay.

var screenplay = aTheme{
	action: lineStyle{
		LeadingBefore: 1,
		LineLength:    60,
	},
	act: lineStyle{
		LeadingBefore: 0,
		LineLength:    35,
		Indent:        25 * en,
	},
	slugLine: lineStyle{
		LineLength:    60,
		LeadingBefore: 2,
		AllCaps:       true,
	},
	centeredText: lineStyle{
		LeadingBefore: 1,
		LineLength:    60,
		Centered:      true,
	},
	looseLyrics: lineStyle{
		LeadingBefore: 1,
		LineLength:    60,
		Italics:       true,
	},
	transition: lineStyle{
		LeadingBefore: 1,
		LineLength:    15,
		Indent:        45 * en,
	},

	character: lineStyle{
		LeadingBefore: 1,
		Indent:        22 * en,
		LineLength:    38,
	},
	parenthetical: lineStyle{
		Indent:          17 * en,
		FirstLineOffset: -1,
		LineLength:      24,
	},
	dialogue: lineStyle{
		Indent:     10 * en,
		LineLength: 40,
	},
	lyrics: lineStyle{
		Indent:     10 * en,
		LineLength: 40,
		Italics:    true,
	},

	dualCharacterOne: lineStyle{
		LeadingBefore: 1,
		Indent:        12 * en,
		LineLength:    17,
	},
	dualParentheticalOne: lineStyle{
		Indent:          3 * en,
		FirstLineOffset: -1,
		LineLength:      17,
	},
	dualDialogueOne: lineStyle{
		LineLength: 29,
	},
	dualLyricsOne: lineStyle{
		LineLength: 29,
		Italics:    true,
	},

	dualCharacterTwo: lineStyle{
		LeadingBefore: 1,
		Indent:        43 * en,
		LineLength:    17,
	},
	dualParentheticalTwo: lineStyle{
		Indent:          34 * en,
		FirstLineOffset: -1,
		LineLength:      17,
	},
	dualDialogueTwo: lineStyle{
		Indent:     31 * en,
		LineLength: 29,
	},
	dualLyricsTwo: lineStyle{
		Indent:     31 * en,
		LineLength: 29,
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
		LineLength:    30,
	},
	titlePageLeft: lineStyle{
		LeadingBefore: 2,
		LineLength:    30,
	},
}
