package frogjump

import "fmt"

func canCross(stones []int) bool {
	if stones[1]-stones[0] != 1 {
		return false
	}
	return canCrossCore(stones, 1, 1, map[string]bool{})
}

func canCrossCore(stones []int, pos int, k int, cache map[string]bool) bool {
	if pos == len(stones)-1 {
		return true
	}
	key := fmt.Sprintf("%d-%d", pos, k)
	if cached, ok := cache[key]; ok {
		return cached
	}
	for i := pos + 1; i < len(stones); i++ {
		if k-1 > 0 && stones[i] == stones[pos]+k-1 {
			if canCrossCore(stones, i, k-1, cache) {
				cache[key] = true
				return true
			}
		}
		if stones[i] == stones[pos]+k {
			if canCrossCore(stones, i, k, cache) {
				cache[key] = true
				return true
			}
		}
		if stones[i] == stones[pos]+k+1 {
			if canCrossCore(stones, i, k+1, cache) {
				cache[key] = true
				return true
			}
		}
		if stones[i] > stones[pos]+k+1 {
			break
		}
	}
	cache[key] = false
	return false
}
