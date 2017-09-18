package lds

import "sort"

func largestDivisibleSubsetSample(nums []int) []int {
	sort.Ints(nums)

	p := make([]int, len(nums))
	t := make([]int, len(nums))

	max := 0
	maxi := 0

	for i := len(nums) - 1; i >= 0; i-- {
		for j := i; j < len(nums); j++ {
			if nums[j]%nums[i] == 0 && t[i] < t[j]+1 {
				t[i] = t[j] + 1
				p[i] = j

				if t[i] > max {
					max = t[i]
					maxi = i
				}
			}
		}
	}

	answer := []int{}
	i := maxi
	for count := 0; count < max; count++ {
		answer = append(answer, nums[i])
		i = p[i]
	}

	return answer
}
