package helpers

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
)

// Flags struct holds inputs values for global flags
type Flags struct {
	Prefix     string
	MissingKey string
}

var (
	GlobalOpts   Flags
	EnvVariables map[string]interface{}
)

// GetVars will get all the environment variables
func GetVars() (enVars map[string]interface{}) {
	enVars = make(map[string]interface{})
	for _, value := range os.Environ() {
		kv := strings.SplitN(value, "=", 2)
		enVars[kv[0]] = kv[1]
	}
	return
}

// ParseString will parse any input provided as string
func ParseString(str string) (*template.Template, error) {
	funcMap := sprig.TxtFuncMap()
	return template.Must(template.New("").Funcs(funcMap).Parse(str)), nil
}

// MatchPrefix will match a given prefix pattern of all env variables and render only those.
func MatchPrefix(prefix string) map[string]interface{} {
	enVars := make(map[string]interface{})
	for _, value := range os.Environ() {
		kv := strings.SplitN(value, "=", 2)
		if strings.HasPrefix(kv[0], prefix) {
			enVars[kv[0]] = kv[1]
		}
	}
	return enVars
}

// GetPathInDir Recursively get all file paths in directory, including sub-directories.
func GetPathInDir(dirpath string) ([]string, error) {
	var paths []string
	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
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

func CreateDirIfNotExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.Mkdir(path, os.ModePerm); err != nil {
			return err
		}
		return err
	}
	return nil
}
