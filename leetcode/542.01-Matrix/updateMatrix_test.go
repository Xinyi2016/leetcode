package matrix

import "testing"

type updateMatrixCase struct {
	mat, res [][]int
}

func updateMatrix(mat [][]int) [][]int {

	return mat

}

func TestUpdateMatrix(t *testing.T) {
	qs := []updateMatrixCase{
		{
			mat: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			res: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
		},
		{
			mat: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
			res: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 2, 1},
			},
		},
	}

	for qi, q := range qs {
		got := updateMatrix(q.mat)
		for i, r := range got {
			if len(r) != len(q.res[i]) {
				t.Fatalf("case %v: got %v, expected %v", qi, got, q.res)
			}
			for pi, px := range r {
				if px != q.res[i][pi] {
					t.Fatalf("case %v: got %v, expected %v", qi, got, q.res)
				}
			}
		}

	}

}
