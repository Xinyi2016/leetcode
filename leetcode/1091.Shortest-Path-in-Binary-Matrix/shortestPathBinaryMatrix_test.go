package shortestpathinbinarymatrix

import (
	"testing"
)

type shortestpathinbinarymatrixCase struct {
	grid [][]int
	res  int
}

var dir = []struct{ r, c int }{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, -1},
	{1, 1},
	{1, -1},
	{-1, 1},
}

func shortestPathBinaryMatrix(grid [][]int) int {

	n := len(grid)
	if grid[0][0]+grid[n-1][n-1] > 0 {
		return -1
	}

	q := [][]int{{0, 0}}

	res := 1
	for len(q) > 0 {
		size := len(q)
		for s := 0; s < size; s++ {
			r, c := q[s][0], q[s][1]
			for _, v := range dir {
				xr, yc := r+v.r, c+v.c
				if xr == n && yc == n {
					return res
				}
				if xr >= 0 && xr < len(grid) && yc >= 0 && yc < len(grid[0]) && grid[xr][yc] == 0 {
					q = append(q, []int{xr, yc})
					grid[xr][yc] = 1

				}
			}

		}

		q = q[size:]
		res++
	}

	return -1
}

func TestShortestPathBinaryMatrix(t *testing.T) {
	qs := []shortestpathinbinarymatrixCase{
		{
			grid: [][]int{
				{0, 1},
				{1, 0},
			},
			res: 2,
		},
		{
			grid: [][]int{
				{0, 0, 0},
				{1, 1, 0},
				{1, 1, 0},
			},
			res: 4,
		},
		{
			grid: [][]int{
				{1, 0, 0},
				{1, 1, 0},
				{1, 1, 0},
			},
			res: -1,
		},
	}

	for qi, q := range qs {
		got := shortestPathBinaryMatrix(q.grid)
		if got != q.res {
			t.Fatalf("case %v: got %v, expected %v", qi, got, q.res)
		}

	}

}
