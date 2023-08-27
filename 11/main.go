package main

import "fmt"

func intersection(firstSlice, secondSlice []int) (res []int) {
	hash := make(map[int]bool)
	for _, val := range firstSlice {
		hash[val] = true
	}
	for _, val := range secondSlice {
		if hash[val] {
			res = append(res, val)
		}
	}
	return
}

func main() {

	firstSlice := []int{-56, -45, -24, -77, 0, 35, 1, 56, 20}
	secondSlice := []int{-77, 35, 6, 2, 5, -5, 1, 20}

	result := intersection(firstSlice, secondSlice)

	fmt.Println(result)

}
