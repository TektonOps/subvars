package dir

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig"

	"github.com/kha7iq/subvars/cmd/helpers"
	"github.com/urfave/cli/v2"
)

type Directory struct {
	Path string
}

func Render() *cli.Command {
	var subVarsOpts Directory
	return &cli.Command{
		Name:    "dir",
		Aliases: []string{"d"},
		Usage:   "Directory lets you render all files in a folder & sub folder.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &subVarsOpts.Path,
				Name:        "path",
				Aliases:     []string{"p"},
				Usage:       "Path of folder containing template file.s",
				EnvVars:     []string{"SUBVARS_DIR_PATH"},
			},
		},
		Action: func(ctx *cli.Context) error {
			paths, err := helpers.GetPathInDir(subVarsOpts.Path)
			if err != nil {
				return err
			}
			for _, v := range paths {
				funcMap := sprig.TxtFuncMap()
				t := template.Must(template.New(filepath.Base(v)).Funcs(funcMap).ParseFiles(v))
				vars := helpers.GetVars()
				t = t.Option("missingkey=" + helpers.GlobalOpts.MissingKey)

				if err := t.Execute(os.Stdout, vars); err != nil {
					return err
				}
			}
			return nil
		},
	}
}
