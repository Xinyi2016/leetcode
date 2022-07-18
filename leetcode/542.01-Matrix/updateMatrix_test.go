package matrix

import (
	"math"
	"testing"
)

type updateMatrixCase struct {
	mat, res [][]int
}

func updateMatrix(mat [][]int) [][]int {
	for i, row := range mat {
		for j, val := range row {
			if val == 0 {
				continue
			}
			left, top := math.MaxInt16, math.MaxInt16
			if i > 0 {
				top = mat[i-1][j] + 1
			}
			if j > 0 {
				left = mat[i][j-1] + 1
			}
			mat[i][j] = min(top, left)
		}
	}
	for i := len(mat) - 1; i >= 0; i-- {
		for j := len(mat[0]) - 1; j >= 0; j-- {
			if mat[i][j] == 0 {
				continue
			}
			right, bottom := math.MaxInt16, math.MaxInt16
			if i < len(mat)-1 {
				bottom = mat[i+1][j] + 1
			}
			if j < len(mat[0])-1 {
				right = mat[i][j+1] + 1
			}
			mat[i][j] = min(mat[i][j], min(bottom, right))
		}
	}
	return mat
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
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
