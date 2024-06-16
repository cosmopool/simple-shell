package main

import (
	"os"
	"strings"
)

const PATH = "PATT"

var environment Environment

type Environment struct {
	path []string
}

// Configure shell environment variables
func setEnvironment() {
	path := strings.Split(os.Getenv(PATH), ":")
	environment = Environment{
		path: path,
	}
}
