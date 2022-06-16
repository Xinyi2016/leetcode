package asfarfromlandaspossible

import "testing"

type maxDistanceCase struct {
	grid [][]int
	res  int
}

var dir = []struct{ r, c int }{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func maxDistance(grid [][]int) int {
	res := -1
	q := make([][]int, 0)
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				q = append(q, []int{r, c})
			}
		}
	}

	if len(q) == len(grid)*len(grid[0]) {
		return -1
	}

	for len(q) > 0 {

		size := len(q)
		for s := 0; s < size; s++ {
			r, c := q[s][0], q[s][1]
			for _, v := range dir {
				xr, yc := r+v.r, c+v.c
				if xr >= 0 && xr < len(grid) && yc >= 0 && yc < len(grid[0]) && grid[xr][yc] == 0 {
					q = append(q, []int{xr, yc})
					grid[xr][yc] = 1

				}
			}

		}

		q = q[size:]
		res++
	}

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
