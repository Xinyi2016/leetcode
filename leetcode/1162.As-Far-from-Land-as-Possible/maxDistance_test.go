package asfarfromlandaspossible

import "testing"

type maxDistanceCase struct {
	grid [][]int
	res  int
}

func maxDistance(grid [][]int) int {
	var res int

	return res

}

func TestMaxDistance(t *testing.T) {
	qs := []maxDistanceCase{
		{
			grid: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
			res: 2,
		},
		{
			grid: [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			res: 4,
		},
	}

	for _, q := range qs {
		got := maxDistance(q.grid)
		if got != q.res {
			t.Fatalf("got %v, expected %v", got, q.res)
		}
	}
}
