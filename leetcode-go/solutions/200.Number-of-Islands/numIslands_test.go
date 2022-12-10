package numberofislands

import (
	"log"
	"testing"
)

type NumIslandsCase struct {
	grid [][]byte
	res  int
}

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func numIslands(grid [][]byte) int {
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

	var isIsland func(x, y int)
	isIsland = func(x, y int) {
		visited[x][y] = true

		for i := 0; i < 4; i++ {
			nx := x + dir[i][0]
			ny := y + dir[i][1]

			if nx >= 0 && nx < m && ny >= 0 && ny < n && !(visited[nx][ny]) && grid[nx][ny] == '1' {
				isIsland(nx, ny)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				isIsland(i, j)
				log.Println(i, j)
				res++
			}
		}
	}

	return res
}

func TestNumIslands(t *testing.T) {
	qs := []NumIslandsCase{
		{
			grid: [][]byte{
				[]byte("11110"),
				[]byte("11010"),
				[]byte("11000"),
				[]byte("00000"),
			},
			res: 1,
		},
		{
			grid: [][]byte{
				[]byte("11000"),
				[]byte("11000"),
				[]byte("00100"),
				[]byte("00011"),
			},
			res: 3,
		},
	}

	for _, q := range qs {
		got := numIslands(q.grid)
		if got != q.res {
			t.Fatalf("got %v, expected %v", got, q.res)
		}

	}
}
