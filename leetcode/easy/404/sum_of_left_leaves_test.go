package _404_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func TestSumOfLeftLeaves(t *testing.T) {
	cases := []struct {
		name string
		root *TreeNode
		res  int
	}{
		{
			name: "1",
			root: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
			},
			res: 24,
		},
		{
			name: "2",
			root: &TreeNode{Val: 1},
			res:  0,
		},
		{
			name: "3",
			root: nil,
			res:  0,
		},
		{
			name: "4",
			root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			res:  2,
		},
		{
			name: "5",
			root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			res:  0,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := sumOfLeftLeaves(c.root)
			assert.Equal(t, c.res, res)
		})
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsLeaf(n TreeNode) bool {
	return n.Left == nil && n.Right == nil
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil || IsLeaf(*root) {
		return 0
	}
	return _sumOfLeftLeaves(root.Right, false) + _sumOfLeftLeaves(root.Left, true)
}

func _sumOfLeftLeaves(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}
	if IsLeaf(*root) {
		if !isLeft {
			return 0
		}
		return root.Val
	}
	return _sumOfLeftLeaves(root.Right, false) + _sumOfLeftLeaves(root.Left, true)
}
