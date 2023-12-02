package test_108

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

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	pivot := len(nums) / 2

	return &TreeNode{
		Val:   nums[pivot],
		Left:  sortedArrayToBST(nums[:pivot:pivot]),
		Right: sortedArrayToBST(nums[pivot+1:]),
	}
}

func Test108(t *testing.T) {
	testCases := []struct {
		name           string
		slice          []int
		expectedResult *TreeNode
	}{
		{
			name:           "nil_slice",
			expectedResult: nil,
			slice:          nil,
		},
		{
			name: "slice",
			expectedResult: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val: -10,
					},
				},
				Right: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
			slice: []int{-10, -3, 0, 5, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			node := sortedArrayToBST(tc.slice)
			assert.True(t, tc.expectedResult.equal(node))
		})
	}
}

func (root *TreeNode) equal(otherRoot *TreeNode) bool {
	if root == nil || otherRoot == nil {
		return root == otherRoot
	}

	if root.Val != otherRoot.Val {
		return false
	}

	return root.Left.equal(otherRoot.Left) && root.Right.equal(otherRoot.Right)
}

func TestTreeEqual(t *testing.T) {
	testCases := []struct {
		name   string
		first  *TreeNode
		second *TreeNode
		res    bool
	}{
		{
			name: "both_nil",
			res:  true,
		},
		{
			name:  "1st_!nil",
			first: &TreeNode{},
			res:   false,
		},
		{
			name:   "2nd_!nil",
			second: &TreeNode{},
			res:    false,
		},
		{
			name:   "both_!nil",
			first:  &TreeNode{},
			second: &TreeNode{},
			res:    true,
		},
		{
			name: "both_!nil_dif_val",
			first: &TreeNode{
				Val: 15,
			},
			second: &TreeNode{
				Val: 42,
			},
			res: false,
		},
		{
			name: "both_!nil_dif_val",
			first: &TreeNode{
				Val: 36,
			},
			second: &TreeNode{
				Val: 36,
			},
			res: true,
		},
		{
			name: "both_!nil_complex",
			first: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 6,
							Left: &TreeNode{
								Val: 7,
							},
						},
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			second: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 6,
							Left: &TreeNode{
								Val: 7,
							},
						},
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			res: true,
		},
		{
			name: "both_!nil_complex_dif",
			first: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 6,
							Left: &TreeNode{
								Val: 7,
							},
						},
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			second: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 999999999,
							Left: &TreeNode{
								Val: 7,
							},
						},
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			res: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.res, tc.first.equal(tc.second))
		})
	}
}
