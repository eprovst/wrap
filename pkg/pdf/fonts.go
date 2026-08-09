package pdf

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	findfont "github.com/flopp/go-findfont"
	"github.com/signintech/gopdf"
)

// Font are fonts that the PDF export can use
type Font struct {
	RomanName      string
	Roman          []string
	BoldName       string
	Bold           []string
	ItalicName     string
	Italic         []string
	BoldItalicName string
	BoldItalic     []string
}

// Selectable fonts
var (
	CourierPrime = Font{
		RomanName:      "Courier Prime",
		Roman:          []string{"Courier Prime.ttf", "Courier Prime Regular.ttf", "CourierPrime-Regular.ttf", "courier-prime.ttf"},
		BoldName:       "Courier Prime Bold",
		Bold:           []string{"Courier Prime Bold.ttf", "CourierPrime-Bold.ttf", "courier-prime-bold.ttf"},
		ItalicName:     "Courier Prime Italic",
		Italic:         []string{"Courier Prime Italic.ttf", "CourierPrime-Italic.ttf", "courier-prime-italic.ttf"},
		BoldItalicName: "Courier Prime Bold Italic",
		BoldItalic:     []string{"Courier Prime Bold Italic.ttf", "CourierPrime-BoldItalic.ttf", "courier-prime-bold-italic.ttf"},
	}

	CourierNew = Font{
		RomanName:      "Courier New",
		Roman:          []string{"Courier New.ttf", "cour.ttf"},
		BoldName:       "Courier New Bold",
		Bold:           []string{"Courier New Bold.ttf", "courbd.ttf"},
		ItalicName:     "Courier New Italic",
		Italic:         []string{"Courier New Italic.ttf", "couri.ttf"},
		BoldItalicName: "Courier New Bold Italic",
		BoldItalic:     []string{"Courier New Bold Italic.ttf", "courbi.ttf"},
	}

	FreeMono = Font{
		RomanName:      "FreeMono",
		Roman:          []string{"FreeMono.ttf"},
		BoldName:       "FreeMono Bold",
		Bold:           []string{"FreeMonoBold.ttf"},
		ItalicName:     "FreeMono Oblique",
		Italic:         []string{"FreeMonoOblique.ttf"},
		BoldItalicName: "FreeMono Bold Oblique",
		BoldItalic:     []string{"FreeMonoBoldOblique.ttf"},
	}
)

// AutoFontSelection enables automatic font selection
var AutoFontSelection = true

// SelectedFont is the font to be used during export if AutoSelect is disabled
var SelectedFont = CourierPrime

var LoadedFont Font

func findFont(fonts []string) (string, error) {
	for _, font := range fonts {
		path, err := findfont.Find(font)

		if err != nil || filepath.Base(path) != font {
			// On some systems spaces get replaced by underscores
			underFont := strings.Replace(font, " ", "_", -1)
			path, err = findfont.Find(underFont)

			if err != nil || filepath.Base(path) != underFont {
				continue
			}
		}

		return path, nil
	}

	// Not found
	return "", errors.New("font not found")
}

func loadFont(font Font) error {
	// Roman
	pathToRegular, err := findFont(font.Roman)

	if err != nil {
		return errors.New("no " + font.RomanName + " installed")
	}

	// Bold
	pathToBold, err := findFont(font.Bold)

	if err != nil {
		return errors.New("no " + font.BoldName + " installed")
	}

	// Italic
	pathToItalic, err := findFont(font.Italic)

	if err != nil {
		return errors.New("no " + font.ItalicName + " installed")
	}

	// Bold italic
	pathToBoldItalic, err := findFont(font.BoldItalic)

	if err != nil {
		return errors.New("no " + font.BoldItalicName + " installed")
	}

	// Successfully found the font
	LoadedFont = font
	thisPDF.AddTTFFont(LoadedFont.RomanName, pathToRegular)

	thisPDF.AddTTFFontWithOption(LoadedFont.RomanName, pathToBold,
		gopdf.TtfOption{
			Style: gopdf.Bold,
		})

	thisPDF.AddTTFFontWithOption(LoadedFont.RomanName, pathToItalic,
		gopdf.TtfOption{
			Style: gopdf.Italic,
		})

	thisPDF.AddTTFFontWithOption(LoadedFont.RomanName, pathToBoldItalic,
		gopdf.TtfOption{
			Style: gopdf.Italic | gopdf.Bold,
		})

	return nil
}

func loadFonts() {
	if !AutoFontSelection {
		err := loadFont(SelectedFont)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: "+err.Error())
			os.Exit(1)
		}

	} else {
		// Attempt auto selection
		// Courier Prime
		err := loadFont(CourierPrime)

		if err != nil {
			// Courier New should be available on macOS and Windows
			fmt.Fprintln(os.Stderr, "Warning: "+err.Error())
			err = loadFont(CourierNew)

			if err != nil {
				// FreeMono as a final attempt
				fmt.Fprintln(os.Stderr, "Warning: "+err.Error())
				err = loadFont(FreeMono)

				if err != nil {
					fmt.Fprintln(os.Stderr, "Error: "+err.Error())
					os.Exit(1)
				}
			}
		}
	}

	// Now prepare the font
	setDefaultFont()
}

func setDefaultFont() {
	thisPDF.SetFont(LoadedFont.RomanName, "", fontSize)
}

func setStyledFont(style string) {
	thisPDF.SetFont(LoadedFont.RomanName, style, fontSize)
}
