package file

import (
	"os"

	"github.com/kha7iq/subvars/cmd/helpers"
	"github.com/urfave/cli/v2"
)

type File struct {
	Path string
	Glob string
}

func Render() *cli.Command {
	var subVarsOpts File
	return &cli.Command{
		Name:    "file",
		Aliases: []string{"f"},
		Usage:   "input from a file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &subVarsOpts.Path,
				Name:        "name",
				Aliases:     []string{"n"},
				Usage:       "SubVars file path",
				EnvVars:     []string{"SUBVARS_FILE_NAME"},
			},
			&cli.StringFlag{
				Destination: &subVarsOpts.Glob,
				Name:        "glob",
				Aliases:     []string{"g"},
				Usage:       "Glob Matching",
				EnvVars:     []string{"SUBVARS_GLOB"},
			},
		},
		Action: func(ctx *cli.Context) error {
			t, err := helpers.ParseFiles(subVarsOpts.Path)
			if err != nil {
				return err
			}
			if len(helpers.GlobalOpts.Prefix) != 0 {
				vars := helpers.MatchPrefix(helpers.GlobalOpts.Prefix)
				t = t.Option("missingkey=" + helpers.GlobalOpts.MissingKey)
				if err := t.Execute(os.Stdout, vars); err != nil {
					return err
				}

				return nil
			}
			vars := helpers.GetVars()
			t = t.Option("missingkey=" + helpers.GlobalOpts.MissingKey)
			if err := t.Execute(os.Stdout, vars); err != nil {
				return err
			}
			return nil
		},
	}
}
