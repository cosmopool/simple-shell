package env

import (
	"os"
	"strings"
)

const PATH = "PATH"
const PWD = "PWD"

// All enviromental variables from this shell session
var SessionEnv Environment

// Enviromental variables from shell
type Environment struct {
	Path []string
	Pwd  string
}

// Set shell environment variables
func SetSessionEnvironment() {
	SessionEnv = Environment{
		Path: strings.Split(os.Getenv(PATH), ":"),
		Pwd:  os.Getenv(PWD),
	}
}
