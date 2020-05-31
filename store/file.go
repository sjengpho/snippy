package store

import (
	"fmt"
	"io/ioutil"
	"os"
)

var readFile = ioutil.ReadFile
var osStat = os.Stat

// FileStore implements snippy.Store.
type FileStore struct {
	path string
}

// NewFileStore returns a store.FileStore.
func NewFileStore(path string) (FileStore, error) {
	store := FileStore{path}

	if _, err := osStat(path); err != nil {
		if !os.IsNotExist(err) {
			return store, err
		}
		return store, store.createNewDataFile()
	}

	return store, nil
}

// GetAll reads the file and returns the contents as string.
//
// It assumes that the contents of the file are snippets separated by
// newlines and defined in the format as described by snippy.snippetRegexp
func (f FileStore) GetAll() (string, error) {
	snippets, err := readFile(f.path)
	if err != nil {
		return "", fmt.Errorf("failed getting snippets: %w", err)
	}
	return string(snippets), nil
}

// createNewDataFile creates a new data file.
//
// If the file already exists it will be overwritten.
func (f FileStore) createNewDataFile() error {
	const exampleData = `Current directory [pwd]
List [ls]
Move or rename files [mv]`

	file, err := os.Create(f.path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	file.WriteString(exampleData)
	return file.Sync()
}
