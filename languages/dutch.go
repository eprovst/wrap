package languages

// DutchTranslation is the Dutch translation of Wrap
var DutchTranslation = Translation{
	Language: Dutch,

	SceneTags: []string{
		"int", "ext", "est", "int./ext", "int/ext", "i/e",
		"bin", "bui", "bin./bui", "open", "bin/bui", "bi/bu",
	},

	StageplaySceneTags: []string{
		"scene", "scène",
	},

	TransitionTags: []string{
		"TO:", "NAAR:",
	},

	BeginActTags: []string{
		"BEDRIJF",
	},

	EndActTags: []string{
		"EINDE BEDRIJF", "EINDE VAN BEDRIJF",
	},

	More: "(MEER)",

	Contd: "(VERDER)",
}
