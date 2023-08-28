package main

import "fmt"

func binarySearch(tar int, nums []int) bool {
	low := 0
	high := len(nums) - 1

	for low <= high {
		median := (low + high) / 2

		if nums[median] < tar {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(nums) || nums[low] != tar {
		return false
	}

	return true
}

func main() {
	nums := []int{1, 2, 3, 4, 6, 7, 8, 9}
	fmt.Println(binarySearch(4, nums))
}
