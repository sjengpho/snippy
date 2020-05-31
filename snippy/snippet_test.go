package snippy

import (
	"testing"
)

func TestParseSnippetSuccess(t *testing.T) {
	tt := []string{
		"Current directory [pwd]",
		"List [ls]",
		"Move or rename files [mv]",
	}

	for _, input := range tt {
		_, err := ParseSnippet(input)
		if err != nil {
			t.Errorf("want %v, got %v", nil, err)
		}
	}
}

func TestParseSnippetError(t *testing.T) {
	tt := []string{
		"Current directory []",
		"Current directory [pwd",
		"Current directory [",
		"Current directory ]",
		"Current directory ]pwd[",
		"Current directory (pwd)",
		"Current directory pwd",
		"Current directory _pwd_",
		"Current directory [pwd)",
		"Current directory (pwd]",
		"[pwd]",
		"Current directory",
		"[pwd] Current directory",
	}

	for _, input := range tt {
		_, err := ParseSnippet(input)
		if err != ErrInvalidSnippet {
			t.Errorf("want %v, got %v", ErrInvalidSnippet, err)
		}
	}
}

func TestSnippetValue(t *testing.T) {
	snippet := Snippet("Current directory [pwd]")
	want := "pwd"
	got := snippet.Value()
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}
