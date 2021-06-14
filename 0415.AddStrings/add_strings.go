package addstrings

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	ret := make([]byte, len(num1)+1)
	sum, plus := byte(0), byte(0)
	for i := 0; i < len(num1); i++ {
		if i < len(num2) {
			sum = (num1[len(num1)-1-i] - '0') + (num2[len(num2)-1-i] - '0') + plus
		} else {
			sum = (num1[len(num1)-1-i] - '0') + plus
		}
		if sum >= 10 {
			sum = sum % 10
			plus = 1
		} else {
			plus = 0
		}
		ret[len(ret)-1-i] = '0' + sum
	}
	if plus == 1 {
		ret[0] = '1'
	} else {
		ret = ret[1:]
	}

	return string(ret)
}
