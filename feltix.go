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
	app.Version = "v0.0.1-indev"
	app.Usage = "Generate HTML (and somewhere in the future PDF) output from Fountain"
	app.Author = "Evert Provoost"
	app.HideHelp = true
	app.Commands = []cli.Command{
		{
			Name:     "html",
			Usage:    "Export file as an HTML webpage",
			Category: "Export formats",
			Action: func(c *cli.Context) error {
				pathToFile := c.Args().First()

				if pathToFile == "" {
					cli.ShowCommandHelpAndExit(c, "html", 0)
				}

				extension := strings.ToLower(filepath.Ext(pathToFile))

				// Check if we're using the feltix extensions
				if extension == ".ftx" || extension == ".feltix" {
					feltixparser.UseFeltixExtensions = true
				}

				script, err := feltixparser.ParseFile(pathToFile)

				if err != nil {
					println(err.Error())
					return err
				}

				// Get the filepath to use during export.
				pathToFile = strings.TrimSuffix(pathToFile, extension) + ".html"

				html := feltixhtml.ToHTML(script)
				err = writeOutput(html, pathToFile)

				if err != nil {
					println(err.Error())
					return err
				}

				return nil
			},
		},
		/*{
			Name:    "pdf",
			Usage:   "Export file as PDF",
			Category: "Export formats:",
			Action: func(c *cli.Context) error {
				pathToFile := c.Args().First()

				if pathToFile == "" {
					cli.ShowCommandHelpAndExit(c, "pdf", 0)
				}

				extension := strings.ToLower(filepath.Ext(pathToFile))

				// Check if we're using the feltix extensions
				if extension == ".ftx" || extension == ".feltix" {
					feltixparser.UseFeltixExtensions = true
				}

				script, err := feltixparser.ParseFile(pathToFile)

				if err != nil {
					println(err.Error())
					return err
				}

				// Get the filepath to use during export.
				pathToFile = strings.TrimSuffix(pathToFile, extension) + "pdf"

				html := feltixpdf.ToPDF(script)
				err = writeOutput(html, pathToFile)

				if err != nil {
					println(err.Error())
					return err
				}

				return nil
			},
		},*/
	}

	app.Run(os.Args)
}
