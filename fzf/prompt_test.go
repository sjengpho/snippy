package fzf

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"testing"
)

func fakeLookPath(file string) (string, error) {
	return "", errors.New("error")
}

func fakeExecCommand(commandName string) func(name string, args ...string) *exec.Cmd {
	return func(name string, args ...string) *exec.Cmd {
		cs := []string{fmt.Sprintf("-test.run=%v", commandName), "--", name}
		cs = append(cs, args...)
		cmd := exec.Command(os.Args[0], cs...)
		cmd.Env = []string{"GO_TEST_PROCESS=1"}
		return cmd
	}
}

func TestNewPromptSuccess(t *testing.T) {
	prompt, _ := NewPrompt()
	want := reflect.TypeOf(&Prompt{})
	got := reflect.TypeOf(prompt)
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestNewPromptError(t *testing.T) {
	lookPath = fakeLookPath
	defer func() { lookPath = exec.LookPath }()

	_, err := NewPrompt()
	if err == nil {
		t.Errorf("want %v, got %v", "error", err)
	}
}

func TestPromptAskSuccess(t *testing.T) {
	execCommand = fakeExecCommand("TestPromptAskExecCommandSuccess")
	defer func() { execCommand = exec.Command }()

	prompt, _ := NewPrompt()
	want := "Current directory [pwd]"
	got, _ := prompt.Ask("")
	if string(got) != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestPromptAskError(t *testing.T) {
	execCommand = fakeExecCommand("TestPromptAskExecCommandError")
	defer func() { execCommand = exec.Command }()

	prompt, _ := NewPrompt()
	_, err := prompt.Ask("")
	if err == nil {
		t.Errorf("want %v, got %v", "error", err)
	}
}

func TestPromptAskExecCommandSuccess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}

	fmt.Println("Current directory [pwd]")
	os.Exit(0)
}

func TestPromptAskExecCommandError(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}

	os.Exit(1)
}
