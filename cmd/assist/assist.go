package assist

// assist package serve as helper functions and aid to app.

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

// Flags struct holds inputs values for global flags
type Flags struct {
	Prefix     string
	MissingKey string
}

var (
	// GlobalFlags variable holds user input flag values
	GlobalFlags Flags
	// EnvVariables is used for saving variables form environment.
	EnvVariables map[string]string
)

// GetVars will get all the environment variables
func GetVars() (enVars map[string]string) {
	enVars = make(map[string]string)
	for _, value := range os.Environ() {
		kv := strings.SplitN(value, "=", 2)
		enVars[kv[0]] = kv[1]
	}
	return
}

// ParseString will parse any input provided as string
func ParseString(str string) (*template.Template, error) {
	funcMap := sprig.TxtFuncMap()
	return template.Must(template.New("").Funcs(funcMap).Funcs(matchFunc()).Parse(str)), nil
}

// ParseFile will parse any input provided as string
func ParseFile(file string) (*template.Template, error) {
	funcMap := sprig.TxtFuncMap()
	return template.Must(template.New(filepath.Base(file)).Funcs(funcMap).Funcs(matchFunc()).ParseFiles(file)), nil
}

// MatchPrefix will match a given prefix pattern of all env variables and render only those.
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

// matchFunc returns a custom functions map
func matchFunc() template.FuncMap {
	functionMap := map[string]interface{}{
		"match": MatchPrefix,
	}
	return functionMap
}

// IsFlagSet function check if flag is set and returns a boolean true or false
//nolint
func IsFlagSet(fl string) bool {
	if len(fl) != 0 {
		return true
	}
	return false
}
