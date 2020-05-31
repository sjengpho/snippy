package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/sjengpho/snippy/fzf"
	"github.com/sjengpho/snippy/snippy"
	"github.com/sjengpho/snippy/store"
)

func main() {
	config := snippy.DefaultConfig()
	store, err := store.NewFileStore(config.FileStorePath)
	if err != nil {
		log.Fatal(err)
	}
	service := snippy.NewSnippetService(store)

	snippets, err := service.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	prompt, err := fzf.NewPrompt()
	if err != nil {
		log.Fatal(err)
	}

	snippet, err := prompt.Ask(snippets)
	if err != nil {
		var e *exec.ExitError
		if errors.As(err, &e) {
			os.Exit(0)
		}
		log.Fatal(err)
	}

	fmt.Print(snippet.Value())
}
