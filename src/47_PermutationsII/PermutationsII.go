package _7_PermutationsII

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}
	dataMap := genMap(nums)
	res = genPermute(dataMap)
	return res
}

func updateMap(dataMap map[int]int, key int) {
	if v, ok := dataMap[key]; ok && v > 0 {
		dataMap[key]--
	} else {
		delete(dataMap, key)
	}
}

func recoverMap(dataMap map[int]int, key int) {
	dataMap[key]++
}

func genPermute(dataMap map[int]int) [][]int {
	//fmt.Println("dataMap:", dataMap)
	res := make([][]int, 0)
	keySet := genKeySet(dataMap)

	for _, v := range keySet {
		updateMap(dataMap, v)
		permute := genPermute(dataMap)
		res = append(res, appendRes(v, permute)...)
		recoverMap(dataMap, v)
	}
	return res
}

func appendRes(num int, res [][]int) [][]int {
	if len(res) == 0 {
		result := []int{num}
		res = append(res, result)
		return res
	}
	for i, _ := range res {
		res[i] = append(res[i], num)
	}
	return res
}

func genMap(nums []int) map[int]int {
	dataMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := dataMap[v]; ok {
			dataMap[v]++
		} else {
			dataMap[v] = 1
		}
	}
	return dataMap
}

func genKeySet(data map[int]int) []int {
	res := make([]int, 0)
	for k, v := range data {
		if v > 0 {
			res = append(res, k)
		}
	}
	return res
}
