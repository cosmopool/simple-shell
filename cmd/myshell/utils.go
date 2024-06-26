package main

import (
	"fmt"
	"os"
)

func writeToStderr(errorString string) {
	_, err := fmt.Fprint(os.Stderr, errorString)
	if err != nil {
		msg := fmt.Sprintf("%v ocurred when trying to write %v to STDERR.", err, errorString)
		panic(msg)
	}
}

func writeToStdout(outString string) {
	_, err := fmt.Fprint(os.Stdout, outString)
	if err != nil {
		msg := fmt.Sprintf("%v ocurred when trying to write %v to STDOUT.", err, outString)
		panic(msg)
	}
}
