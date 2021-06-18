package battleshipsinaboard

func countBattleships(board [][]byte) int {
	m, n := len(board), len(board[0])

	count, length := 0, 0

	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if board[x][y] == 'X' {
				length++
			} else if length > 1 {
				for length > 0 {
					board[x][y-length] = '.'
					length--
				}
				count++
			} else {
				length = 0
			}
			if y == n-1 {
				if length > 1 {
					for length > 0 {
						board[x][y-length+1] = '.'
						length--
					}
					count++
				} else {
					length = 0
				}
			}
		}
	}

	for y := 0; y < n; y++ {
		for x := 0; x < m; x++ {
			if board[x][y] == 'X' {
				length++
			} else if length > 1 {
				for length > 0 {
					board[x-length][y] = '.'
					length--
				}
				count++
			} else {
				length = 0
			}
			if x == m-1 {
				if length > 1 {
					for length > 0 {
						board[x-length+1][y] = '.'
						length--
					}
					count++
				} else {
					length = 0
				}
			}
		}
	}

	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if board[x][y] == 'X' {
				board[x][y] = '.'
				count++
			}
		}
	}

	return count
}
