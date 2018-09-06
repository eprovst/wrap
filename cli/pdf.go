package cli

import (
	"errors"

	"github.com/Wraparound/wrap/pdf"
	"github.com/spf13/cobra"
)

// pdfCmd represents the pdf command
var pdfCmd = &cobra.Command{
	Use:              "pdf [path to input file]",
	Short:            "Export file as PDF",
	Args:             cobra.MaximumNArgs(1),
	TraverseChildren: true,
	Long:             longDescription,
	Run:              pdfRun,
}

var (
	pdfNoSceneNumbersFlag bool
	useCourierPrime       bool
	useCourierNew         bool
	useFreeMono           bool
)

func init() {
	pdfCmd.Flags().BoolVarP(&pdfNoSceneNumbersFlag, "no-scene-numbers", "s", false, "remove scene numbers from output")
	pdfCmd.Flags().BoolVar(&useCourierPrime, "use-courier-prime", false, "force the usage of Courier Prime")
	pdfCmd.Flags().BoolVar(&useCourierNew, "use-courier-new", false, "force the usage of Courier New")
	pdfCmd.Flags().BoolVar(&useFreeMono, "use-freemono", false, "force the usage of GNU FreeMono")

	WrapCmd.AddCommand(pdfCmd)
}

func pdfRun(cmd *cobra.Command, args []string) {
	// Evaluate font selection
	pdf.AutoFontSelection = false

	if useCourierPrime && useCourierNew || useCourierPrime && useFreeMono || useCourierNew && useFreeMono {
		// The fonts are mutualy exclusive so throw an error
		handle(errors.New("tried to force multiple fonts at the same time"))
	}

	if useCourierPrime {
		pdf.SelectedFont = pdf.CourierPrime

	} else if useCourierNew {
		pdf.SelectedFont = pdf.CourierNew

	} else if useFreeMono {
		pdf.SelectedFont = pdf.FreeMono

	} else {
		// Else use automatic font selection
		pdf.AutoFontSelection = true
	}

	if pdfNoSceneNumbersFlag {
		pdf.AddSceneNumbers = false
	}

	export(args, "pdf", pdf.WritePDF)
}
