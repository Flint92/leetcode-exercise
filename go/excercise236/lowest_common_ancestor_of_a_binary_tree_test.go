package excercise236

import (
	"log"
	"testing"
)

func TestLowestCommonAncestor1(t *testing.T) {
	p := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 8,
		},
	}

	root := &TreeNode{
		Val:   3,
		Left:  p,
		Right: q,
	}

	ancestor := lowestCommonAncestor(root, p, q)
	if ancestor == nil {
		log.Fatal("ancestor should not be nil")
	}

	if ancestor.Val != 3 {
		log.Fatal("ancestor val should be 3")
	}
}

func TestLowestCommonAncestor2(t *testing.T) {
	p := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	q := &TreeNode{
		Val: 4,
	}

	root := &TreeNode{
		Val:  3,
		Left: p,
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	ancestor := lowestCommonAncestor(root, p, q)
	if ancestor == nil {
		log.Fatal("ancestor should not be nil")
	}

	if ancestor.Val != 5 {
		log.Fatal("ancestor val should be 5")
	}
}
