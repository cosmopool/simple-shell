package main

import (
	"fmt"
	"os"
	"strings"
)

func parseEchoCommand(args []string) {
	fmt.Fprintln(os.Stdout, strings.Join(args, " "))
}
