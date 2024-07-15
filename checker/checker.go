package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/push-swap/pkg/validate"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	stackA, err := validate.ParseInput(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
	}

	stackB := []int{}
	scanner := bufio.NewScanner(os.Stdin)
}
