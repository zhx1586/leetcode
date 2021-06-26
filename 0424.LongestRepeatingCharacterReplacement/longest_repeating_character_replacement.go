package longestrepeatingcharacterreplacement

func characterReplacement(s string, k int) int {
	ret, n := 0, len(s)
	lo, hi := 0, 0
	counts := map[byte]int{s[0]: 1}
	//fmt.Println(s)

	for hi < n {
		//fmt.Printf("lo:%d, hi:%d, ret:%d, counts:%v\n", lo, hi, ret, counts)
		validate := isValiade(counts, hi-lo+1, k)
		if validate {
			ret = max(ret, hi-lo+1)
			hi++
			if hi < n {
				counts[s[hi]]++
			}
		} else if lo < hi {
			counts[s[lo]]--
			lo++
		} else {
			hi++
			if hi < n {
				counts[s[hi]]++
			}
		}
		if hi == n {
			break
		}
	}

	return ret
}

func isValiade(counts map[byte]int, interval int, k int) bool {
	for _, count := range counts {
		if count+k >= interval {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
