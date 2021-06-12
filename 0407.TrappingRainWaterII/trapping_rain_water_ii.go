package trappingrainwaterii

func trapRainWater(heightMap [][]int) int {
	//fmt.Println(heightMap)
	m, n := len(heightMap), len(heightMap[0])

	ret := 0
	height := 1
	for {
		marks := make([][]int, m)
		for i := 0; i < m; i++ {
			marks[i] = make([]int, n)
		}

		for x := 0; x < m; x++ {
			for y := 0; y < n; y++ {
				if heightMap[x][y] >= height {
					marks[x][y] = 2
				}
			}
		}
		height++
		for i := 0; i < m; i++ {
			marks = search(i, 0, marks)
			marks = search(i, n-1, marks)
		}
		for i := 1; i < n; i++ {
			marks = search(0, i, marks)
			marks = search(m-1, i, marks)
		}

		//fmt.Printf("height: %d,  marks: %+v\n", height, marks)

		inc := 0
		for x := 0; x < m; x++ {
			for y := 0; y < n; y++ {
				if marks[x][y] == 0 {
					inc++
				}
			}
		}
		ret = ret + inc
		if inc == 0 && ret != 0 {
			break
		}
		//fmt.Println(ret)
	}

	return ret
}

func search(x0 int, y0 int, marks [][]int) [][]int {
	if marks[x0][y0] > 0 {
		return marks
	}

	marks[x0][y0] = 1
	next := [][]int{{x0, y0}}

	for len(next) > 0 {
		curr := next[0]
		next = next[1:]

		x, y := curr[0]-1, curr[1]
		if canBeMarked(x, y, marks) {
			next = append(next, []int{x, y})
			marks[x][y] = 1
			//fmt.Printf("(%d,%d)->(%d,%d)\n", curr[0], curr[1], x, y)
		}
		x, y = curr[0]+1, curr[1]
		if canBeMarked(x, y, marks) {
			next = append(next, []int{x, y})
			marks[x][y] = 1
			//fmt.Printf("(%d,%d)->(%d,%d)\n", curr[0], curr[1], x, y)
		}
		x, y = curr[0], curr[1]-1
		if canBeMarked(x, y, marks) {
			next = append(next, []int{x, y})
			marks[x][y] = 1
			//fmt.Printf("(%d,%d)->(%d,%d)\n", curr[0], curr[1], x, y)
		}
		x, y = curr[0], curr[1]+1
		if canBeMarked(x, y, marks) {
			next = append(next, []int{x, y})
			marks[x][y] = 1
			//fmt.Printf("(%d,%d)->(%d,%d)\n", curr[0], curr[1], x, y)
		}
	}

	return marks
}

func canBeMarked(x int, y int, marks [][]int) bool {
	if x < 0 || x >= len(marks) {
		return false
	}
	if y < 0 || y >= len(marks[0]) {
		return false
	}
	if marks[x][y] > 0 {
		return false
	}
	return true
}
