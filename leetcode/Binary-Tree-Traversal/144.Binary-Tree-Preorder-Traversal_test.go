package binarytreetraversal

import (
	"log"
	"reflect"
	"testing"

	"github.com/xinyi2016/leetcode-go/structures"
)

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	preorder(root, &res)
	return res
}

func preorder(root *TreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		preorder(root.Left, res)
		preorder(root.Right, res)

	}
}

// iteratively using Stack
func preorderTraversalIter(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node != nil {
			res = append(res, node.Val)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)

		}

		if node.Left != nil {
			stack = append(stack, node.Left)

		}

	}

	return res

}

func TestPreorder(t *testing.T) {
	qs := []TraversalCase{
		{
			input:    []int{1, structures.NULL, 2, 3},
			expected: []int{1, 2, 3},
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
		if res := preorderTraversalIter(root); !reflect.DeepEqual(res, q.expected) {
			log.Fatalf("[ERROR] %d: got %v, expected %v", i, res, q.expected)
		}
	}

}
