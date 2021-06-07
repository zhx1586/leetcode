package frogjump

func canCross(stones []int) bool {
	if stones[1]-stones[0] != 1 {
		return false
	}
	return canCrossCore(stones, 1, 1)
}

func canCrossCore(stones []int, pos int, k int) bool {
	if pos == len(stones)-1 {
		return true
	}
	for i := pos + 1; i < len(stones); i++ {
		if k-1 > 0 && stones[i] == stones[pos]+k-1 {
			//fmt.Printf("jump k-1: %+v, %d, %d, (%d->%d)\n", stones, pos, k, stones[pos], stones[i])
			if canCrossCore(stones, i, k-1) {
				return true
			}
		}
		if stones[i] == stones[pos]+k {
			//fmt.Printf("jump k: %+v, %d, %d, (%d->%d)\n", stones, pos, k, stones[pos], stones[i])
			if canCrossCore(stones, i, k) {
				return true
			}
		}
		if stones[i] == stones[pos]+k+1 {
			//fmt.Printf("jump k+1: %+v, %d, %d, (%d->%d)\n", stones, pos, k, stones[pos], stones[i])
			if canCrossCore(stones, i, k+1) {
				return true
			}
		}
		if stones[i] > stones[pos]+k+1 {
			break
		}
	}
	return false
}
