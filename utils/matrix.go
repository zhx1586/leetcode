package util

import (
	"sort"
	"strconv"
	"strings"
)

func Matrix(input string) [][]int {
	input = remBrackets(input)
	sliceStrs := strings.Split(input, "],[")

	if len(sliceStrs[0]) == 0 {
		return [][]int{}
	}

	matrix := make([][]int, len(sliceStrs))
	for i, sliceStr := range sliceStrs {
		sliceStr = remBrackets(sliceStr)
		matrix[i] = toInts(sliceStr)
		if i > 0 && len(matrix[i-1]) != len(matrix[i]) {
			panic("rows must have same length")
		}
	}

	return matrix
}

func remBrackets(s string) string {
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	return s
}

func toInts(s string) []int {
	if len(s) == 0 {
		return []int{}
	}
	nums := strings.Split(s, ",")
	ints := make([]int, len(nums))
	for i, num := range nums {
		n, _ := strconv.Atoi(num)
		ints[i] = n
	}
	return ints
}

type SortMatrix [][]int

func (m SortMatrix) Len() int { return len(m) }
func (m SortMatrix) Less(i, j int) bool {
	if m[i][0] < m[j][0] {
		return true
	} else if m[i][0] > m[j][0] {
		return false
	} else {
		return m[i][1] < m[j][1]
	}
}
func (m SortMatrix) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func MatrixEqual(a, b SortMatrix) bool {
	sort.Sort(a)
	sort.Sort(b)
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := 0; j < len(a[0]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
