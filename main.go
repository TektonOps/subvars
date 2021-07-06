package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/kha7iq/subvars/cmd/file"
	"github.com/kha7iq/subvars/cmd/helpers"

	"github.com/urfave/cli/v2"
)

var version string

func main() {
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		file.Render(),
	}
	app.Name = "subvars"
	app.Version = version
	app.Usage = "Substitute environment variables defined as go templates."
	app.Flags = append([]cli.Flag{}, helpers.GlobalFlags...)
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
