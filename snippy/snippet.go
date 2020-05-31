package snippy

import (
	"errors"
	"regexp"
	"strings"
)

// ErrInvalidSnippet means the string doesn't have the format snippy.Snippet is expected to have.
var ErrInvalidSnippet = errors.New("invalid snippet format")

// snippetRegexp defines the expected format of a snippy.Snippet. The string should start
// with one or more characters and end with one or more characters between square brackets.
// Example: Current directory [pwd]
var snippetRegexp = regexp.MustCompile(`^.+\[.+\]$`)

// Snippet represents a snippet.
type Snippet string

// Value returns the value between the square brackets.
func (s Snippet) Value() string {
	bracketIndex := strings.Index(string(s), "[")
	value := string(s[bracketIndex+1:]) // Extract everything after the first bracket.
	value = strings.TrimSpace(value)
	value = strings.Trim(value, "]")
	return value
}

// ParseSnippet parses the string and returns a snippy.Snippet.
func ParseSnippet(s string) (Snippet, error) {
	if !snippetRegexp.MatchString(s) {
		return Snippet(""), ErrInvalidSnippet
	}

	return Snippet(s), nil
}
