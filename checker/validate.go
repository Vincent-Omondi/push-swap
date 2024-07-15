package checker

import (
	"errors"
	"strconv"
)

func ParseInput(args []string) ([]int, error) {
	stack := []int{}
	seen := map[int]bool{}

	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			return nil, err
		}
		if seen[num] {
			return nil, errors.New("duplicate number found")
		}
		seen[num] = true
		stack = append(stack, num)
	}
	return stack, nil
}

func IsSorted(stack []int) bool {
	for i := 0; i < len(stack)-1; i++ {
		if stack[i] > stack[i+1] {
			return false
		}
	}
	return true
}
