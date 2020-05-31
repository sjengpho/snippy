package snippy

import (
	"fmt"
	"os"
	"strings"
)

// Config represents the configuration.
type Config struct {
	FileStorePath string
}

// DefaultConfig returns a snippy.Config with default values.
func DefaultConfig() Config {
	home, _ := os.UserHomeDir()
	dir := fmt.Sprintf("/%v/.config/snippy", strings.TrimLeft(home, "/"))
	return Config{FileStorePath: dir + "/data"}
}
