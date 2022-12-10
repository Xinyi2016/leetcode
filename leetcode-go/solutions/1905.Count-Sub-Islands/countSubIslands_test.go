package countsubislands

import (
	"testing"
)

type countSubIslandsCase struct {
	grid1, grid2 [][]int
	res          int
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	var res int

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= len(grid1) || c < 0 || c >= len(grid1[0]) || grid2[r][c] == 0 {
			return
		}

		grid2[r][c] = 0
		dfs(r+1, c)
		dfs(r, c+1)
		dfs(r, c-1)
		dfs(r-1, c)

	}

	for r := 0; r < len(grid1); r++ {
		for c := 0; c < len(grid1[0]); c++ {
			if grid2[r][c] == 1 && grid1[r][c] == 0 {
				dfs(r, c)
			}
		}
	}

	for r := 0; r < len(grid1); r++ {
		for c := 0; c < len(grid1[0]); c++ {
			if grid2[r][c] == 1 {
				dfs(r, c)
				res++
			}
		}
	}

	return res
}

func TestCountSubIslands(t *testing.T) {
	qs := []countSubIslandsCase{
		{
			grid1: [][]int{
				{1, 1, 1, 0, 0},
				{0, 1, 1, 1, 1},
				{0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 1, 1},
			},
			grid2: [][]int{
				{1, 1, 1, 0, 0},
				{0, 0, 1, 1, 1},
				{0, 1, 0, 0, 0},
				{1, 0, 1, 1, 0},
				{0, 1, 0, 1, 0},
			},
			res: 3,
		},
		{
			grid1: [][]int{
				{1, 0, 1, 0, 1},
				{1, 1, 1, 1, 1},
				{0, 0, 0, 0, 0},
				{1, 1, 1, 1, 1},
				{1, 0, 1, 0, 1},
			},
			grid2: [][]int{
				{0, 0, 0, 0, 0},
				{1, 1, 1, 1, 1},
				{0, 1, 0, 1, 0},
				{0, 1, 0, 1, 0},
				{1, 0, 0, 0, 1},
			},
			res: 2,
		},
	}

	for _, q := range qs {
		got := countSubIslands(q.grid1, q.grid2)
		if got != q.res {
			t.Fatalf("got %v, expected %v", got, q.res)
		}

	}
}
