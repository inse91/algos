package same_tree_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
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
			//assert.Equal(t, tc.res, tc.first.equal(tc.second))
			assert.Equal(t, tc.res, isSameTree(tc.first, tc.second))
		})
	}
}
