package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(bin_search(nums, 13))
}

func bin_search(nums []int, target int) int {
	mid_idx := len(nums) / 2


	if len(nums) == 1 && nums[mid_idx] != target {
		return -1
	}

	if nums[mid_idx] == target {
		return nums[mid_idx]
	}else if target < nums[mid_idx] {
		return bin_search(nums[:mid_idx], target)
	}else if target > nums[mid_idx] {
		return bin_search(nums[mid_idx+1:], target)
	}

	return -1
}