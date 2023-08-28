package main

import (
	"fmt"
	"strings"
)

func reverse(str string) string {
	strs := strings.Split(str, " ")
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}
	return strings.Join(strs, " ")
}

func main() {
	str := "dog car snow"

	fmt.Println(reverse(str))
}
