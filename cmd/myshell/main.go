package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Returns ([command], [arguments]) from input
func getArgumentsFromInput(input string) (string, []string) {
	args := strings.Split(input, " ")
	return args[0], args[1:]
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
		command, args := getArgumentsFromInput(input)
		switch command {
		case "exit":
			parseExitCommand(args)
		case "echo":
			parseEchoCommand(args)
		default:
			result := fmt.Sprintf("%s: command not found\n", command)
			fmt.Fprint(os.Stderr, result)
		}
	}
}
