package pacificatlanticwaterflow

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])

	pacific, atlantic := [][]int{}, [][]int{}
	for i := 0; i < m; i++ {
		pacific = append(pacific, []int{i, 0})
		atlantic = append(atlantic, []int{i, n - 1})
	}
	for i := 1; i < n; i++ {
		pacific = append(pacific, []int{0, i})
		atlantic = append(atlantic, []int{m - 1, i - 1})
	}

	marks := [][]int{}
	for i := 0; i < m; i++ {
		marks = append(marks, make([]int, n))
	}
	marks = searchCore(heights, pacific, marks)
	marks = searchCore(heights, atlantic, marks)

	ret := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if marks[i][j] == 2 {
				ret = append(ret, []int{i, j})
			}
		}
	}

	return ret
}

func searchCore(heights [][]int, bounds [][]int, marks [][]int) [][]int {
	visited := [][]bool{}
	m, n := len(heights), len(heights[0])
	for i := 0; i < m; i++ {
		visited = append(visited, make([]bool, n))
	}

	queue := [][]int{}
	for _, b := range bounds {
		x0, y0 := b[0], b[1]
		visited[x0][y0] = true
		marks[x0][y0]++
		queue = append(queue, []int{x0, y0})
	}

	dirctions := []struct{ x, y int }{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		for _, d := range dirctions {
			x, y := q[0]+d.x, q[1]+d.y
			if x < 0 || x >= m || y < 0 || y >= n || visited[x][y] {
				continue
			}
			if heights[x][y] >= heights[q[0]][q[1]] {
				visited[x][y] = true
				marks[x][y]++
				queue = append(queue, []int{x, y})
			}
		}
	}

	return marks
}
