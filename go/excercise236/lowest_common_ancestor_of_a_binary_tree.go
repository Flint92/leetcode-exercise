package excercise236

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		// 左子树为空，肯定都在右子树上
		return right
	}

	if right == nil {
		// 右子树为空，肯定都在左子树上
		return left
	}

	// 左右子树都不为空，一个在左子树一个在右子树，所以root就是他们的最近公共祖先节点。
	return root
}
