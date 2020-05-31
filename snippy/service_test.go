package snippy

import (
	"errors"
	"reflect"
	"testing"
)

type fakeStore struct {
	data        string
	returnError bool
}

func (f *fakeStore) GetAll() (string, error) {
	if f.returnError {
		return "", errors.New("error")
	}
	return f.data, nil
}

func TestNewSnippetService(t *testing.T) {
	want := reflect.TypeOf(&SnippetService{})
	got := reflect.TypeOf(NewSnippetService(&fakeStore{}))
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestSnippetServiceGetAllSuccess(t *testing.T) {
	data := "Current directory [pwd]"
	service := NewSnippetService(&fakeStore{data: data})
	want := data
	got, _ := service.GetAll()
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestSnippetServiceGetAllError(t *testing.T) {
	service := NewSnippetService(&fakeStore{data: "", returnError: true})
	_, err := service.GetAll()
	if err == nil {
		t.Errorf("want %v, got %v", "error", err)
	}
}

func TestSnippetServiceGetAllNoSnippets(t *testing.T) {
	service := NewSnippetService(&fakeStore{data: ""})
	_, err := service.GetAll()
	if err == nil {
		t.Errorf("want %v, got %v", "error", err)
	}
}
