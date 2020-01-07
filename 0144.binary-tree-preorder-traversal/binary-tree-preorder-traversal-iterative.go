package problem0144

func preorderTraversalIterative(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	if root != nil {
		stack = append(stack, root)
	}
	var curr *TreeNode
	for len(stack) != 0 {
		curr = stack[len(stack)-1]
		stack = stack[0:(len(stack) - 1)]
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
		res = append(res, curr.Val)
	}
	return res
}
