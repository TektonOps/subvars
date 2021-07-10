package dir

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/kha7iq/subvars/cmd/assist"
	"github.com/urfave/cli/v2"
)

type flagsDir struct {
	InputDir string
	OutDir   string
}

// Render function will render all the template files in any given input folder and saves
// them in target folder.
func Render() *cli.Command {
	var subVarsOpts flagsDir
	return &cli.Command{
		Name:    "dir",
		Aliases: []string{"d"},
		Usage:   "Directory lets you render all files in a folder & subfolder.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &subVarsOpts.InputDir,
				Name:        "input",
				Aliases:     []string{"i"},
				Usage:       "Path of folder containing template files.",
				EnvVars:     []string{"SUBVARS_INPUTDIR"},
			},
			&cli.StringFlag{
				Destination: &subVarsOpts.OutDir,
				Name:        "out",
				Aliases:     []string{"o"},
				Usage:       "Output folder path. If folder does not exist it will be created automatically.",
				EnvVars:     []string{"SUBVARS_OUTDIR"},
			},
		},
		Action: func(ctx *cli.Context) error {
			paths, err := getPathInDir(subVarsOpts.InputDir)
			if err != nil {
				return err
			}

			for _, v := range paths {
				t, err := assist.ParseFile(v)
				if err != nil {
					return fmt.Errorf("unable to parse file \nError: %v", err)
				}

				if assist.IsFlagSet(assist.GlobalFlags.Prefix) {
					assist.EnvVariables = assist.MatchPrefix(assist.GlobalFlags.Prefix)
				} else {
					assist.EnvVariables = assist.GetVars()
				}

				t = t.Option("missingkey=" + assist.GlobalFlags.MissingKey)
				if assist.IsFlagSet(subVarsOpts.OutDir) {
					if err := createDirIfNotExist(subVarsOpts.OutDir); err != nil {
						return err
					}
					_, outfile := path.Split(v)
					file, err := os.Create(subVarsOpts.OutDir + "/" + outfile)
					if err != nil {
						return err
					}
					if err := t.Execute(file, assist.EnvVariables); err != nil {
						return err
					}
					file.Close()
				} else {
					if err := t.Execute(os.Stdout, assist.EnvVariables); err != nil {
						return err
					}
				}
			}

			return nil
		},
	}
}

// getPathInDir recursively get all file paths in directory, including sub-directories.
func getPathInDir(pid string) ([]string, error) {
	var paths []string
	err := filepath.Walk(pid, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return paths, nil
}

// createDirIfNotExist will check if folder does not exist it will create it.
func createDirIfNotExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.Mkdir(path, os.ModePerm); err != nil {
			return err
		}
		return err
	}
	return nil
}
