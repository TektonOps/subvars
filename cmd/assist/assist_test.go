package assist

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"text/template"
)

//nolint
func TestGetVars(t *testing.T) {
	tests := []struct {
		name       string
		wantEnVars map[string]string
	}{

		{"Get Variables", map[string]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEnVars := GetVars(); reflect.DeepEqual(gotEnVars, tt.wantEnVars) {
				t.Errorf("GetVars() = %v, want %v", gotEnVars, tt.wantEnVars)
			}
		})
	}
}

//nolint
func TestIsFlagSet(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"Set", "on", true},
		{"Not Set", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFlagSet(tt.args); got != tt.want {
				t.Errorf("IsFlagSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

//nolint
func TestMatchPrefix(t *testing.T) {
	_ = os.Setenv("TESTING_PREFIX_MATCH", "matched")
	tests := []struct {
		name   string
		prefix string
		want   map[string]string
	}{
		{"Prefix match test", "TESTING_", map[string]string{"TESTING_PREFIX_MATCH": "matched"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchPrefix(tt.prefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

//nolint
func TestParseFile(t *testing.T) {
	tmpFile, _ := ioutil.TempFile(os.TempDir(), "prefix-")
	filename := tmpFile.Name()
	defer os.Remove(tmpFile.Name())
	tests := []struct {
		name    string
		file    string
		wantErr bool
	}{
		{"Parse files", filename, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseFile(tt.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

//nolint
func TestParseString(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    *template.Template
		wantErr bool
	}{
		{"Parse Success", "{{ .USER }}", template.New("{{ .USER }}"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseString(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
