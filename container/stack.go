package container

import (
	"fmt"
	"strings"
)

type Stack[T any] struct {
	null T
	sl   []T
}

// @in_param size 栈初始长度
func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{
		sl: make([]T, 0, size),
	}
}

func (s *Stack[T]) Push(el ...T) {
	s.sl = append(s.sl, el...)
}

// @out_param T 值
// @out_param bool 栈是否为空，true 时 T 为空值
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		return s.null, true
	}
	el := s.sl[s.Len()-1]
	s.sl = s.sl[:s.Len()-1]
	return el, false
}

func (s *Stack[T]) Clear() {
	s.sl = make([]T, 0)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Len() int {
	return len(s.sl)
}

func (s *Stack[T]) String() string {
	strs := make([]string, 0, s.Len())
	for _, el := range s.sl {
		strs = append(strs, fmt.Sprintf("%+v", el))
	}
	return fmt.Sprintf("[%s]", strings.Join(strs, ", "))
}
