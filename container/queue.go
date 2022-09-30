package container

import (
	"fmt"
	"strings"
)

type Queue[T any] struct {
	null T
	sl   []T
}

// @in_param size 初始容量
func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		sl: make([]T, 0, size),
	}
}

func (q *Queue[T]) Push(el T) {
	q.sl = append(q.sl, el)
}

// @out_param T 值
// @out_param bool 是否为空 为空时 T 为 null
func (q *Queue[T]) Pop() (T, bool) {
	if q.IsEmpty() {
		return q.null, true
	}
	rlt := q.sl[0]
	q.sl = q.sl[1:]
	return rlt, false
}

func (q *Queue[T]) Clear() {
	q.sl = make([]T, 0)
}

func (q *Queue[T]) Len() int {
	return len(q.sl)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue[T]) String() string {
	strs := make([]string, 0, q.Len())
	for _, el := range q.sl {
		strs = append(strs, fmt.Sprintf("%v", el))
	}
	return fmt.Sprintf("[%s]", strings.Join(strs, ", "))
}
