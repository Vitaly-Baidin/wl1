package main

import (
	"fmt"
	"strings"
)

// TODO| Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

type Set struct {
	data map[string]struct{}
}

func NewSet() *Set {
	return &Set{
		data: make(map[string]struct{}),
	}
}

func (s *Set) Add(element string) bool {
	_, ok := s.data[element]
	if !ok {
		s.data[element] = struct{}{}
	}

	return !ok
}

func (s *Set) Delete(element string) bool {
	_, ok := s.data[element]
	if ok {
		delete(s.data, element)
	}

	return ok
}

func (s *Set) String() string {
	var builder strings.Builder
	i := 0
	for elem := range s.data {
		if i == 0 {
			builder.WriteString("[")
		} else {
			builder.WriteString(" ")
		}
		builder.WriteString(elem)
		i++
	}
	builder.WriteString("]")
	return builder.String()
}

func main() {
	set := NewSet()
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	for _, e := range arr {
		set.Add(e)
	}

	fmt.Println(set)
}
