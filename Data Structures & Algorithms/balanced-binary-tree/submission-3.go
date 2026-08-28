/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	return dfs(root).b
}

type Res struct {
	h int
	b bool
}

func dfs(node *TreeNode) Res {
	if node == nil {
		return Res{0, true}
	}

	left, right := dfs(node.Left), dfs(node.Right)
	b := left.b && right.b && abs(left.h - right.h) <= 1

	return Res{b: b, h: 1 + max(left.h, right.h)}
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}