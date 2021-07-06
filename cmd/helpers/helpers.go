package helpers

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/urfave/cli/v2"
)

type Flags struct {
	Prefix     string
	MissingKey string
}

var (
	GlobalOpts Flags

	GlobalFlags = []cli.Flag{
		&cli.StringFlag{
			Name:        "prefix, pr",
			Usage:       "Match only variables with given prefix pattern",
			Destination: &GlobalOpts.Prefix,
			EnvVars:     []string{"SUBVARS_PREFIX_PATTERN"},
		},
		&cli.StringFlag{
			Destination: &GlobalOpts.MissingKey,
			Name:        "missingkey",
			Usage:       "Behavior for missing key when parsing variables",
			EnvVars:     []string{"SUBVARS_MISSINGKEY"},
			Value:       "default",
		},
	}
)

func GetVars() (enVars map[string]interface{}) {
	enVars = make(map[string]interface{})
	for _, value := range os.Environ() {
		kv := strings.SplitN(value, "=", 2)
		enVars[kv[0]] = kv[1]
	}
	return
}

func ParseString(str string) (*template.Template, error) {
	funcMap := sprig.TxtFuncMap()
	return template.Must(template.New("").Funcs(funcMap).Parse(str)), nil
}

func ParseFiles(files ...string) (*template.Template, error) {
	funcMap := sprig.TxtFuncMap()
	return template.Must(template.New(filepath.Base(files[0])).Funcs(funcMap).ParseFiles(files...)), nil
}

func MatchPrefix(prefix string) map[string]string {
	enVars := make(map[string]string)
	for _, value := range os.Environ() {
		kv := strings.SplitN(value, "=", 2)
		if strings.HasPrefix(kv[0], prefix) {
			enVars[kv[0]] = kv[1]
		}
	}
	return enVars
}
