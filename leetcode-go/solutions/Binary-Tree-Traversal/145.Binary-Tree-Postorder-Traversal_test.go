package binarytreetraversal

import (
	"log"
	"reflect"
	"testing"

	"github.com/xinyi2016/leetcode-go/structures"
)

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	postorder(root, &res)
	return res
}

func postorder(root *TreeNode, res *[]int) {
	if root != nil {
		postorder(root.Left, res)
		postorder(root.Right, res)
		*res = append(*res, root.Val)

	}
}

func TestPostorder(t *testing.T) {
	qs := []TraversalCase{
		{
			input:    []int{1, structures.NULL, 2, 3},
			expected: []int{3, 2, 1},
		},
		{
			input:    []int{},
			expected: []int{},
		},
		{
			input:    []int{1},
			expected: []int{1},
		},
	}

	for i, q := range qs {
		root := structures.NewTreeNode(q.input)
		if res := postorderTraversal(root); !reflect.DeepEqual(res, q.expected) {
			log.Fatalf("[ERROR] %d: got %v, expected %v", i, res, q.expected)
		}
	}

}
