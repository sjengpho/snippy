package store

import (
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"
)

type fakeFileInfo struct{}

func (f fakeFileInfo) Name() string {
	return "file name"
}

func (f fakeFileInfo) Size() int64 {
	return 1
}

func (f fakeFileInfo) Mode() os.FileMode {
	return os.ModeTemporary
}

func (f fakeFileInfo) ModTime() time.Time {
	return time.Now()
}

func (f fakeFileInfo) IsDir() bool {
	return f.Mode().IsDir()
}

func (f fakeFileInfo) Sys() interface{} {
	return nil
}

func fakeReadFile(filename string) ([]byte, error) {
	return []byte("snippy"), nil
}

func fakeReadFileError(filename string) ([]byte, error) {
	return []byte(""), errors.New("error")
}

func fakeOsStatSuccess(name string) (os.FileInfo, error) {
	return fakeFileInfo{}, nil
}

func fakeOsStatError(name string) (os.FileInfo, error) {
	return nil, errors.New("error")
}

func TestNewFileStoreSuccess(t *testing.T) {
	defer func() { os.Remove("data") }()
	store, _ := NewFileStore("data")
	want := reflect.TypeOf(FileStore{})
	got := reflect.TypeOf(store)
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestNewFileStoreCreateDataFileSuccess(t *testing.T) {
	osStat = fakeOsStatSuccess
	defer func() { osStat = os.Stat }()
	defer func() { os.Remove("data") }()
	store, _ := NewFileStore("data")
	want := reflect.TypeOf(FileStore{})
	got := reflect.TypeOf(store)
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestNewFileStoreError(t *testing.T) {
	osStat = fakeOsStatError
	defer func() { osStat = os.Stat }()
	defer func() { os.Remove("data") }()
	_, err := NewFileStore("data")
	if err == nil {
		t.Errorf("want %v, got %v", "error", err)
	}
}

func TestFileStoreGetAllSuccess(t *testing.T) {
	readFile = fakeReadFile
	defer func() { readFile = ioutil.ReadFile }()

	store, _ := NewFileStore("")
	want := "snippy"
	got, _ := store.GetAll()
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestFileStoreGetAllError(t *testing.T) {
	readFile = fakeReadFileError
	defer func() { readFile = ioutil.ReadFile }()

	store, _ := NewFileStore("")
	_, err := store.GetAll()
	if err == nil {
		t.Errorf("want %v, got %v", "error", err)
	}
}
