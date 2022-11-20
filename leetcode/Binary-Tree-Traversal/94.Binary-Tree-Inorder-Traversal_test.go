package binarytreetraversal

import (
	"log"
	"reflect"
	"testing"

	"github.com/xinyi2016/leetcode-go/structures"
)

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	inorder(root.Left, &res)
	res = append(res, root.Val)
	inorder(root.Right, &res)
	return res
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inorder(root.Left, res)
	}
	*res = append(*res, root.Val)
	if root.Right != nil {
		inorder(root.Right, res)
	}
}

func TestInorder(t *testing.T) {
	qs := []TraversalCase{
		{
			input:    []int{1, structures.NULL, 2, 3},
			expected: []int{1, 3, 2},
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
		if res := inorderTraversal(root); !reflect.DeepEqual(res, q.expected) {
			log.Fatalf("[ERROR] %d: got %v, expected %v", i, res, q.expected)
		}
	}

}
