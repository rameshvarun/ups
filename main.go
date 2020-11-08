package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codegangsta/cli"
	"github.com/rameshvarun/ups/operations"
	"github.com/rameshvarun/ups/reader"
	"github.com/rameshvarun/ups/writer"
)

func main() {
	app := cli.NewApp()
	app.Name = "ups"
	app.Usage = "Utilities for manipulating / creating UPS patch files."

	app.Commands = []*cli.Command{
		{
			Name:    "apply",
			Usage:   "Apply a .ups patch to a file.",
			Aliases: []string{"patch"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "base, b",
					Usage: "The base file on top of which to apply the patch.",
				},
				&cli.StringFlag{
					Name:  "patch, p",
					Usage: "The patch file to apply.",
				},
				&cli.StringFlag{
					Name:  "output, o",
					Usage: "The file in which to write the patched data.",
				},
			},
			Action: func(c *cli.Context) error {
				if c.String("base") == "" || c.String("patch") == "" || c.String("output") == "" {
					if c.String("base") == "" {
						fmt.Printf("Missing required argument 'base'.\n")
					}
					if c.String("patch") == "" {
						fmt.Printf("Missing required argument 'patch'.\n")
					}
					if c.String("output") == "" {
						fmt.Printf("Missing required argument 'output'.\n")
					}
					fmt.Println()

					err := cli.ShowCommandHelp(c, "apply")
					if err != nil {
						panic(err)
					}
					return cli.Exit("", 1)
				}

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

				err = ioutil.WriteFile(c.String("output"), result, 0644)
				if err != nil {
					panic(err)
				}

				return nil
			},
		},
		{
			Name:    "diff",
			Usage:   "Diff two files, creating a UPS patch.",
			Aliases: []string{"create"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "base, b",
					Usage: "The base file.",
				},
				&cli.StringFlag{
					Name:  "modified, m",
					Usage: "The modified file.",
				},
				&cli.StringFlag{
					Name:  "output, o",
					Usage: "The file in which to write the patch data.",
				},
			},
			Action: func(c *cli.Context) error {
				if c.String("base") == "" || c.String("modified") == "" || c.String("output") == "" {
					if c.String("base") == "" {
						fmt.Printf("Missing required argument 'base'.\n")
					}
					if c.String("modified") == "" {
						fmt.Printf("Missing required argument 'patch'.\n")
					}
					if c.String("output") == "" {
						fmt.Printf("Missing required argument 'output'.\n")
					}
					fmt.Println()

					err := cli.ShowCommandHelp(c, "diff")
					if err != nil {
						panic(err)
					}
					return cli.Exit("", 1)
				}

				base, err := ioutil.ReadFile(c.String("base"))
				if err != nil {
					panic(err)
				}

				modified, err := ioutil.ReadFile(c.String("modified"))
				if err != nil {
					panic(err)
				}

				// Create patch data and write it to file.
				patch := operations.Diff(base, modified)
				err = ioutil.WriteFile(c.String("output"), writer.WriteUPS(patch), 0644)
				if err != nil {
					panic(err)
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
