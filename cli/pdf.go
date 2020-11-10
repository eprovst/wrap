package cli

import (
	"errors"
	"strings"

	"github.com/Wraparound/wrap/pdf"
	"github.com/spf13/cobra"
)

// pdfCmd represents the pdf command
var pdfCmd = &cobra.Command{
	Use:              "pdf [path to input file]",
	Example:          "  wrap pdf MyScript.fountain\n  cat OtherScript.wrap | wrap pdf",
	Short:            "Export file as PDF",
	Args:             cobra.MaximumNArgs(1),
	TraverseChildren: true,
	Long:             longDescription,
	Run:              pdfRun,
}

var (
	pdfProductionFlag bool
	useCourierPrime   bool
	useCourierNew     bool
	useFreeMono       bool
	pageSize          string
	font              string
)

func init() {
	pdfCmd.Flags().BoolVarP(&pdfProductionFlag, "production", "p", false, "add scene numbers and other production text")
	pdfCmd.Flags().BoolVar(&useCourierPrime, "use-courier-prime", false, "force the usage of Courier Prime")
	pdfCmd.Flags().BoolVar(&useCourierNew, "use-courier-new", false, "force the usage of Courier New")
	pdfCmd.Flags().BoolVar(&useFreeMono, "use-freemono", false, "force the usage of GNU FreeMono")
	pdfCmd.Flags().StringVar(&pageSize, "page-size", "", "choose page size (letter or a4)")
	pdfCmd.RegisterFlagCompletionFunc("page-size", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"letter", "a4"}, cobra.ShellCompDirectiveDefault
	})
	pdfCmd.Flags().StringVar(&font, "font", "", "provide font as \"regular.ext, bold.ext, italic.ext, bolditalic.ext\"")

	WrapCmd.AddCommand(pdfCmd)
}

func pdfRun(cmd *cobra.Command, args []string) {
	// Evaluate font selection
	pdf.AutoFontSelection = false

	if (font != "" && atLeastOne(useCourierPrime, useCourierNew, useFreeMono)) ||
		moreThanOne(useCourierPrime, useCourierNew, useFreeMono) {
		// The fonts are mutualy exclusive so throw an error
		handle(errors.New("tried to force multiple fonts at the same time"))
	}

	if font != "" {
		fontfiles := strings.Split(font, ",")

		if len(fontfiles) != 4 {
			handle(errors.New("need four files for font"))
		}

		for i := range fontfiles {
			fontfiles[i] = strings.TrimSpace(fontfiles[i])
		}

		pdf.SelectedFont = pdf.Font{
			RomanName:      fontfiles[0],
			Roman:          []string{fontfiles[0]},
			BoldName:       fontfiles[1],
			Bold:           []string{fontfiles[1]},
			ItalicName:     fontfiles[2],
			Italic:         []string{fontfiles[2]},
			BoldItalicName: fontfiles[3],
			BoldItalic:     []string{fontfiles[3]},
		}

	} else if useCourierPrime {
		pdf.SelectedFont = pdf.CourierPrime

	} else if useCourierNew {
		pdf.SelectedFont = pdf.CourierNew

	} else if useFreeMono {
		pdf.SelectedFont = pdf.FreeMono

	} else {
		// Else use automatic font selection
		pdf.AutoFontSelection = true
	}

	pdf.Production = pdfProductionFlag

	if pageSize != "" {
		pdf.PageSize = pageSize
	}

	export(args, "pdf", pdf.WritePDF)
}
