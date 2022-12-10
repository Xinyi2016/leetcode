package countsmaller

import (
	"sort"
	"testing"

	"github.com/xinyi2016/leetcode-go/structures"
)

type countSmallerCase struct {
	nums, counts []int
}

func countSmaller(nums []int) []int {
	allNums, res := make([]int, len(nums)), []int{}

	copy(allNums, nums)
	sort.Ints(allNums)
	k := 1
	kth := map[int]int{allNums[0]: k}
	for i := 1; i < len(allNums); i++ {
		if allNums[i] != allNums[i-1] {
			k++
			kth[allNums[i]] = k
		}
	}
	bit := structures.BinaryIndexedTree{}
	bit.Init(k)
	for i := len(nums) - 1; i >= 0; i-- {
		res = append(res, bit.Query(kth[nums[i]]-1))
		bit.Add(kth[nums[i]], 1)
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func TestCountSmaller(t *testing.T) {
	qs := []countSmallerCase{
		{
			nums:   []int{5, 2, 6, 1},
			counts: []int{2, 1, 1, 0},
		},
		{
			nums:   []int{-1},
			counts: []int{0},
		},
		{
			nums:   []int{-1, -1},
			counts: []int{0, 0},
		},
	}

	for idx, q := range qs {
		got := countSmaller(q.nums)
		for i, g := range got {
			if g != q.counts[i] {
				t.Fatalf("case %v: got %v, expected %v ", idx, got, q.counts)
			}
		}
	}
}
