package CutRod

import (
	"math"
)

func CutRodTopDown(p []int, n int) int {
	if n == 0 {
		return 0
	}

	value := math.MinInt16
	for i := 1; i <= n; i ++ {
		value = max(value, p[i]+CutRodTopDown(p, n-i))
	}
	return value
}

func MemCutRoadTopDown(p []int, n int, mem []int) int {
	if mem[n] > 0 {
		return mem[n]
	}

	if n == 0 {
		return 0
	}

	value := math.MinInt16
	for i := 1; i <= n; i++ {
		value = max(value, p[i]+MemCutRoadTopDown(p, n-i, mem))
	}
	mem[n] = value
	return value
}

func MemCutRoadDownTop(p []int, n int, mem []int) int {
	/**
	保证在每解一个大的问题之前，它所依赖的子问题都得到了解决
	 */
	if n == 0 {
		return 0
	}

	for i := 1; i <= n; i++ {
		value := math.MinInt16
		for j := 1; j <= i; j++ {
			value = max(value, p[j]+mem[i-j])
		}
		mem[i] = value
		printArray(mem)
	}
	return mem[n]
}
