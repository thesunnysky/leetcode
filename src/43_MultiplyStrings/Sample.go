package MultiplyStrings

func multiply3(num1 string, num2 string) string {
	sum := make([]byte, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		var carry, bit byte
		for j := len(num2) - 1; j >= 0; j-- {
			bit = byte(num1[i]-'0')*byte(num2[j]-'0') + carry
			if sum[i+j+1] != 0 {
				bit += byte(sum[i+j+1] - '0')
			}
			sum[i+j+1] = bit%10 + '0'
			carry = bit / 10
		}
		/**
		 每次一个数字遍历完了之后,如果此时有进位,就将进位放到sum[i](outer循环的最新)
		 */
		if carry != 0 {
			sum[i] = carry + '0'
		}
	}

	for i := 0; i < len(sum); i++ {
		if sum[i] != 0 && sum[i]-'0' > 0 {
			return string(sum[i:])
		}
	}
	return "0"
}
