package env

import (
	"fmt"
	"os"
	"strings"
)

const PATH = "PATH"
const PWD = "PWD"
const HOME = "HOME"

// All enviromental variables from this shell session
var SessionEnv Environment

// Enviromental variables from shell
type Environment struct {
	Path []string
	Pwd  string
	Home string
}

func (e *Environment) SetEnv(env string, val string) error {
	switch env {
	case PATH:
		e.Path = strings.Split(val, ":")
	case PWD:
		e.Pwd = val
	case HOME:
		e.Home = val
	default:
		return fmt.Errorf("No such variable with this name: %s", env)
	}

	return os.Setenv(env, val)
}

// Set shell environment variables
func SetSessionEnvironment() {
	SessionEnv = Environment{
		Path: strings.Split(os.Getenv(PATH), ":"),
		Pwd:  os.Getenv(PWD),
		Home:  os.Getenv(HOME),
	}
}
