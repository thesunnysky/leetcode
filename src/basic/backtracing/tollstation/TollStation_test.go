package tollstation

import (
	"testing"
	"fmt"
)

func TestSet(t *testing.T) {
	s := []int{1, 2, 2, 2, 3, 3, 3, 4, 5, 5, 5, 6, 7, 8, 10}
	set := MySet{data: s}
	x := make([]int, 7)
	find := turnpike(x, set, 6)
	fmt.Println("find:", find)
	fmt.Println("result:", x)
}
