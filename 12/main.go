package main

import "fmt"

func main() {
	strSlice := []string{"cat", "cat", "dog", "cat", "tree"}
	hash := make(map[string]bool)
	res := make([]string, 0)
	for _, value := range strSlice {
		if _, ok := hash[value]; !ok {
			hash[value] = true
		}
	}
	for keys := range hash {
		res = append(res, keys)
	}
	fmt.Println(res)
}
