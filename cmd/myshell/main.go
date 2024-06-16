package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	setEnvironment()
	for {
		writeToStdout("$ ")
		// wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			writeToStderr(err.Error())
			continue
		}

		command := getCommandFromInput(input)
		// execute builtin command
		if command.isBuiltin {
			command.execute(command.args)
			continue
		}

		// execute non-builtin command
		result := fmt.Sprintf("%s: command not found\n", command.name)
		writeToStderr(result)
	}
}
