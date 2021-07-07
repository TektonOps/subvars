package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/kha7iq/subvars/cmd/dir"
	"github.com/kha7iq/subvars/cmd/helpers"

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
			Destination: &helpers.GlobalOpts.Prefix,
			EnvVars:     []string{"SUBVARS_PREFIX"},
		},
		&cli.StringFlag{
			Destination: &helpers.GlobalOpts.MissingKey,
			Name:        "missingkey",
			Aliases:     []string{"m"},
			Usage: "Behavior for missing key when parsing variables." +
				"Available options 'invalid', 'error' or 'zero'",
			EnvVars: []string{"SUBVARS_MISSINGKEY"},
			Value:   "invalid",
		},
	}

	app.Action = func(c *cli.Context) error {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("unable to read input \nERROR: %v\n", err)
		}

		t, err := helpers.ParseString(string(b))
		if err != nil {
			return fmt.Errorf("unable to parse string stream \nERROR: %v\n", err)
		}

		if len(helpers.GlobalOpts.Prefix) != 0 {
			helpers.EnvVariables = helpers.MatchPrefix(helpers.GlobalOpts.Prefix)
		} else {
			helpers.EnvVariables = helpers.GetVars()
		}
		t = t.Option("missingkey=" + helpers.GlobalOpts.MissingKey)
		if err := t.Execute(os.Stdout, helpers.EnvVariables); err != nil {
			return err
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
