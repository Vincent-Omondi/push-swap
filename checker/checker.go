package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/push-swap/pkg/execute"
	"github.com/push-swap/pkg/validate"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	// Split the input args into individual numbers
	inputArgs := strings.Split(os.Args[1], " ")
	stackA, err := validate.ParseInput(inputArgs)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	stackB := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	instructions := []string{}

	for scanner.Scan() {
		instruction := scanner.Text()
		if instruction == "" {
			break
		}
		instructions = append(instructions, instruction)
	}

	for _, instruction := range instructions {
		if err := execute.ExecuteInstruction(instruction, &stackA, &stackB); err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
	}

	if validate.IsSorted(stackA) && len(stackB) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
