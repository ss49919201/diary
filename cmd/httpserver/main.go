package main

import (
	"fmt"
	"os"

	"github.com/ss49919201/diary/customlogger"
	"github.com/ss49919201/diary/server"
)

func main() {
	if err := server.Run(customlogger.NewLogger()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
