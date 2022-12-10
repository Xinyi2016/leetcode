package binarytreetraversal

import "github.com/xinyi2016/leetcode-go/structures"

type TreeNode = structures.TreeNode

type TraversalCase struct {
	input, expected []int
}

type BFSCase struct {
	input    []int
	expected [][]int
}
