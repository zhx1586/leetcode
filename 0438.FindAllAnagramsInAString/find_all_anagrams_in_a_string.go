package findallanagramsinastring

type counter [26]int

func (c *counter) Inc(letter byte) {
	c[int(letter-'a')]++
}

func (c *counter) Dec(letter byte) {
	c[int(letter-'a')]--
}

func (c counter) In(another counter) bool {
	for i := 0; i < len(c); i++ {
		if c[i] > another[i] {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	ret := []int{}
	if len(s) == 0 {
		return ret
	}

	pattern, window := counter{}, counter{}
	for i := 0; i < len(p); i++ {
		pattern.Inc(p[i])
	}
	window.Inc(s[0])

	lo, hi := 0, 0
	for {
		if window.In(pattern) {
			if hi-lo+1 == len(p) {
				ret = append(ret, lo)
				window.Dec(s[lo])
				lo++
			} else {
				hi++
				if hi < len(s) {
					window.Inc(s[hi])
				} else {
					break
				}
			}
		} else {
			if lo < hi {
				window.Dec(s[lo])
				lo++
			} else {
				window.Dec(s[lo])
				lo++
				hi++
				if hi < len(s) {
					window.Inc(s[hi])
				} else {
					break
				}
			}
		}
	}

	return ret
}
