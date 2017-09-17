package MultiplyStrings

import (
	"strconv"
	"bytes"
)

func multiply2(num1 string, num2 string) string {
	product := make([]int, len(num1)+len(num2)+1)

	for k := len(num1) - 1; k >= 0; k-- {
		for m := len(num2) - 1; m >= 0; m-- {
			numi, _ := strconv.Atoi(string(num1[k]))
			numj, _ := strconv.Atoi(string(num2[m]))

			i, j := len(num1)-k-1, len(num2)-m-1
			sum := numi*numj + product[i+j]
			product[i+j] = sum % 10
			sum2 := sum/10 + product[i+j+1]
			product[i+j+1] = sum2 % 10
			product[i+j+2] = sum2/10 + product[i+j+2]
		}
	}
	var buf bytes.Buffer
	from := 0
	for i := len(product) - 1; i >= 0; i-- {
		if product[i] != 0 {
			from = i
			break
		}
	}
	for ; from >= 0; from -- {
		buf.WriteString(strconv.Itoa(product[from]))
	}
	return buf.String()
}
