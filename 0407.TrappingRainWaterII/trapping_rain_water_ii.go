package trappingrainwaterii

import (
	"container/heap"
)

type direction struct {
	Dx int
	Dy int
}

type block struct {
	H int
	X int
	Y int
}

type blockHeap []block

func (h blockHeap) Len() int            { return len(h) }
func (h blockHeap) Less(i, j int) bool  { return h[i].H < h[j].H }
func (h blockHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *blockHeap) Push(x interface{}) { *h = append(*h, x.(block)) }
func (h *blockHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func trapRainWater(heightMap [][]int) int {
	m, n := len(heightMap), len(heightMap[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	h := make(blockHeap, 0)
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if x == 0 || x == m-1 || y == 0 || y == n-1 {
				h = append(h, block{H: heightMap[x][y], X: x, Y: y})
				visited[x][y] = true
			}
		}
	}
	heap.Init(&h)

	directions := []direction{
		{Dx: -1, Dy: 0},
		{Dx: 1, Dy: 0},
		{Dx: 0, Dy: -1},
		{Dx: 0, Dy: 1},
	}

	ret, height := 0, 0

	for len(h) > 0 {
		curr := heap.Pop(&h).(block)
		if curr.H > height {
			height = curr.H
		}
		for _, d := range directions {
			x, y := curr.X+d.Dx, curr.Y+d.Dy
			if x < 0 || x >= m || y < 0 || y >= n || visited[x][y] {
				continue
			}
			if heightMap[x][y] < height {
				ret = ret + (height - heightMap[x][y])
			}
			visited[x][y] = true
			heap.Push(&h, block{H: heightMap[x][y], X: x, Y: y})
		}
	}

	return ret
}
