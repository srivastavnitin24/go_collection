package main

import (
	"fmt"
	"sort"
)

var exists = struct{}{}

type set struct {
	m map[string]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *set) Set(str string) {
	s.m[str] = exists
}
func (s *set) Delete(str string) {
	delete(s.m, str)
}
func (s *set) Contains(str string) bool {
	_, ok := s.m[str]
	return ok
}

func main() {
	s := NewSet()
	s.Set("nitin")
	s.Set("mayank")
	s.Set("mohit")
	s.Set("india")
	s.Set("fiji")
	s.Set("Ao")
	//fmt.Println(s.Contains("nitin"))
	//s.Delete("mohit")
	//fmt.Println(s.Contains("mohit"))

	fmt.Println("before sorting")
	for k := range s.m {
		fmt.Println(k)
	}
	fmt.Println("==========================")

	type ss struct {
		key string
	}
	var stru []ss
	for k, _ := range s.m {
		stru = append(stru, ss{k})
	}

	sort.Slice(stru, func(i, j int) bool {
		return stru[i].key < stru[j].key
	})

	fmt.Println("after sorting")
	for _, v := range stru {
		fmt.Println(v.key)
	}
}
