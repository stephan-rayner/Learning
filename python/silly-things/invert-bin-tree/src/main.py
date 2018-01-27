from node import TreeNode
from utils import *

def reverse_tree(node):
    """This does what it says on the tin. This function takes a node in a binary tree and L-R reverses the tree. 
    
    Given a node this function swaps the children below it before
    recursively moving on to do the same for them.
    
    Arguments:
        node {./node/TreeNode} -- This is the node who's child nods we swap.
    """
    if node == None:
        return
    temp = node.left
    node.left = node.right
    node.right = temp
    reverse_tree(node.left)
    reverse_tree(node.right)

def is_symetric(node):

    if node.left == None and node.right == None:
        return True

    if node.left == None or node.right == None:
        return False

    return is_symetric(node.left) and is_symetric(node.right)

def main():
    nodeA = build_tree(TreeNode)
    result_print(nodeA, "Before")
    reverse_tree(nodeA)
    result_print(nodeA, "After")

if __name__ == '__main__':
    main()