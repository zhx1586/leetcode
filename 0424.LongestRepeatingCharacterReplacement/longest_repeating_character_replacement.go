package longestrepeatingcharacterreplacement

func characterReplacement(s string, k int) int {
	ret, n := 0, len(s)
	lo, hi := 0, 0
	counts := make([]int, 26)
	counts[s[0]-'A'] = 1

	for hi < n {
		validate := isValiade(counts, hi-lo+1, k)
		if validate {
			ret = max(ret, hi-lo+1)
			hi++
			if hi < n {
				counts[s[hi]-'A']++
			}
		} else if lo < hi {
			counts[s[lo]-'A']--
			lo++
		} else {
			hi++
			if hi < n {
				counts[s[hi]-'A']++
			}
		}
		if hi == n {
			break
		}
	}

	return ret
}

func isValiade(counts []int, interval int, k int) bool {
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
