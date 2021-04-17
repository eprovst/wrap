package pdf

// This file contains the definitions for a stageplay.

var stageplay = aTheme{
	action: lineStyle{
		Indent:        12.5 * en,
		LineLength:    33,
		LeadingBefore: 1,
	},
	act: lineStyle{
		LeadingBefore: 0,
		LineLength:    35,
		Indent:        25 * en,
		AllCaps:       true,
		Underline:     true,
	},
	slugLine: lineStyle{
		Indent:        25 * en,
		LineLength:    35,
		LeadingBefore: 1,
		Underline:     true,
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
		Indent:        25 * en,
		LineLength:    35,
		LeadingBefore: 1,
		AllCaps:       true,
	},

	character: lineStyle{
		Indent:        25 * en,
		LineLength:    35,
		LeadingBefore: 1,
	},
	parenthetical: lineStyle{
		Indent:     12.5 * en,
		LineLength: 33,
	},
	dialogue: lineStyle{
		LineLength: 60,
	},
	lyrics: lineStyle{
		LineLength: 60,
		Italics:    true,
	},

	dualCharacterOne: lineStyle{
		LeadingBefore: 1,
		Indent:        8 * en,
		LineLength:    20,
	},
	dualParentheticalOne: lineStyle{
		Indent:     3.5 * en,
		LineLength: 18,
	},
	dualDialogueOne: lineStyle{
		LineLength: 28,
	},
	dualLyricsOne: lineStyle{
		LineLength: 28,
		Italics:    true,
	},

	dualCharacterTwo: lineStyle{
		LeadingBefore: 1,
		Indent:        40 * en,
		LineLength:    20,
	},
	dualParentheticalTwo: lineStyle{
		Indent:     35.5 * en,
		LineLength: 18,
	},
	dualDialogueTwo: lineStyle{
		Indent:     32 * en,
		LineLength: 28,
	},
	dualLyricsTwo: lineStyle{
		Indent:     32 * en,
		LineLength: 28,
		Italics:    true,
	},

	more: lineStyle{
		Indent: 25 * en,
	},

	titlePageTitle: lineStyle{
		LeadingBefore: 15,
		Indent:        25 * en,
		LineLength:    35,
		Underline:     true,
		AllCaps:       true,
	},
	titlePageSubtitle: lineStyle{
		LeadingBefore: 1,
		Indent:        25 * en,
		LineLength:    35,
	},
	titlePageCredit: lineStyle{
		LeadingBefore: 1,
		Indent:        25 * en,
		LineLength:    35,
	},
	titlePageAuthor: lineStyle{
		LeadingBefore: 1,
		Indent:        25 * en,
		LineLength:    35,
	},
	titlePageSource: lineStyle{
		LeadingBefore: 1,
		Indent:        25 * en,
		LineLength:    35,
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
