package container

import (
	"fmt"
	"strings"
)

type Set[T comparable] map[T]struct{}

// @in_param els 初始值
func NewSet[T comparable](els ...T) Set[T] {
	set := make(Set[T], len(els))
	for _, el := range els {
		set[el] = struct{}{}
	}
	return set
}

func (s Set[T]) Add(el T) {
	s[el] = struct{}{}
}

func (s Set[T]) Remove(el T) {
	delete(s, el)
}

func (s Set[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s Set[T]) AddFirst(el T) bool {
	if _, ok := s[el]; ok {
		return false
	}
	s.Add(el)
	return true
}

func (s Set[T]) Merge(t Set[T]) {
	for k, v := range t {
		s[k] = v
	}
}

func (s Set[T]) Range(f func(el T) bool) {
	for k, _ := range s {
		if !f(k) {
			break
		}
	}
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) String() string {
	strs := make([]string, 0, s.Len())
	for el, _ := range s {
		strs = append(strs, fmt.Sprintf("%+v", el))
	}
	return fmt.Sprintf("[%s]", strings.Join(strs, ", "))
}
