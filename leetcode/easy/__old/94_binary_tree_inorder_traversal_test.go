package __old

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var out []int
	if root == nil {
		return out
	}
	if root.Right == nil && root.Left == nil {
		return []int{root.Val}
	}
	sl := append([]int{root.Val}, inorderTraversal(root.Left)...)
	return append(sl, inorderTraversal(root.Right)...)
}

func Test94(t *testing.T) {
	testCases := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "nil",
			root: nil,
			want: nil,
		},
		{
			name: "one_root_elem",
			root: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			want: []int{1},
		},
		{
			name: "two_elems",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			want: []int{1, 2},
		},
		{
			name: "complex",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val:   8,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   19,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   6,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   13,
							Left:  nil,
							Right: nil,
						},
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val:   11,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   23,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: 14,
						Left: &TreeNode{
							Val:   16,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   43,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: []int{1, 2, 5, 8, 19, 4, 6, 13, 3, 7, 11, 23, 14, 16, 43},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := inorderTraversal(tc.root)
			assert.ElementsMatch(t, tc.want, res)
		})
	}
}
