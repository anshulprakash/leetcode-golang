package problem0094

func inorderTraversalRecursive(root *TreeNode) []int {
	var inorder []int
	return helper(root, inorder)
}

func helper(node *TreeNode, inorder []int) []int {
	if node == nil {
		return inorder
	}
	inorder = helper(node.Left, inorder)
	inorder = append(inorder, node.Val)
	inorder = helper(node.Right, inorder)

	return inorder
}
