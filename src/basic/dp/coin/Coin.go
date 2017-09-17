package coin

import "math"

func MinCoin(target int, cost []int) (int) {

	if target < 0 {
		return math.MaxInt8
	} else if target == 0 {
		cost[target] = 0
	} else if target == 1 || target == 3 || target == 5 {
		cost[target] = 1
	} else {
		cost[target] = min(MinCoin(target-5, cost),
			MinCoin(target-3, cost), MinCoin(target-1, cost)) + 1
	}

	return cost[target]
}

func min(a int, b int, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}

}
