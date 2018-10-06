package languages

// FrenchTranslation is the French translation of Wrap
var FrenchTranslation = Translation{
	Language: French,

	SceneTags: []string{
		"int", "ext", "est", "int./ext", "int/ext", "i/e",
	},

	StageplaySceneTags: []string{
		"scene", "scène",
	},

	TransitionTags: []string{
		"TO:", "A:", "À:",
	},

	BeginActTags: []string{
		"ACTE",
	},

	EndActTags: []string{
		"FIN D'ACTE", "FIN D' ACTE",
	},

	More: "(PLUS)",

	Contd: "(suite)",

	Continued: "CONTINUÉ:",
}
