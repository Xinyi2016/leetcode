package numberofclosedislands

import (
	"testing"
)

type ClosedIslandCase struct {
	grid [][]int
	res  int
}

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func closedIsland(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])
	if n == 0 {
		return 0
	}

	res, visited := 0, make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var isIsland func(x, y int, isEdge *bool)
	isIsland = func(x, y int, isEdge *bool) {
		visited[x][y] = true
		if x == 0 || x == len(grid)-1 || y == 0 || y == len(grid[0])-1 {
			*isEdge = true
		}

		for i := 0; i < 4; i++ {
			nx := x + dir[i][0]
			ny := y + dir[i][1]

			if nx >= 0 && nx < m && ny >= 0 && ny < n && !(visited[nx][ny]) && grid[nx][ny] == 0 {
				isIsland(nx, ny, isEdge)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			isEdge := false
			if grid[i][j] == 0 && !visited[i][j] {
				isIsland(i, j, &isEdge)
				if !isEdge {
					res++
				}

			}
		}
	}

	return res

}

func TestClosedIsland(t *testing.T) {
	qs := []ClosedIslandCase{
		{
			grid: [][]int{
				{0, 0, 1, 1, 0, 1, 0, 0, 1, 0},
				{1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
				{1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
				{0, 1, 1, 0, 0, 0, 0, 1, 0, 1},
				{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
				{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
				{1, 0, 1, 0, 1, 1, 0, 0, 0, 1},
				{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 1, 0, 1, 0, 1},
				{1, 1, 1, 0, 1, 1, 0, 1, 1, 0},
			},
			res: 5,
		},
		{
			grid: [][]int{
				{1, 1, 1, 1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 1, 0},
				{1, 0, 1, 0, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 0, 1},
				{1, 1, 1, 1, 1, 1, 1, 0},
			},
			res: 2,
		},
		{
			grid: [][]int{
				{0, 0, 1, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 1, 1, 1, 0},
			},
			res: 1,
		},
	}

	for qi, q := range qs {
		got := closedIsland(q.grid)
		if got != q.res {
			t.Fatalf("case %v: got %v, expected %v", qi, got, q.res)
		}

	}

}
