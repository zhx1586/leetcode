package removekdigits

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}

	ret := make([]byte, len(num)-k)
	removeKdigitsCore(num, k, 0, &ret, 0)
	return leftTrim(string(ret))
}

func removeKdigitsCore(num string, k int, start int, ret *[]byte, pos int) {
	if k == 0 {
		for i := 0; i < len(num)-start; i++ {
			(*ret)[pos+i] = num[start+i]
		}
		return
	}
	if pos == len(*ret) {
		return
	}

	minValue, minValueIndex := findMinValueAndIndex(num, start, start+k+1)
	(*ret)[pos] = minValue
	removeKdigitsCore(num, k-(minValueIndex-start), minValueIndex+1, ret, pos+1)

	return
}

func findMinValueAndIndex(num string, lo int, hi int) (minValue byte, minValueIndex int) {
	minValue = num[lo]
	minValueIndex = lo

	for i := lo + 1; i < hi; i++ {
		if num[i] < minValue {
			minValue = num[i]
			minValueIndex = i
		}
	}

	return minValue, minValueIndex
}

func leftTrim(num string) string {
	count := 0
	for i := 0; i < len(num)-1; i++ {
		if num[i] == '0' {
			count++
		} else {
			break
		}
	}
	if count == 0 {
		return num
	}
	ret := make([]byte, len(num)-count)
	for i := count; i < len(num); i++ {
		ret[i-count] = num[i]
	}
	return string(ret)
}
