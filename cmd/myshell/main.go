package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Returns ([command], [arguments]) from input
func getArgumentsFromCommand(input string) (string, []string) {
	args := strings.Split(input, " ")
	return args[0], args[1:]
}

func parseExitCommand(args []string) {
	var exitCode int

	if len(args) == 0 {
		os.Exit(0)
	}

	exitCode, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(exitCode)
	}
	os.Exit(exitCode)
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		rawInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			os.Stderr.Write([]byte(fmt.Sprint(err)))
			os.Exit(1)
		}

		input := strings.Trim(rawInput, "\n")
		command, args := getArgumentsFromCommand(input)
		switch command {
		case "exit":
			parseExitCommand(args)
		default:
			result := fmt.Sprintf("%s: command not found\n", command)
			fmt.Fprint(os.Stderr, result)
		}
	}
}
