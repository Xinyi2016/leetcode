package climbingstairs

import (
	"testing"
)

type climbStairsCase struct {
	n, res int
}

func climbStairs(n int) int {
	if n < 4 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]

}

func TestCountSmaller(t *testing.T) {
	qs := []climbStairsCase{
		{
			n:   2,
			res: 2,
		},
		{
			n:   3,
			res: 3,
		},
		{
			n:   4,
			res: 5,
		},
	}

	for idx, q := range qs {
		if got := climbStairs(q.n); got != q.res {
			t.Fatalf("case %v: got %v, expected %v ", idx, got, q.res)
		}

	}
}
