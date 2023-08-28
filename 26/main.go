package main

import (
	"fmt"
	"strings"
)

func containsDuplicate(str string) bool {
	container := make(map[string]bool)
	str = strings.ToLower(str)
	for _, value := range str {
		if container[string(value)] {
			return false
		} else {
			container[string(value)] = true
		}
	}
	return true
}

func main() {
	str1 := "еваЕВА"
	str2 := "йцукеФЫВАП"

	fmt.Println(containsDuplicate(str1), containsDuplicate(str2))
}
