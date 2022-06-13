package numberofenclaves

import (
	"testing"
)

type NumEnclavesCase struct {
	grid [][]int
	res  int
}

var dir = []struct{ x, y int }{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func numEnclaves(grid [][]int) int {
	var res int
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}

	var isEnclave func(x, y int)
	isEnclave = func(x, y int) {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || visited[x][y] || grid[x][y] == 0 {
			return
		}

		visited[x][y] = true
		for _, v := range dir {
			isEnclave(x+v.x, y+v.y)
		}

	}

	for i := range grid {
		isEnclave(i, 0)
		isEnclave(i, len(grid[0])-1)
	}

	for i := 0; i < len(grid[0]); i++ {
		isEnclave(0, i)
		isEnclave(len(grid)-1, i)
	}

	for i, v := range grid {
		for j := range v {
			if grid[i][j] == 1 && visited[i][j] == false {
				res++
			}
		}
	}

	return res

}

func TestClosedIsland(t *testing.T) {
	qs := []NumEnclavesCase{
		{
			grid: [][]int{
				{0, 0, 0, 0},
				{1, 0, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
			},
			res: 3,
		},
		{
			grid: [][]int{
				{0, 1, 1, 0},
				{0, 0, 1, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
			res: 0,
		},
	}

	for qi, q := range qs {
		got := numEnclaves(q.grid)
		if got != q.res {
			t.Fatalf("case %v: got %v, expected %v", qi, got, q.res)
		}

	}

}
