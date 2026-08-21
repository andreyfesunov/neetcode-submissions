/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	curr := head
	visited := make(map[*ListNode]struct{})
	for curr != nil {
		if _, ok := visited[curr]; ok {
			return true
		}
		visited[curr] = struct{}{}
		curr = curr.Next
	}
	return false
}
