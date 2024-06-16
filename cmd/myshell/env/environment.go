package env

import (
	"os"
	"strings"
)

const PATH = "PATH"

// All enviromental variables from this shell session
var SessionEnv Environment

// Enviromental variables from shell
type Environment struct {
	Path []string
}

// Set shell environment variables
func SetEnvironment() {
	path := strings.Split(os.Getenv(PATH), ":")
	SessionEnv = Environment{
		Path: path,
	}
}
