package floodfill

import (
	"testing"
)

type FloodFillCase struct {
	image, res       [][]int
	sr, sc, newColor int
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	if color == newColor {
		return image
	}

	var dfs func(r, c int)

	dfs = func(r, c int) {
		if image[r][c] == color {
			image[r][c] = newColor
			if r >= 1 {
				dfs(r-1, c)
			}
			if r+1 < len(image) {
				dfs(r+1, c)
			}

			if c >= 1 {
				dfs(r, c-1)
			}

			if c+1 < len(image[0]) {
				dfs(r, c+1)
			}
		}
	}

	dfs(sr, sc)

	return image
}

func TestFloodFill(t *testing.T) {
	qs := []FloodFillCase{
		{
			image: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
			sr:       0,
			sc:       0,
			newColor: 2,
			res: [][]int{
				{2, 2, 2},
				{2, 2, 2},
			},
		},
		{
			image: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},
			sr:       1,
			sc:       1,
			newColor: 2,
			res: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
	}

	for _, q := range qs {
		got := floodFill(q.image, q.sr, q.sc, q.newColor)
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
