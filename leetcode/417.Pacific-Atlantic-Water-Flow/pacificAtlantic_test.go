package pacificatlanticwaterflow

import (
	"testing"
)

type pacificAtlanticCase struct {
	heights, res [][]int
}

func pacificAtlantic(heights [][]int) [][]int {

	return heights
}

func TestPacificAtlantic(t *testing.T) {
	qs := []pacificAtlanticCase{
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
