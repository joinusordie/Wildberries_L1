package main

import (
	"fmt"
)

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[0]
	var low, high []int

	for _, num := range nums[1:] {
		if num <= pivot {
			low = append(low, num)
		} else {
			high = append(high, num)
		}
	}

	res := append(quickSort(low), pivot)
	res = append(res, quickSort(high)...)

	return res
}

func main() {

	nums := []int{-4, -9, 2, 4, 1, 0, 2, -5, 8, -3, 10, 9}

	nums = quickSort(nums)
	fmt.Println(nums)

}
