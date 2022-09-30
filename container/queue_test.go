package container

import (
	"fmt"
	"testing"
)

func TestQueu(t *testing.T) {
	q := NewQueue[int](0)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	l, _ := q.Pop()
	fmt.Println(l)
	l, _ = q.Pop()
	fmt.Println(l)
	fmt.Println(q.String())
}
