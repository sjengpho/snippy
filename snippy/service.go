package snippy

import (
	"errors"
	"strings"
)

// Store is the interface implemented by an object that provides
// access to the snippet storage.
//
// GetAll returns a multiline string in which snippets are separated by
// newlines. See snippy.snippetRegexp for the expected format of a snippet.
type Store interface {
	GetAll() (string, error)
}

// SnippetService provides access to snippets.
type SnippetService struct {
	store Store
}

// NewSnippetService returns a snippy.SnippetService.
func NewSnippetService(store Store) *SnippetService {
	return &SnippetService{store}
}

// GetAll returns a multiline string. See the interface for more info.
func (s *SnippetService) GetAll() (string, error) {
	snippets, err := s.store.GetAll()
	if err != nil {
		return "", err
	}

	if strings.TrimSpace(snippets) == "" {
		return "", errors.New("no snippets available")
	}

	return snippets, nil
}
