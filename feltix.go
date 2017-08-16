package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Feltix/feltixhtml"
	"github.com/Feltix/feltixparser"
	"github.com/urfave/cli"
)

func writeOutput(content string, path string) error {
	return ioutil.WriteFile(path, []byte(content), 0664)
}

func main() {
	app := cli.NewApp()
	app.Name = "Feltix"
	app.Version = "v0.0.1"
	app.Usage = "Generate HTML (and somewhere in the future PDF) output from Fountain"
	app.Author = "Evert Provoost"
	app.Commands = []cli.Command{
		{
			Name:     "html",
			Usage:    "Export file as an HTML webpage",
			Category: "Export formats",
			Action: func(c *cli.Context) error {
				pathToFile := c.Args().First()
				extension := filepath.Ext(pathToFile)

				// Check if we're using the feltix extensions
				if extension == ".ftx" {
					feltixparser.UseFeltixExtensions = true
				}

				script, err := feltixparser.ParseFile(pathToFile)

				// Get the filepath to use during export.
				pathToFile = strings.TrimSuffix(pathToFile, extension) + ".html"

				html := feltixhtml.ToHTML(script)
				writeOutput(html, pathToFile)
				return err
			},
		},
		/*{
			Name:    "pdf",
			Usage:   "Export file as PDF",
			Category: "Export formats:",
			Action: func(c *cli.Context) error {
				pathToFile := c.Args().First()
				extension := filepath.Ext(pathToFile)

				// Check if we're using the feltix extensions
				if extension == "ftx" {
					feltixparser.UseFeltixExtensions = true
				}

				script, err := feltixparser.ParseFile(pathToFile)

				// Get the filepath to use during export.
				pathToFile = strings.TrimSuffix(pathToFile, extension) + "pdf"

				html := feltixpdf.ToPDF(script)
				writeOutput(html, pathToFile)
				return err
			},
		},*/
	}

	app.Run(os.Args)
}
