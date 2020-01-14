package problem0094

func inorderTraversalIterative(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack []*TreeNode
	var curr = root
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		res = append(res, curr.Val)
		curr = curr.Right
	}
	return res
}
