package cobraCmd

import (
	"fmt"
	"os"

	"Product-Management-API/pkg/cmd"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
