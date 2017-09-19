package tollstation

import (
	"math"
	"fmt"
	"sort"
)

type MySet struct {
	data []int
}

func insert(set MySet, ele int) MySet {
	if ele == 4 {
		fmt.Println(ele)
	}
	set.data = append(set.data, ele)
	sort.Ints(set.data)
	fmt.Println("insert:", ele, "data:", set.data)
	return set
}

func remove(set MySet, ele int) MySet {
	for k, v := range set.data {
		if v == ele {
			if k == len(set.data)-1 {
				set.data = set.data[:k]
			} else {
				set.data = append(set.data[:k], set.data[k+1:]...)
			}
			sort.Ints(set.data)
			fmt.Println("remove:", v, "left:", set.data)
			break
		}
	}
	return set
}

func deleteMax(set MySet) (int, MySet) {
	max := math.MinInt8
	index := 0
	for k, v := range set.data {
		if v > max {
			max = v
			index = k
		}
	}
	if index < len(set.data) {
		set.data = append(set.data[:index], set.data[index+1:]...)
	} else {
		set.data = set.data[:index]
	}
	sort.Ints(set.data)
	return max, set
}

func contains(set MySet, ele int) bool {
	for _, v := range set.data {
		if v == ele {
			return true
		}
	}
	return false
}

func isEmpty(set MySet) bool {
	return len(set.data) == 0
}

func findMax(set MySet) int {
	max := math.MinInt8
	for _, v := range set.data {
		if max < v {
			max = v
		}
	}
	return max
}

func abs(num int) int {
	if num < 0 {
		return 0 - num
	}
	return num
}

func place(x []int, set MySet, n int, left int, right int) (bool, MySet) {
	if isEmpty(set) {
		return true, set
	}
	if left > right {
		return true, set
	}

	found := false
	max := findMax(set)
	/* let x[right] = max */
	check := true
	for i := 1; i <= n; i++ {
		fmt.Println("max:", max, "i:", i, "left:", left, "right:", right, "distance:", abs(max-x[i]))
		if (i < left || i > right) && !contains(set, abs(max-x[i])) {
			check = false
			break
		}
	}
	if check {
		x[right] = max
		for i := 1; i <= n; i++ {
			if i < left || i > right {
				set = remove(set, abs(max-x[i]))
			}
		}
		found, set = place(x, set, n, left, right-1)

		if !found {
			for i := 1; i <= n; i++ {
				if i < left || i > right {
					set = insert(set, abs(max-x[i]))
				}
			}
		}
	}

	/* try set[left] = min */
	check = true
	for i := 1; i <= n; i++ {
		fmt.Println("max:", max, "i:", i, "left:", left, "right:", right, "distance:", abs(x[n]-max-x[i]))
		if (i < left || i > right) && !contains(set, abs(x[n]-max-x[i])) {
			check = false
			break
		}
	}
	if !found && check {
		x[left] = x[n] - max
		for i := 1; i <= n; i++ {
			if i < left || i > right {
				set = remove(set, abs(x[n]-max-x[i]))
			}
		}
		found, set = place(x, set, n, left+1, right)
		if !found {
			for i := 1; i <= n; i++ {
				if i < left || i > right {
					set = insert(set, abs(x[n]-max-x[i]))
				}
			}
		}
	}
	return found, set
}

func turnpike(x []int, set MySet, n int) bool {
	x[1] = 0
	x[n], set = deleteMax(set)
	x[n-1], set = deleteMax(set)
	if contains(set, x[n]-x[n-1]) {
		set = remove(set, x[n]-x[n-1])
		result, _ := place(x, set, n, 2, n-2)
		return result
	}
	return false
}
