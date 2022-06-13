package countsubislands

import (
	"testing"
)

type countSubIslandsCase struct {
	grid1, grid2 [][]int
	res          int
}

// TODO
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	var res int

	return res
}

func TestFloodFill(t *testing.T) {
	qs := []countSubIslandsCase{
		{
			grid1: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
			grid2: [][]int{
				{2, 2, 2},
				{2, 2, 2},
			},
			res: 3,
		},
		{
			grid1: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
			grid2: [][]int{
				{2, 2, 2},
				{2, 2, 2},
			},
			res: 3,
		},
	}

	for _, q := range qs {
		got := countSubIslands(q.grid1, q.grid2)
		if got != q.res {
			t.Fatalf("got %v, expected %v", got, q.res)
		}

	}
}
