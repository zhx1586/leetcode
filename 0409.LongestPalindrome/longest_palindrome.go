package longestpalindrome

func longestPalindrome(s string) int {
	counts := map[rune]int{}
	for _, letter := range s {
		counts[letter]++
	}

	ret := 0
	odd := false
	for _, v := range counts {
		if v%2 == 1 {
			ret = ret + v - 1
			odd = true
		} else {
			ret = ret + v
		}
	}
	if odd {
		ret = ret + 1
	}

	return ret
}
