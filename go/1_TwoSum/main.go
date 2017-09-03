package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var a []int
	for i := 0; i < len(nums)-1; i ++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				a = append(append(a, i), j)
			}
		}
	}
	return a
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	result := twoSum(nums, target)
	fmt.Println(result)
}
