package shortestbridge

import "testing"

type shortestBridgeCase struct {
	grid [][]int
	res  int
}

var dir = []struct{ r, c int }{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func shortestBridge(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	var p, q int
getFirst:
	for r, row := range grid {
		for c, col := range row {
			if col == 1 {
				p, q = r, c
				break getFirst

			}
		}
	}

	isChecked := [100][100]bool{}
	isChecked[p][q] = true

	island := make([][]int, 0)
	island = append(island, []int{p, q})

	for idx := 0; idx < len(island); idx++ {
		i, j := island[idx][0], island[idx][1]
		for _, v := range dir {
			x, y := i+v.r, j+v.c
			if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 && !isChecked[x][y] {
				island = append(island, []int{x, y})
				isChecked[x][y] = true
			}
		}
	}

	bridge := island
	var res int

	for len(bridge) > 0 {
		size := len(bridge)
		for idx := 0; idx < size; idx++ {
			land := bridge[idx]
			i, j := land[0], land[1]
			for _, v := range dir {
				x, y := i+v.r, j+v.c
				if x >= 0 && x < m && y >= 0 && y < n && !isChecked[x][y] {
					if grid[x][y] == 1 {
						return res
					}

					grid[x][y] = 1
					bridge = append(bridge, []int{x, y})
					isChecked[x][y] = true
				}
			}
		}
		bridge = bridge[size:]
		res++
	}

	panic("Expected 2 islands. Found 1.")
}

func TestShortestBridge(t *testing.T) {
	qs := []shortestBridgeCase{
		{
			grid: [][]int{
				{0, 1},
				{1, 0},
			},
			res: 1,
		},
		{
			grid: [][]int{
				{0, 1, 0},
				{0, 0, 0},
				{0, 0, 1},
			},
			res: 2,
		},
		{
			grid: [][]int{
				{1, 1, 1, 1, 1},
				{1, 0, 0, 0, 1},
				{1, 0, 1, 0, 1},
				{1, 0, 0, 0, 1},
				{1, 1, 1, 1, 1},
			},
			res: 1,
		},
	}

	for _, q := range qs {
		got := shortestBridge(q.grid)
		if got != q.res {
			t.Fatalf("got %v, expected %v", got, q.res)
		}

	}
}
