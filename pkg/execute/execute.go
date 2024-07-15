package execute

import "errors"

func ExecuteInstruction(instruction string, stackA, stackB *[]int) error {
	switch instruction {
	case "sa":
		swap(stackA)
	case "sb":
		swap(stackB)
	case "ss":
		swap(stackA)
		swap(stackB)
	case "pa":
		push(stackB, stackA)
	case "pb":
		push(stackA, stackB)
	case "ra":
		rotate(stackA)
	case "rb":
		rotate(stackB)
	case "rr":
		rotate(stackA)
		rotate(stackB)
	case "rra":
		reverseRotate(stackA)
	case "rrb":
		reverseRotate(stackB)
	case "rrr":
		reverseRotate(stackA)
		reverseRotate(stackB)
	default:
		return errors.New("invalid instruction")
	}
	return nil
}

func swap(stack *[]int) {
	if len(*stack) < 2 {
		return
	}
	(*stack)[0], (*stack)[1] = (*stack)[1], (*stack)[0]
}

func push(fromStack, toStack *[]int) {
	if len(*fromStack) == 0 {
		return
	}
	*toStack = append([]int{(*fromStack)[0]}, *toStack...)
	*fromStack = (*fromStack)[1:]
}

func rotate(stack *[]int) {
	if len(*stack) < 2 {
		return
	}
	*stack = append((*stack)[1:], (*stack)[0])
}

func reverseRotate(stack *[]int) {
	if len(*stack) < 2 {
		return
	}
	*stack = append([]int{(*stack)[len(*stack)-1]}, (*stack)[:len(*stack)-1]...)
}
