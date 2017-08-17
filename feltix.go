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

func main() {
	app := cli.NewApp()
	app.Name = "Feltix"
	app.Version = "v0.0.1-indev"
	app.Usage = "Generate HTML (and somewhere in the future PDF) output from Fountain"
	app.Author = "Evert Provoost"
	app.HideHelp = true
	app.ArgsUsage = "path/to/file"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "out, o",
			Usage: "specify the `file` name to be used",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "html",
			Usage: "export file as an HTML webpage",
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
				if c.GlobalString("out") != "" {
					pathToFile = c.GlobalString("out")
				} else {
					pathToFile = strings.TrimSuffix(pathToFile, extension) + ".html"
				}

				html := feltixhtml.ToHTML(script)
				err = ioutil.WriteFile(pathToFile, []byte(html), 0664)

				if err != nil {
					println(err.Error())
					return err
				}

				return nil
			},
		},
		/*{
			Name:    "pdf",
			Usage:   "export file as PDF",
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
				if c.GlobalString("out") != "" {
					pathToFile = c.GlobalString("out")
				} else {
					pathToFile = strings.TrimSuffix(pathToFile, extension) + ".pdf"
				}

				pdf := feltixpdf.ToPDF(script)
				err = ioutil.WriteFile(pathToFile, []byte(pdf), 0664)

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
