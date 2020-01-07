package problem0144

func preorderTraversalRecursive(root *TreeNode) []int {
	var preOrder []int
	return helper(root, preOrder)
}

func helper(node *TreeNode, preOrder []int) []int {
	if node == nil {
		return preOrder
	}
	preOrder = append(preOrder, node.Val)
	preOrder = helper(node.Left, preOrder)
	preOrder = helper(node.Right, preOrder)
	return preOrder
}
