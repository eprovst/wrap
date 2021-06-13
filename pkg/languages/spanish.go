package languages

// SpanishTranslation is the Spanish translation of Wrap

var SpanishTranslation = Translation{
	Language: Spanish,

	SceneTags: []string{
        "int ", "ext ", "est ", "int./ext ", "int/ext ", "i/e ", "i./e ",
		"int.", "ext.", "est.", "int./ext.", "int/ext.", "i/e.", "i./e.",
	},

	StageplaySceneTags: []string{
		"scene ", "escena ",
	},

	TransitionTags: []string{
		" TO:", " A:",
	},

	BeginActTags: []string{
		"ACTO ",
	},

	EndActTags: []string{
		"FIN DE ACTO ",
	},

	More: "(MÁS)",

	Contd: "(CONTINÚA)",
}
