/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type QueueNode struct {
    n *TreeNode
    l int
}

func rightSideView(root *TreeNode) []int {
if root == nil {
    return []int{}
}

    m := make(map[int]int)
    queue := make([]QueueNode, 0)

queue = append(queue, QueueNode{
    n: root, l: 0,
})

for len(queue) != 0 {
node := queue[0]
queue = queue[1:]

_, ok := m[node.l]
if !ok {
    m[node.l] = node.n.Val
}

if node.n.Left != nil {
    queue = append([]QueueNode{QueueNode{node.n.Left, node.l+1}}, queue...)
}
if node.n.Right != nil {
    queue = append([]QueueNode{QueueNode{node.n.Right, node.l+1}}, queue...)
}
}

result := make([]int, 0)
i := 0
for m[i] != 0 {
    result = append(result, m[i])
    i++
}

return result
}
