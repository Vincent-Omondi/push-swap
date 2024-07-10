package pushswap

import (
	"testing"
)

// // type Stack struct {
// //     elements []int
// // }

// // func NewStack() *Stack {
// //     return &Stack{elements: []int{}}
// // }

// func (s *Stack) Push(element int) {
//     s.elements = append(s.elements, element)
// }

// func (s *Stack) Pop() (int, error) {
//     if len(s.elements) == 0 {
//         return 0, errors.New("stack is empty")
//     }
//     element := s.elements[len(s.elements)-1]
//     s.elements = s.elements[:len(s.elements)-1]
//     return element, nil
// }

func TestPush(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	if len(stack.elements) != 2 || stack.elements[0] != 2 || stack.elements[1] != 1 {
		t.Error("Push failed")
	}
}

func TestPop(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	elem, err := stack.Pop()
	if err != nil || elem != 1 || len(stack.elements) != 0 {
		t.Error("Pop failed")
	}
}

func TestPa(t *testing.T) {
	a := NewStack()
	b := NewStack()
	b.Push(1)
	a.Push(2)
	err := Pa(a, b)
	if err != nil || len(a.elements) != 2 || len(b.elements) != 0 || a.elements[0] != 1 {
		t.Error("pa failed")
	}
}

func TestPb(t *testing.T) {
	a := NewStack()
	b := NewStack()
	a.Push(1)
	b.Push(2)
	err := Pb(a, b)
	if err != nil || len(b.elements) != 2 || len(a.elements) != 0 || b.elements[0] != 1 {
		t.Error("pb failed")
	}
}

func TestSa(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	Sa(stack)
	if stack.elements[0] != 1 || stack.elements[1] != 2 {
		t.Error("sa failed")
	}
}

func TestSb(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	Sb(stack)
	if stack.elements[0] != 1 || stack.elements[1] != 2 {
		t.Error("sb failed")
	}
}

func TestSs(t *testing.T) {
	a := NewStack()
	b := NewStack()
	a.Push(1)
	a.Push(2)
	b.Push(3)
	b.Push(4)
	Ss(a, b)
	if a.elements[0] != 1 || a.elements[1] != 2 || b.elements[0] != 3 || b.elements[1] != 4 {
		t.Error("ss failed")
	}
}

func TestRa(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	Ra(stack)
	if stack.elements[0] != 2 || stack.elements[1] != 1 || stack.elements[2] != 3 {
		t.Error("ra failed")
	}
}

func TestRb(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	Rb(stack)
	if stack.elements[0] != 2 || stack.elements[1] != 1 || stack.elements[2] != 3 {
		t.Error("rb failed")
	}
}

func TestRr(t *testing.T) {
	a := NewStack()
	b := NewStack()
	a.Push(1)
	a.Push(2)
	a.Push(3)
	b.Push(4)
	b.Push(5)
	b.Push(6)
	Rr(a, b)
	if a.elements[0] != 2 || a.elements[1] != 1 || a.elements[2] != 3 || b.elements[0] != 5 || b.elements[1] != 4 || b.elements[2] != 6 {
		t.Error("rr failed")
	}
}

func TestRra(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	Rra(stack)
	if stack.elements[0] != 1 || stack.elements[1] != 3 || stack.elements[2] != 2 {
		t.Error("rra failed")
	}
}

func TestRrb(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	Rrb(stack)
	if stack.elements[0] != 1 || stack.elements[1] != 3 || stack.elements[2] != 2 {
		t.Error("rrb failed")
	}
}

func TestRrr(t *testing.T) {
	a := NewStack()
	b := NewStack()
	a.Push(1)
	a.Push(2)
	a.Push(3)
	b.Push(4)
	b.Push(5)
	b.Push(6)
	Rrr(a, b)
	if a.elements[0] != 1 || a.elements[1] != 3 || a.elements[2] != 2 || b.elements[0] != 4 || b.elements[1] != 6 || b.elements[2] != 5 {
		t.Error("rrr failed")
	}
}


