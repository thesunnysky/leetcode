package lds

import (
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	mem := make([]int, len(nums))
	sort.Ints(nums)
	mem[0] = 1
	sub := 0
	for i := 1; i < len(nums); i ++ {
		mem[i] = 1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && mem[j]+1 > mem[i] {
				mem[i] = mem[j] + 1
				if mem[i] > mem[sub] {
					sub = i
				}
			}
		}
	}
	temp := nums[sub]
	curMem := mem[sub]
	res := make([]int, 0)
	for i := sub; i >= 0; i-- {
		if temp%nums[i] == 0 && mem[i] == curMem {
			res = append(res, nums[i])
			temp = nums[i]
			curMem--
		}
	}
	return res
}
