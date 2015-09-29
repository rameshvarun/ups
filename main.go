package main

import (
	"io/ioutil"
	"os"

	"github.com/codegangsta/cli"
	"github.com/rameshvarun/ups/operations"
	"github.com/rameshvarun/ups/reader"
)

func main() {
	app := cli.NewApp()
	app.Name = "ups"
	app.Usage = "Utilities for manipulating / creating UPS patch files."

	app.Commands = []cli.Command{
		{
			Name:    "apply",
			Usage:   "Apply a .ups patch to a file.",
			Aliases: []string{"patch"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "base, b",
					Usage: "The base file on top of which to apply the patch.",
				},
				cli.StringFlag{
					Name:  "patch, p",
					Usage: "The patch file to apply.",
				},
				cli.StringFlag{
					Name:  "output, o",
					Usage: "The file in which to write the patched data.",
				},
			},
			Action: func(c *cli.Context) {
				base, err := ioutil.ReadFile(c.String("base"))
				if err != nil {
					panic(err)
				}

				patchBytes, err := ioutil.ReadFile(c.String("patch"))
				if err != nil {
					panic(err)
				}
				patch, err := reader.ReadUPS(patchBytes)
				if err != nil {
					panic(err)
				}

				result, err := operations.Apply(base, patch, false)
				if err != nil {
					panic(err)
				}

				ioutil.WriteFile(c.String("output"), result, 0644)
			},
		},
		{
			Name:    "diff",
			Usage:   "Diff two files, creating a .ups patch.",
			Aliases: []string{"create"},
			Action: func(c *cli.Context) {
				panic("Not yet implemented.")
			},
		},
		{
			Name:  "revert",
			Usage: "Revert a patched file to it's original state.",
			Action: func(c *cli.Context) {
				panic("Not yet implemented.")
			},
		},
		{
			Name:    "merge",
			Usage:   "Merge two .ups files, creating a file equivalent to applying both files.",
			Aliases: []string{"combine"},
			Action: func(c *cli.Context) {
				panic("Not yet implemented.")
			},
		},
	}

	app.Run(os.Args)
}
