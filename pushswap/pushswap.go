package pushswap

import "errors"

type Stack struct {
    elements []int
}

func NewStack() *Stack {
    return &Stack{elements: []int{}}
}

func (s *Stack) Push(element int) {
    s.elements = append([]int{element}, s.elements...)
}

func (s *Stack) Pop() (int, error) {
    if len(s.elements) == 0 {
        return 0, errors.New("stack is empty")
    }
    element := s.elements[0]
    s.elements = s.elements[1:]
    return element, nil
}

func Pa(stackA, stackB *Stack) error {
    if len(stackB.elements) == 0 {
        return errors.New("stack B is empty")
    }
    element := stackB.elements[0]
    stackB.elements = stackB.elements[1:]
    stackA.elements = append([]int{element}, stackA.elements...)
    return nil
}

func Pb(stackA, stackB *Stack) error {
    if len(stackA.elements) == 0 {
        return errors.New("stack A is empty")
    }
    element := stackA.elements[0]
    stackA.elements = stackA.elements[1:]
    stackB.elements = append([]int{element}, stackB.elements...)
    return nil
}

func Sa(stackA *Stack) error {
    if len(stackA.elements) < 2 {
        return errors.New("stack A has fewer than 2 elements")
    }
    stackA.elements[0], stackA.elements[1] = stackA.elements[1], stackA.elements[0]
    return nil
}

func Sb(stackB *Stack) error {
    if len(stackB.elements) < 2 {
        return errors.New("stack B has fewer than 2 elements")
    }
	stackB.elements[0], stackB.elements[1] = stackB.elements[1], stackB.elements[0]
    return nil
}

func Ss(stackA, stackB *Stack) error {
    if err := Sa(stackA); err != nil {
        return err
    }
    return Sb(stackB)
}

func Ra(stackA *Stack) error {
    if len(stackA.elements) < 1 {
        return errors.New("stack A is empty")
    }
    first := stackA.elements[0]
    stackA.elements = append(stackA.elements[1:], first)
    return nil
}

func Rb(stackB *Stack) error {
    if len(stackB.elements) < 1 {
        return errors.New("stack B is empty")
    }
    first := stackB.elements[0]
    stackB.elements = append(stackB.elements[1:], first)
    return nil
}

func Rr(stackA, stackB *Stack) error {
    if err := Ra(stackA); err != nil {
        return err
    }
    return Rb(stackB)
}

func Rra(stackA *Stack) error {
    if len(stackA.elements) < 1 {
        return errors.New("stack A is empty")
    }
    last := stackA.elements[len(stackA.elements)-1]
    stackA.elements = append([]int{last}, stackA.elements[:len(stackA.elements)-1]...)
    return nil
}

func Rrb(stackB *Stack) error {
    if len(stackB.elements) < 1 {
        return errors.New("stack B is empty")
    }
    last := stackB.elements[len(stackB.elements)-1]
    stackB.elements = append([]int{last}, stackB.elements[:len(stackB.elements)-1]...)
    return nil
}

func Rrr(stackA, stackB *Stack) error {
    if err := Rra(stackA); err != nil {
        return err
    }
    return Rrb(stackB)
}