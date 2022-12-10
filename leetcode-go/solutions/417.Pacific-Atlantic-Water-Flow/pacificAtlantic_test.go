package pacificatlanticwaterflow

import (
	"testing"
)

type pacificAtlanticCase struct {
	heights, res [][]int
}

var dir = []struct{ x, y int }{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func pacificAtlantic(heights [][]int) [][]int {
	pVisited := make([][]bool, len(heights))
	aVisited := make([][]bool, len(heights))
	for c := 0; c < len(heights); c++ {
		pVisited[c] = make([]bool, len(heights[0]))
		aVisited[c] = make([]bool, len(heights[0]))
	}

	res := make([][]int, 0)

	var toOcean func(r, c int, visited [][]bool)
	toOcean = func(r, c int, visited [][]bool) {
		visited[r][c] = true
		for _, d := range dir {
			x, y := r+d.x, c+d.y
			if x < 0 || x >= len(heights) || y < 0 || y >= len(heights[0]) || visited[x][y] || heights[x][y] < heights[r][c] {
				continue
			}

			toOcean(x, y, visited)
		}

	}

	for i := 0; i < len(heights); i++ {
		toOcean(i, 0, pVisited)
		toOcean(i, len(heights[0])-1, aVisited)
	}

	for j := 0; j < len(heights[0]); j++ {
		toOcean(0, j, pVisited)
		toOcean(len(heights)-1, j, aVisited)
	}

	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			if pVisited[i][j] && aVisited[i][j] {
				res = append(res, []int{i, j})
			}

		}
	}

	return res
}

func TestPacificAtlantic(t *testing.T) {
	qs := []pacificAtlanticCase{
		{
			heights: [][]int{
				{1, 1},
				{1, 1},
				{1, 1},
			},
			res: [][]int{
				{0, 0},
				{0, 1},
				{1, 0},
				{1, 1},
				{2, 0},
				{2, 1},
			},
		},
		{
			heights: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			res: [][]int{
				{0, 4},
				{1, 3},
				{1, 4},
				{2, 2},
				{3, 0},
				{3, 1},
				{4, 0},
			},
		},
		{
			heights: [][]int{
				{2, 1},
				{1, 2},
			},
			res: [][]int{
				{0, 0},
				{0, 1},
				{1, 0},
				{1, 1},
			},
		},
	}

	for _, q := range qs {
		got := pacificAtlantic(q.heights)
		for i, r := range got {
			if len(r) != len(q.res[i]) {
				t.Fatalf("got %v, expected %v", got, q.res)
			}
			for pi, px := range r {
				if px != q.res[i][pi] {
					t.Fatalf("got %v, expected %v", got, q.res)
				}
			}
		}

	}
}
