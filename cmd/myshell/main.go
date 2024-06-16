package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// wait for user input
		rawInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			writeToStderr(err.Error())
			continue
		}

		command := getCommandFromInput(rawInput)
		if command.isBuiltin {
			command.execute(command.args)
			continue
		}

		// execute non-builtin command
		result := fmt.Sprintf("%s: command not found\n", command.name)
		writeToStderr(result)
	}
}
