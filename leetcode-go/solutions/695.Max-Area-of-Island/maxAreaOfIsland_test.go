package maxareaofisland

import (
	"testing"
)

type MaxAreaOfIslandCase struct {
	grid [][]int
	res  int
}

func maxAreaOfIsland(grid [][]int) int {
	visited := make([][]bool, len(grid))
	n := len(grid[0])
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, n)
	}

	var area func(r, c int) int

	area = func(r, c int) int {
		if !(r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && !visited[r][c] && grid[r][c] == 1) {
			return 0
		}

		visited[r][c] = true
		return (1 + area(r+1, c) + area(r-1, c) + area(r, c+1) + area(r, c-1))

	}

	var res int
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			rcArea := area(r, c)
			if rcArea > res {
				res = rcArea
			}
		}
	}

	return res

}
func TestMaxAreaOfIsland(t *testing.T) {
	qs := []MaxAreaOfIslandCase{
		{
			grid: [][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			res: 6,
		},
		{
			grid: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			res: 0,
		},
	}

	for _, q := range qs {
		got := maxAreaOfIsland(q.grid)
		if got != q.res {
			t.Fatalf("got %v, expected %v", got, q.res)
		}

	}
}
