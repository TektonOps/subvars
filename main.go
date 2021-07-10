package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/kha7iq/subvars/cmd/assist"
	"github.com/kha7iq/subvars/cmd/dir"

	"github.com/urfave/cli/v2"
)

var version string

func main() {
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		dir.Render(),
	}
	app.Name = "subvars"
	app.Version = version
	app.Usage = "Substitute environment variables defined as go templates."
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "prefix",
			Aliases:     []string{"pr"},
			Usage:       "Match only variables with given prefix pattern",
			Destination: &assist.GlobalFlags.Prefix,
			EnvVars:     []string{"SUBVARS_PREFIX"},
		},
		&cli.StringFlag{
			Destination: &assist.GlobalFlags.MissingKey,
			Name:        "missingkey",
			Aliases:     []string{"m"},
			Usage: "Behavior for missing key when parsing variables." +
				" Available options 'invalid', 'error' or 'zero'",
			EnvVars: []string{"SUBVARS_MISSINGKEY"},
			Value:   "invalid",
		},
	}

	app.Action = func(c *cli.Context) error {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("unable to read input \nError: %v", err)
		}

		t, err := assist.ParseString(string(b))
		if err != nil {
			return fmt.Errorf("unable to parse string stream \nError: %v", err)
		}

		if assist.IsFlagSet(assist.GlobalFlags.Prefix) {
			assist.EnvVariables = assist.MatchPrefix(assist.GlobalFlags.Prefix)
		} else {
			assist.EnvVariables = assist.GetVars()
		}
		t = t.Option("missingkey=" + assist.GlobalFlags.MissingKey)
		if err := t.Execute(os.Stdout, assist.EnvVariables); err != nil {
			return err
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
