package dir

import (
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig"

	"github.com/kha7iq/subvars/cmd/helpers"
	"github.com/urfave/cli/v2"
)

type Directory struct {
	Path    string
	OutDir  string
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
				Name:        "input",
				Aliases:     []string{"i"},
				Usage:       "Path of folder containing template file.s",
				EnvVars:     []string{"SUBVARS_INPUTDIR"},
			},
			&cli.StringFlag{
				Destination: &subVarsOpts.OutDir,
				Name:        "out",
				Aliases:     []string{"o"},
				Usage:       "Output folder path. If folder does not exist it will be created.",
				EnvVars:     []string{"SUBVARS_OUTDIR"},
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

				if len(helpers.GlobalOpts.Prefix) != 0 {
					helpers.EnvVariables = helpers.MatchPrefix(helpers.GlobalOpts.Prefix)
				} else {
					helpers.EnvVariables = helpers.GetVars()
				}

				t = t.Option("missingkey=" + helpers.GlobalOpts.MissingKey)
				if len(subVarsOpts.OutDir) != 0 {
					if err := helpers.CreateDirIfNotExist(subVarsOpts.OutDir); err != nil {
						return err
					}
					_, outfile := path.Split(v)
					file, err := os.Create(subVarsOpts.OutDir + "/" + outfile)
					if err != nil {
						return err
					}
					if err := t.Execute(file, helpers.EnvVariables); err != nil {
						return err
					}
					file.Close()
					return nil
				}
				if err := t.Execute(os.Stdout, helpers.EnvVariables); err != nil {
					return err
				}
			}

			return nil
		},
	}
}
