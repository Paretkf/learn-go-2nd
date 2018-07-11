package main

import (
	"errors"
	"fmt"
)

type stack []int

func main() {
	// var a []string
	// a = append(a, "a", "b", "c", "d", "e", "f", "i", "g", "h", "i")
	// // fmt.Println(a[:2])

	// // b := append([]string(nil), a...) // copy
	// // fmt.Println(b)

	// c := append(a, a...)
	// fmt.Println(c)

	// fmt.Println(reverse(a))

	var a stack
	a = append(a, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(a.push(20))
	fmt.Println(a.pop())
}

func (s stack) push(a int) stack {
	s = append(s, a)
	return s
}

func (s stack) pop() (int, error) {
	if len(s) == 0 {
		return 0, errors.New("type")
	}
	a := s[len(s)-1]
	s = append(s[len(s):], s[:len(s)-1]...)
	return a, nil
}

func reverse(arr []string) []string {
	a := arr
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}
