// main.go
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

	input := strings.Split(os.Args[1], " ")
	stackA, err := validate.ParseInput(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	if validate.IsSorted(stackA) {
		return
	}

	stackB := []int{}
	commands := []string{}

	// Stage 2: Sorting using calculated optimal commands
	for len(stackA) > 3 {
		targetIndex, minOps := -1, int(^uint(0)>>1)
		for i := 0; i < len(stackA); i++ {
			ops := calculateOperations(stackA, stackB, i)
			if ops < minOps {
				minOps = ops
				targetIndex = i
			}
		}
		moveToTop(&stackA, targetIndex, &commands, "ra", "rra")
		commands = append(commands, "pb")
		execute.ExecuteInstruction("pb", &stackA, &stackB)
	}

	sortThree(&stackA, &commands)

	// Stage 3: Push everything back to stackA
	for len(stackB) > 0 {
		position := findPositionToInsert(stackA, stackB[0])
		moveToTop(&stackA, position, &commands, "ra", "rra")
		commands = append(commands, "pa")
		execute.ExecuteInstruction("pa", &stackA, &stackB)
	}

	// Stage 4: Final arrangement to bring the minimum to the top
	minIndex := findMinIndex(stackA)
	moveToTop(&stackA, minIndex, &commands, "ra", "rra")

	for _, command := range commands {
		fmt.Println(command)
	}
}

func calculateOperations(stackA, stackB []int, targetIndex int) int {
	element := stackA[targetIndex]
	position := findPositionToInsert(stackB, element)
	ops := 0

	if targetIndex <= len(stackA)/2 {
		ops += targetIndex
	} else {
		ops += len(stackA) - targetIndex
	}

	if position <= len(stackB)/2 {
		ops += position
	} else {
		ops += len(stackB) - position
	}

	return ops
}

func moveToTop(stack *[]int, index int, commands *[]string, rotateCmd, reverseRotateCmd string) {
	if index <= len(*stack)/2 {
		for i := 0; i < index; i++ {
			*commands = append(*commands, rotateCmd)
			execute.ExecuteInstruction(rotateCmd, stack, nil)
		}
	} else {
		for i := 0; i < len(*stack)-index; i++ {
			*commands = append(*commands, reverseRotateCmd)
			execute.ExecuteInstruction(reverseRotateCmd, stack, nil)
		}
	}
}

func sortThree(stack *[]int, commands *[]string) {
	a, b, c := (*stack)[0], (*stack)[1], (*stack)[2]
	if a > b && b < c && a < c {
		*commands = append(*commands, "sa")
		execute.ExecuteInstruction("sa", stack, nil)
	} else if a > b && b > c && a > c {
		*commands = append(*commands, "sa", "rra")
		execute.ExecuteInstruction("sa", stack, nil)
		execute.ExecuteInstruction("rra", stack, nil)
	} else if a > b && b < c && a > c {
		*commands = append(*commands, "ra")
		execute.ExecuteInstruction("ra", stack, nil)
	} else if a < b && b > c && a < c {
		*commands = append(*commands, "sa", "ra")
		execute.ExecuteInstruction("sa", stack, nil)
		execute.ExecuteInstruction("ra", stack, nil)
	} else if a < b && b > c && a > c {
		*commands = append(*commands, "rra")
		execute.ExecuteInstruction("rra", stack, nil)
	}
}

func findPositionToInsert(stack []int, value int) int {
	for i, v := range stack {
		if v > value {
			return i
		}
	}
	return len(stack)
}

func findMinIndex(stack []int) int {
	minIdx, minVal := 0, stack[0]
	for i, v := range stack {
		if v < minVal {
			minIdx = i
			minVal = v
		}
	}
	return minIdx
}
