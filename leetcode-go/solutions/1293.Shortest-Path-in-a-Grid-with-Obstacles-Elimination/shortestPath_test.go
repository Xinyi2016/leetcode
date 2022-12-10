package shortestpathinagridwithobstacleselimination

import (
	"testing"
)

type shortestPathCase struct {
	grid   [][]int
	k, res int
}

func TestClosedIsland(t *testing.T) {
	qs := []shortestPathCase{
		{
			grid: [][]int{
				{0, 0, 0},
				{1, 1, 0},
				{0, 0, 0},
				{0, 1, 1},
				{0, 0, 0},
			},
			k:   1,
			res: 6,
		},
		{
			grid: [][]int{
				{0, 1, 1},
				{1, 1, 1},
				{1, 0, 0},
			},
			k:   1,
			res: -1,
		},
	}

	for qi, q := range qs {
		got := shortestPath(q.grid, q.k)
		if got != q.res {
			t.Fatalf("case %v: got %v, expected %v", qi, got, q.res)
		}

	}

}
