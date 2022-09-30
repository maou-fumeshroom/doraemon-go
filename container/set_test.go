package container

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	fmt.Println(s.AddFirst(1))
	fmt.Println(s.AddFirst(2))
	s.AddFirst(5)
	fmt.Println(s.String())
	fmt.Println(s.Len())
	s.Remove(1)
	s.Range(func(el int) bool {
		fmt.Println("range: ", el)
		return true
	})
	ss := NewSet[int]()
	ss.Add(3)
	s.Merge(ss)
	fmt.Println(s)
	s.Clear()
	fmt.Println(s)
}
