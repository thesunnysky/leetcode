package main

/*
found
delete
insert
findmax
 */

func main() {

}

func turnpike(x []int, d []int, n int) int {

	return 0
}

func place(x *[]int, d *[]int, n int, left int, right int) {
	var dMax int
	var found bool
	if isEmpty(d) {
		return
	}

	dMax = findMax(d)

	for i := 1; i < left; i++ {
		if belongTo(CalcAbs(x[i] - dMax), d) {
		}
	}
}

func CalcAbs(a int) (ret int) {
	ret = (a ^ a>>31) - a>>31
	return
}

func myDelete(d *[]int, num int) {
	var temp []int
	for _, v := range *d {
		if num != v {
			temp = append(temp, v)
		}
	}
	d = &temp
}

func myInsert(d *[]int, num int) {
	temp := make([]int, 3)
	for _, v := range *d {
		temp = append(temp, v)
	}
	temp = append(temp, num)
	d = &temp
}

func findMax(d *[]int) int {
	temp := -1
	for _, v := range *d {
		if v > temp {
			temp = v
		}
	}
	return temp
}

func isEmpty(d *[]int) bool {
	return len(*d) == 0
}

func belongTo(num int, d *[]int) bool {
	for _, v := range *d {
		if num == v {
			return true
		}
	}
}
