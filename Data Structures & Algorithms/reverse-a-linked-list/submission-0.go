/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(curr *ListNode) *ListNode {
    result := (*ListNode)(nil)

	for curr != nil {
		temp := curr.Next
		curr.Next = result
		result = curr
		curr = temp
	}

	return result
}
