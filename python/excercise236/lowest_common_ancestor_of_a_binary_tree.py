# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x, left=None, right=None):
        self.val = x
        self.left = left
        self.right = right


def lowest_common_ancestor(root: 'TreeNode', p: 'TreeNode', q: 'TreeNode'):
    if root is None or root.val == p.val or root.val == q.val:
        return root
    left = lowest_common_ancestor(root.left, p, q)
    right = lowest_common_ancestor(root.right, p, q)

    if left and right:
        return root

    if left:
        return left

    return right
