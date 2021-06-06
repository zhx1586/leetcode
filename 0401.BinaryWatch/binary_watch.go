package binarywatch

import "fmt"

var maxBitsForHours = 4
var maxBitsForMinuts = 6
var cacheForHours = map[int][]int{}
var cacheForMinuts = map[int][]int{}

func readBinaryWatch(turnedOn int) []string {
	rets := []string{}
	for i := 0; i <= maxBitsForHours && i <= turnedOn; i++ {
		hours := buildWithCache(i, maxBitsForHours, cacheForHours)
		minutes := buildWithCache(turnedOn-i, maxBitsForMinuts, cacheForMinuts)
		for _, hour := range hours {
			for _, minute := range minutes {
				if ret, ok := validateAndPrint(hour, minute); ok {
					rets = append(rets, ret)
				}
			}
		}
	}

	return rets
}

func buildWithCache(count int, maxPos int, cache map[int][]int) []int {
	if cached, ok := cache[count]; ok {
		return cached
	}
	ret := []int{}
	build(count, 0, maxPos, 0, &ret)
	cache[count] = ret
	return ret
}

func build(count int, pos int, maxPos int, curr int, ret *[]int) {
	if count > (maxPos-pos)+1 {
		return
	}
	if count == 0 {
		*ret = append(*ret, curr)
		return
	}
	build(count-1, pos+1, maxPos, curr+(1<<pos), ret)
	build(count, pos+1, maxPos, curr, ret)
}

func validateAndPrint(hours int, minutes int) (string, bool) {
	if hours < 0 || hours > 11 {
		return "", false
	}
	if minutes < 0 || minutes > 59 {
		return "", false
	}
	return fmt.Sprintf("%d:%.2d", hours, minutes), true
}
