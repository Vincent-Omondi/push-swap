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
	}
}
