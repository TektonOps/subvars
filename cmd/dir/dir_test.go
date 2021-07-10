package dir

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

//nolint
func Test_createDirIfNotExist(t *testing.T) {
	dir, _ := ioutil.TempDir("", "example")
	defer os.RemoveAll(dir) // clean up
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"Cerate Dir Success", dir, false},
		{"Crate Dir Failure", "/temp/t2", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createDirIfNotExist(tt.path); (err != nil) != tt.wantErr {
				t.Errorf("createDirIfNotExist() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//nolint
func Test_getPathInDir(t *testing.T) {
	content := []byte("temporary file's content")
	dir, _ := ioutil.TempDir("", "example")
	defer os.RemoveAll(dir)
	tmpfn := filepath.Join(dir, "tmpfile")
	_ = ioutil.WriteFile(tmpfn, content, 0666)

	tests := []struct {
		name    string
		args    string
		want    []string
		wantErr bool
	}{
		{"Get Path Test", dir, []string{tmpfn}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPathInDir(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPathInDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPathInDir() got = %v, want %v", got, tt.want)
			}
		})
	}
}
