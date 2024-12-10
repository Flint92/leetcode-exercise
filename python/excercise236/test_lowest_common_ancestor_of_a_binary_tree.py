import unittest

from excercise236.lowest_common_ancestor_of_a_binary_tree import lowest_common_ancestor, TreeNode


class TestCommonAncestor(unittest.TestCase):
    def test_common_ancestor(self):
        p = TreeNode(5, left=TreeNode(6), right=TreeNode(2, left=TreeNode(7), right=TreeNode(4)))
        q = TreeNode(1, left=TreeNode(0), right=TreeNode(8))
        root = TreeNode(3, left=p, right=q)

        ancestor = lowest_common_ancestor(root, p, q)
        self.assertIsNotNone(ancestor)
        self.assertEqual(ancestor.val, 3)

        q = TreeNode(4)
        ancestor = lowest_common_ancestor(root, p, q)
        self.assertIsNotNone(ancestor)
        self.assertEqual(ancestor.val, 5)


if __name__ == '__main__':
    unittest.main()
