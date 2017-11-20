package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Feltix/feltixpdf"

	"github.com/Feltix/feltixhtml"
	"github.com/Feltix/feltixparser"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Feltix"
	app.Version = "v0.1.1"
	app.Usage = "Generate HTML and/or PDF output from Fountain files"
	app.Author = "Evert Provoost"
	app.Email = "evert.provoost@gmail.com"
	app.HideHelp = true
	app.ArgsUsage = "path/to/file"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "benchmark",
			Usage: "measure the time spend on certain tasks",
		},
	}
	app.Action = func(c *cli.Context) error {
		return cli.ShowAppHelp(c)
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
				cli.StringFlag{
					Name:  "out, o",
					Usage: "specify the `file` name to be used",
				},
				cli.BoolFlag{
					Name:  "noscenenumbers, s",
					Usage: "remove scnenenumbers from output",
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

				t0 := time.Now()
				script, err := feltixparser.ParseFile(pathToFile)

				if err != nil {
					println(err.Error())
					return err
				}

				// Get the filepath to use during export.
				if c.String("out") != "" {
					pathToFile = c.String("out")
				} else {
					pathToFile = strings.TrimSuffix(pathToFile, extension) + ".html"
				}

				t1 := time.Now()
				var html string

				if c.Bool("noscenenumbers") {
					feltixhtml.AddSceneNumbers = false
				}

				if c.Bool("embedable") {
					html = feltixhtml.ToHTML(script)
					err = ioutil.WriteFile(pathToFile, []byte(html), 0664)
				} else {
					err = feltixhtml.WriteHTMLPage(script, pathToFile)
				}

				if err != nil {
					println(err.Error())
					return err
				}

				t2 := time.Now()
				if c.GlobalBool("benchmark") {
					print("Parsing:   ")
					print(t1.Sub(t0) / time.Millisecond)
					println(" ms")
					print("Exporting: ")
					print(t2.Sub(t1) / time.Millisecond)
					println(" ms")
					print("Total:     ")
					print(t2.Sub(t0) / time.Millisecond)
					println(" ms")
				}

				return nil
			},
		},
		{
			Name:  "pdf",
			Usage: "export file as PDF",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "out, o",
					Usage: "specify the `file` name to be used",
				},
				cli.BoolFlag{
					Name:  "noscenenumbers, s",
					Usage: "remove scnenenumbers from output",
				},
			},
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

				t0 := time.Now()
				script, err := feltixparser.ParseFile(pathToFile)

				if err != nil {
					println(err.Error())
					return err
				}

				// Get the filepath to use during export.
				if c.String("out") != "" {
					pathToFile = c.String("out")
				} else {
					pathToFile = strings.TrimSuffix(pathToFile, extension) + ".pdf"
				}

				if c.Bool("noscenenumbers") {
					feltixpdf.AddSceneNumbers = false
				}

				t1 := time.Now()
				err = feltixpdf.WritePDFFile(script, pathToFile)

				if err != nil {
					println(err.Error())
					return err
				}

				t2 := time.Now()
				if c.GlobalBool("benchmark") {
					print("Parsing:   ")
					print(t1.Sub(t0) / time.Millisecond)
					println(" ms")
					print("Exporting: ")
					print(t2.Sub(t1) / time.Millisecond)
					println(" ms")
					print("Total:     ")
					print(t2.Sub(t0) / time.Millisecond)
					println(" ms")
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}
