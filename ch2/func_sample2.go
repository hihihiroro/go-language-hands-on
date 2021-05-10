package main

import (
	"fmt"
)

func main() {
	m := []string{"one", "two", "three"}
	fmt.Println(m)
	//	m = insert(m, "*", 2)
	//	m = insert(m, "*", 1)
	m = push(m, "1", "2", "3")
	fmt.Println(m)
}

func insert(a []string, v string, p int) (s []string) {
	s = append(a, "")
	s = append(s[:p+1], s[p:len(s)-1]...)
	s[p] = v
	return
}

func push(a []string, v ...string) (s []string) {
	s = append(a, v...)
	return
}
