package _94_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test94(t *testing.T) {
	cases := []struct {
		name string
		root *TreeNode
		res  []int
	}{
		{
			name: "1",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			res: []int{1, 3, 2},
		},
		{
			name: "2",
			root: &TreeNode{
				Val: 1,
			},
			res: []int{1},
		},
		{
			name: "3",
			root: nil,
			res:  []int{},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := inorderTraversal(c.root)
			assert.Equal(t, c.res, res)
		})
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	if root.Left != nil {
		res = append(res, inorderTraversal(root.Left)...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, inorderTraversal(root.Right)...)
	}

	return res
}
