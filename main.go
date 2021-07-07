package main

import (
	"errors"
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
			Name:        "prefix, pr",
			Usage:       "Match only variables with given prefix pattern",
			Destination: &helpers.GlobalOpts.Prefix,
			EnvVars:     []string{"SUBVARS_PREFIX"},
		},
		&cli.StringFlag{
			Destination: &helpers.GlobalOpts.MissingKey,
			Name:        "missingkey",
			Usage:       "Behavior for missing key when parsing variables." +
				"Available options 'invalid', 'error' or 'zero'",
			EnvVars:     []string{"SUBVARS_MISSINGKEY"},
			Value:       "invalid",
		},
	}

	app.Action = func(c *cli.Context) error {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return errors.New("unable to read input")
		}

		t, err := helpers.ParseString(string(b))
		if err != nil {
			return errors.New("error parsing string stream")
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
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
