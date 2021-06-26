package minimumgeneticmutation

func minMutation(start string, end string, bank []string) int {
	edges := make([][]int, len(bank))
	for i := 0; i < len(bank); i++ {
		edges[i] = make([]int, 0)
		for j := 0; j < len(bank); j++ {
			if isMutation(bank[i], bank[j]) {
				edges[i] = append(edges[i], j)
			}
		}
	}
	visited := make([]bool, len(bank))
	queue, next := make([]int, 0), make([]int, 0)
	for i := 0; i < len(bank); i++ {
		if isMutation(start, bank[i]) {
			queue = append(queue, i)
		}
	}
	ret := 1
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		visited[curr] = true
		if isEqual(bank[curr], end) {
			return ret
		}
		for _, e := range edges[curr] {
			if !visited[e] {
				next = append(next, e)
			}
		}
		if len(queue) == 0 {
			queue = next
			next = make([]int, 0)
			ret++
		}
	}

	return -1
}

func isMutation(from string, to string) bool {
	diff := 0
	for i := 0; i < len(from); i++ {
		if from[i] != to[i] {
			diff++
		}
	}
	return diff == 1
}

func isEqual(a, b string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
