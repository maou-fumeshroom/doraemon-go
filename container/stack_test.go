package container

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack[int](0)
	s.Push(1)
	n, empty := s.Pop()
	fmt.Println(n)
	fmt.Println(empty)
	s.Push(2)
	s.Push(3)
	n, empty = s.Pop()
	fmt.Println(n)
	fmt.Println(empty)
	fmt.Println(s.String())
}
