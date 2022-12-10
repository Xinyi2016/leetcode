package binarytreetraversal

import (
	"log"
	"reflect"
	"testing"

	"github.com/xinyi2016/leetcode-go/structures"
)

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	res = append(res, []int{root.Val})
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		tmp := make([]*TreeNode, 0)
		levelRes := []int{}
		for _, q := range queue {
			if q != nil {
				levelRes = append(levelRes, q.Val)
				tmp = append(tmp, []*TreeNode{q.Left, q.Right}...)
			}
		}
		if len(levelRes) > 0 {
			res = append(res, levelRes)
		}

		queue = tmp
	}

	return res
}

func TestLevelorder(t *testing.T) {
	qs := []BFSCase{
		{
			input: []int{3, 9, 20, structures.NULL, structures.NULL, 15, 7},
			expected: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
		{
			input:    []int{},
			expected: [][]int{},
		},
		{
			input: []int{1},
			expected: [][]int{
				{1},
			},
		},
	}

	for i, q := range qs {
		root := structures.NewTreeNode(q.input)
		if res := levelOrder(root); !reflect.DeepEqual(res, q.expected) {
			log.Fatalf("[ERROR] %d: got %v, expected %v", i, res, q.expected)
		}
	}

}
