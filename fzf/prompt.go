package fzf

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sjengpho/snippy/snippy"
)

var lookPath = exec.LookPath
var execCommand = exec.Command

// Prompt implements snippy.Prompt.
type Prompt struct{}

// NewPrompt returns a fzf.Prompt.
func NewPrompt() (*Prompt, error) {
	if _, err := lookPath("fzf"); err != nil {
		return nil, fmt.Errorf("failed creating prompt: %w", err)
	}

	return &Prompt{}, nil
}

// Ask uses fzf (https://github.com/junegunn/fzf) that the user can
// use to filter and choose a snippet.
func (p *Prompt) Ask(snippets string) (snippy.Snippet, error) {
	var buf bytes.Buffer
	cmd := execCommand("fzf", "--no-preview")
	cmd.Stdin = strings.NewReader(snippets)
	cmd.Stdout = &buf
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return snippy.Snippet(""), err
	}
	return snippy.ParseSnippet(strings.TrimSpace(buf.String()))
}
