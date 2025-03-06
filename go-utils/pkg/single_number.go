package main

import "sort"

//	a := []int{2, 2, 1}
//	fmt.Println(singleNumber(a))
//
// output: 2
func singleNumber(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	sort.Ints(nums)

	for i := 1; i < n-1; i++ {
		if nums[i] == nums[i-1] || nums[i] == nums[i+1] {
			continue
		}
		return nums[i]
	}

	return nums[n-1]
}
