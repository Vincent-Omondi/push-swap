// pushswap/push-swap.go
package main

import (
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

	inputArgs := strings.Split(os.Args[1], " ")
	stackA, err := validate.ParseInput(inputArgs)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	instructions := generateInstructions(stackA)

	for _, instruction := range instructions {
		fmt.Println(instruction)
	}
}

// Helper function to execute and store instructions
func addInstruction(instruction string, stackA, stackB *[]int, instructions *[]string) {
	*instructions = append(*instructions, instruction)
	execute.ExecuteInstruction(instruction, stackA, stackB)
}

// Function to find the index of the smallest element
func findMinIndex(stack []int) int {
	minIndex := 0
	for i := 1; i < len(stack); i++ {
		if stack[i] < stack[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

// Function to generate the sequence of instructions
func generateInstructions(stackA []int) []string {
	stackB := []int{}
	instructions := []string{}

	// If the stack is already sorted, return immediately
	if validate.IsSorted(stackA) {
		return instructions
	}

	for len(stackA) > 0 {
		minIndex := findMinIndex(stackA)

		// Bring the smallest element to the top
		if minIndex > len(stackA)/2 {
			for i := len(stackA) - 1; i >= minIndex; i-- {
				addInstruction("rra", &stackA, &stackB, &instructions)
			}
		} else {
			for i := 0; i < minIndex; i++ {
				addInstruction("ra", &stackA, &stackB, &instructions)
			}
		}

		// Push the smallest element to stack B
		addInstruction("pb", &stackA, &stackB, &instructions)

		// Try to sort within stack A
		if len(stackA) > 1 && stackA[0] > stackA[1] {
			addInstruction("sa", &stackA, &stackB, &instructions)
		}
	}

	// Push all elements back to stack A
	for len(stackB) > 0 {
		addInstruction("pa", &stackA, &stackB, &instructions)
	}

	return instructions
}
