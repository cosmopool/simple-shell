package builtin

import (
	"fmt"
	"os"
	"strings"
)

func executeEchoCommand(args []string) {
	fmt.Fprintln(os.Stdout, strings.Join(args, " "))
}
