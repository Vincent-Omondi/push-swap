package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/push-swap/pkg/execute"
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
