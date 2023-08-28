package main

import "fmt"

func deleteFromSlice(s []int, i int) []int {
	if i >= len(s) {
		panic("index out of range")
	}

	return append(s[:i], s[i+1:]...)
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = deleteFromSlice(s, 2)
	fmt.Println(s)
}
