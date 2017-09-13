package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

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
		cli.BoolFlag{
			Name:  "benchmark",
			Usage: "measure the time spend on certain tasks",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "html",
			Usage: "export file as an HTML webpage",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "embedable, e",
					Usage: "only output the play itself",
				},
			},
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

				file, err := os.Open(pathToFile)
				defer file.Close() // Make sure it's closed at one point.

				if err != nil {
					println(err.Error())
					return err
				}

				t0 := time.Now()
				script, err := feltixparser.Parser(file)

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

				t1 := time.Now()
				var html string
				if c.Bool("embedable") {
					html = feltixhtml.ToHTML(script)
				} else {
					html = feltixhtml.ToHTMLPage(script)
				}
				t2 := time.Now()

				err = ioutil.WriteFile(pathToFile, []byte(html), 0664)

				if err != nil {
					println(err.Error())
					return err
				}

				if c.GlobalBool("benchmark") {
					print("Parsing:   ")
					print(t1.Sub(t0) / time.Millisecond)
					println(" ms")
					print("Exporting: ")
					print(t2.Sub(t1) / time.Millisecond)
					println(" ms")
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
