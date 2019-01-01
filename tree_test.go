package tree

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func getPathOfThisFile() (path string, err error) {
	path, err = os.Getwd()
	return
}

func TestItemGetName(t *testing.T) {
	name := "test 1"
	item := Item{name: name}
	obtainedName := GetName(item)
	if obtainedName != name {
		t.Errorf("Wrong name returned. Expected: %s, obtained: %s", name, obtainedName)
	}
}

func TestItemGetPath(t *testing.T) {
	path := "test 1"
	item := Item{path: path}
	obtainedPath := GetPath(item)
	if obtainedPath != path {
		t.Errorf("Wrong path returned. Expected: %s, obtained: %s", path, obtainedPath)
	}
}

func TestGetItemsInFolder(t *testing.T) {
	currentPath, error := getPathOfThisFile()
	if error != nil {
		t.Errorf("Error: %s", error)
	}
	folderPath := filepath.Join(currentPath, "test_data")
	content, error := ioutil.ReadDir(folderPath)
	if error != nil {
		t.Errorf("Error: %s", error)
	}

	if len(content) == 0 {
		t.Error("Content has length 0")
	}
	for _, file := range content {
		t.Log(file.Name())
	}
}
