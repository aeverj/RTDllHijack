package main

import (
	"AGHD/models"
	"AGHD/runner"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	var params models.Parameters
	app := &cli.App{
		Name:  "PE File Parser",
		Usage: "Parses PE file import tables and generates hijackable DLL source code",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "compiler",
				Aliases:     []string{"c"},
				Usage:       "Compiler to use (mingw or msvc)",
				Destination: &params.Compiler,
				Value:       "msvc",
			},
			&cli.StringFlag{
				Name:        "input",
				Aliases:     []string{"i"},
				Usage:       "Input file or directory path",
				Destination: &params.InputPath,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "Output directory path",
				Destination: &params.OutputPath,
				Value:       "",
			},
			&cli.StringFlag{
				Name:        "exclude",
				Aliases:     []string{"e"},
				Usage:       "Exclude file or directory name pattern",
				Destination: &params.ExcludePattern,
			},
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"v"},
				Usage:       "Enable verbose output",
				Destination: &params.Verbose,
				Value:       false,
			},
		},
		Action: func(c *cli.Context) error {
			r, err := runner.New(&params)
			if err != nil {
				return err
			}
			r.Run()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error:\n", err)
	}
}
