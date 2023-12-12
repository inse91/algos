package _222

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestCountTreeNodes(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "nil_root",
			root:     nil,
			expected: 0,
		},
		{
			name:     "root_only",
			root:     &TreeNode{},
			expected: 1,
		},
		{
			name: "tree",
			root: &TreeNode{
				Left:  &TreeNode{},
				Right: &TreeNode{},
			},
			expected: 3,
		},
		{
			name: "tree_big",
			root: &TreeNode{
				Left: &TreeNode{
					Left: &TreeNode{
						Left: &TreeNode{
							Left:  &TreeNode{},
							Right: &TreeNode{},
						},
						Right: &TreeNode{},
					},
					Right: &TreeNode{},
				},
				Right: &TreeNode{},
			},
			expected: 9,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, countNodes(tc.root))
		})
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Right) + countNodes(root.Left)
}
