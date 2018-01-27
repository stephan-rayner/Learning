
class TreeNode(object):
    """Definition for a binary tree node.
    """
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

    def add_nodes(self, left, right):
        self.left = left
        self.right = right
