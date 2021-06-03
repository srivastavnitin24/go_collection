package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"something": 10, "yo": 20, "blah": 20}

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	fmt.Println("by val : ", ss)

	//sort.Slice(ss, func(i, j int) bool {
	//	return ss[i].Key > ss[j].Key
	//})
	//fmt.Println("by key : ",ss)

	for _, kv := range ss {
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
	}
}

//==================================

//package main
//
//import (
//"fmt"
//"sort"
//)
//
//func main() {
//	// To create a map as input
//	m := make(map[int]string)
//	m[1] = "a"
//	m[2] = "c"
//	m[0] = "b"
//
//	keys := make([]string, len(m))
//	i := 0
//	for _, v := range m {
//		keys[i] = v
//		i++
//	}
//	sort.Strings(keys)
//	fmt.Println("Key:", keys)
//
//	// To perform the opertion you want
//	for _, k := range keys {
//		fmt.Println("Key:", k, "Value:", m[k])
//	}
//}
