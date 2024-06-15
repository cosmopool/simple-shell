package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	rawInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		os.Stderr.Write([]byte(fmt.Sprint(err)))
		os.Exit(1)
	}

	command := strings.Trim(rawInput, "\n")
	for {
		if command == "exit" {
			os.Exit(0)
		}
		result := fmt.Sprintf("%s: command not found", command)
		os.Stderr.Write([]byte(result))
	}
}
