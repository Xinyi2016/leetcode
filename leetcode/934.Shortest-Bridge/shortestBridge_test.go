package shortestbridge

import "testing"

type shortestBridgeCase struct {
	grid [][]int
	res  int
}

// TODO
func shortestBridge(grid [][]int) int {
	var res int
	return res

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
