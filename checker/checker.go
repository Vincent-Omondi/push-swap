package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	stackA, err := checker.ParseInput(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
	}
}
