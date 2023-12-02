package test_104

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/maximum-depth-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return maxx(root.Right, root.Left)
}

func maxx(a, b *TreeNode) int {
	if a != nil && b != nil {
		am := maxx(a.Left, a.Right)
		bm := maxx(b.Left, b.Right)
		if am > bm {
			return am + 1
		}
		return bm + 1
	}
	if a != nil {
		return maxx(a.Left, a.Right) + 1
	}
	if b != nil {
		return maxx(b.Left, b.Right) + 1
	}
	return 1
}

func Test104(t *testing.T) {
	testCases := []struct {
		name           string
		root           *TreeNode
		expectedResult int
	}{
		{
			name:           "nil_root",
			root:           nil,
			expectedResult: 0,
		},
		{
			name:           "root_only",
			root:           &TreeNode{},
			expectedResult: 1,
		},
		{
			name: "tree",
			root: &TreeNode{
				Left:  &TreeNode{},
				Right: &TreeNode{},
			},
			expectedResult: 2,
		},
		{
			name: "tree_big",
			root: &TreeNode{
				Left: &TreeNode{
					Left: &TreeNode{
						Left: &TreeNode{
							Left: &TreeNode{},
						},
					},
				},
				Right: &TreeNode{
					Left: &TreeNode{
						Right: &TreeNode{},
					},
				},
			},
			expectedResult: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expectedResult, maxDepth(tc.root))
		})
	}
}
