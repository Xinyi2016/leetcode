package shortestpathinagridwithobstacleselimination

type pos struct {
	x, y, obstacle, step int
}

var dir = []struct{ x, y int }{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func shortestPath(grid [][]int, k int) int {
	queue, m, n := []pos{}, len(grid), len(grid[0])
	visitor := make([][][]int, m)
	if m == 1 && n == 1 {
		return 0
	}

	for i := 0; i < m; i++ {
		visitor[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			visitor[i][j] = make([]int, k+1)
		}
	}
	visitor[0][0][0] = 1
	queue = append(queue, pos{})
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			for _, d := range dir {
				newX := node.x + d.x
				newY := node.y + d.y
				if newX == m-1 && newY == n-1 {
					if node.obstacle != 0 {
						if node.obstacle <= k {
							return node.step + 1
						} else {
							continue
						}
					}
					return node.step + 1
				}
				if isInBoard(grid, newX, newY) {
					if grid[newX][newY] == 1 {
						if node.obstacle+1 <= k && visitor[newX][newY][node.obstacle+1] != 1 {
							queue = append(queue, pos{x: newX, y: newY, obstacle: node.obstacle + 1, step: node.step + 1})
							visitor[newX][newY][node.obstacle+1] = 1
						}
					} else {
						if node.obstacle <= k && visitor[newX][newY][node.obstacle] != 1 {
							queue = append(queue, pos{x: newX, y: newY, obstacle: node.obstacle, step: node.step + 1})
							visitor[newX][newY][node.obstacle] = 1
						}
					}

				}
			}

		}
	}

	return -1

}

func isInBoard(board [][]int, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}
