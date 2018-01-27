"""Collection of utilities for the binary tree inverstion.

Utilities for building a binary tree, printing the tree, and printing the result.
"""

def build_tree(TreeNode):
    """This a helper funciton designed to build the binary_tree
    
    Given a dependencey injected ./node/TreeNode Class definition. 
    
    Arguments:
        TreeNode {./node/TreeNode::Class} -- from node import TreeNode
    
    Returns:
        [./node/TreeNode] -- The root of the fully formed binary tree
    """
    nodeA = TreeNode('A')
    nodeB = TreeNode('B')
    nodeC = TreeNode('C')
    nodeD = TreeNode('D')
    nodeE = TreeNode('E')
    nodeF = TreeNode('F')
    nodeG = TreeNode('G')

    # nodeC.add_nodes(nodeF, nodeG)
    nodeC.add_nodes(nodeF, None)
    
    nodeB.add_nodes(nodeD, nodeE)
    nodeA.add_nodes(nodeB, nodeC)
    return nodeA

def print_tree(root):
    """This is a helper function designed to print the test binary tree.
    
    Arguments:
        root {./node/TreeNode} -- This is the root of the binary tree.
    """
    
    def get_level(root):
        """This function extracts the current level of the tree to be printed.
        
        Arguments:
            root {./node/TreeNode} -- This is the node above the level to print.
        """

        level = [None, None]

        if(root.left != None):
            level[0] = root.left.val

        if(root.right != None):
            level[1] = root.right.val

        return level

    # TODO: This is hardcoded to the binary tree it would be better if it wasn't
    print(root.val)
    print(get_level(root))
    print(get_level(root.left), get_level(root.right))

def result_print(tree_root, title: str) -> None:
    """[summary]
    
    [description]

    Arguments:
        root {./node/TreeNode} -- This is the root of the binary tree.
        title {str} -- This is the title of the pretty printed block
    """

    print()
    print(title)
    print("=======================================")
    print_tree(tree_root)
    print("=======================================")
    print()
