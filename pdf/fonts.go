package pdf

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	findfont "github.com/flopp/go-findfont"
	"github.com/signintech/gopdf"
)

func findFont(font string) (string, error) {
	path, err := findfont.Find(font)

	if err != nil || filepath.Base(path) != font {
		// On some systems spaces get replaced by underscores
		underFont := strings.Replace(font, " ", "_", -1)
		path, err = findfont.Find(underFont)

		if err != nil || filepath.Base(path) != underFont {
			return "", errors.New("font not found")
		}

		return path, nil
	}

	return path, nil
}

func loadFonts() {
	// Courier Prime
	err := loadCourierPrime()

	if err != nil {
		if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
			// Courier New should be available on macOS and Windows
			fmt.Println("Warning: " + err.Error())
			err = loadCourierNew()
		}

		if err != nil {
			fmt.Println("Error: " + err.Error())
			os.Exit(1)
		}
	}

	// Now prepare the font
	setDefaultFont()
}

func loadCourierPrime() error {
	// Regular
	pathToRegular, err := findFont("Courier Prime.ttf")
	if err != nil {
		pathToRegular, err = findFont("Courier Prime Regular.ttf")
	}

	if err != nil {
		return errors.New("no Courier Prime installed")
	}

	// Bold
	pathToBold, err := findFont("Courier Prime Bold.ttf")

	if err != nil {
		return errors.New("no Courier Prime Bold installed")
	}

	// Italic
	pathToItalic, err := findFont("Courier Prime Italic.ttf")

	if err != nil {
		return errors.New("no Courier Prime Italic installed")
	}

	// Bold italic
	pathToBoldItalic, err := findFont("Courier Prime Bold Italic.ttf")

	if err != nil {
		return errors.New("no Courier Prime Bold Italic installed")
	}

	// Successfully found Courier Prime
	thisPDF.AddTTFFont("courier", pathToRegular)

	thisPDF.AddTTFFontWithOption("courier", pathToBold,
		gopdf.TtfOption{
			Style: gopdf.Bold,
		})

	thisPDF.AddTTFFontWithOption("courier", pathToItalic,
		gopdf.TtfOption{
			Style: gopdf.Italic,
		})

	thisPDF.AddTTFFontWithOption("courier", pathToBoldItalic,
		gopdf.TtfOption{
			Style: gopdf.Italic | gopdf.Bold,
		})

	return nil
}

func loadCourierNew() error {
	// A font installed by default on both Windows and macOS

	// Regular
	pathToRegular, err := findFont("Courier New.ttf")
	if err != nil {
		pathToRegular, err = findFont("cour.ttf")
	}

	if err != nil {
		return errors.New("no Courier New available")
	}

	// Bold
	pathToBold, err := findFont("Courier New Bold.ttf")
	if err != nil {
		pathToBold, err = findFont("courbd.ttf")
	}

	if err != nil {
		return errors.New("no Courier New Bold available")
	}

	// Italic
	pathToItalic, err := findFont("Courier New Italic.ttf")
	if err != nil {
		pathToRegular, err = findFont("couri.ttf")
	}

	if err != nil {
		return errors.New("no Courier New Italic available")
	}

	// Bold italic
	pathToBoldItalic, err := findFont("Courier New Bold Italic.ttf")
	if err != nil {
		pathToRegular, err = findFont("courbi.ttf")
	}

	if err != nil {
		return errors.New("no Courier New Bold Italic available")
	}

	// Successfully found Courier New
	thisPDF.AddTTFFont("courier", pathToRegular)

	thisPDF.AddTTFFontWithOption("courier", pathToBold,
		gopdf.TtfOption{
			Style: gopdf.Bold,
		})

	thisPDF.AddTTFFontWithOption("courier", pathToItalic,
		gopdf.TtfOption{
			Style: gopdf.Italic,
		})

	thisPDF.AddTTFFontWithOption("courier", pathToBoldItalic,
		gopdf.TtfOption{
			Style: gopdf.Italic | gopdf.Bold,
		})

	return nil
}

func setDefaultFont() {
	thisPDF.SetFont("courier", "", fontSize)
}

func setStyledFont(style string) {
	thisPDF.SetFont("courier", style, fontSize)
}
