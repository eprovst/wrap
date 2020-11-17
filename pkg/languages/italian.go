package languages

// ItalianTranslation is the Italian translation of Wrap

var ItalianTranslation = Translation{
	Language: Italian,

	SceneTags: []string{
		"int ", "ext ", "est ", "int./ext ", "int/ext ", "i/e ", "i./e ", "i./e ",
		"int.", "ext.", "est.", "int./ext.", "int/ext.", "i/e.", "i./e.", "i./e.",
		"d'amb ", "amb ", "int./est ", "int/est ",
		"d'amb.", "amb.", "int./est.", "int/est.",
	},

	StageplaySceneTags: []string{
		"scene ", "scena ",
	},

	TransitionTags: []string{
		" TO:", " A:",
	},

	BeginActTags: []string{
		"ATTO ",
	},

	EndActTags: []string{
		"FINE ATTO ",
	},

	More: "(SEGUE)",

	Contd: "(CONTINUA)",
}
