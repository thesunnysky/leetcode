package lis

/**
 这种是贪心算法的思想，不能得到问题的解
 */
func lengthOfLIS0(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	maxLen := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			maxLen = maxLen + 1
		}
		if maxLen == 1 {
			max = nums[i]
		}
	}
	return maxLen
}

/**
感觉如何定义子问题很关键
 */
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mem := make([]int, len(nums))
	mem[0] = 1
	max := 1
	for i := 1; i < len(nums); i ++ {
		mem[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && mem[j] > mem[i]-1 {
				mem[i] = mem[j] + 1
				if mem[i] > max {
					max = mem[i]
				}
			}
		}
	}
	return max
}
