/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeNodes(head *ListNode) *ListNode {
    dummy := &ListNode{Val: 0}
    tail, current, total := dummy, head.Next, 0
    for current != nil {
        if current.Val == 0 {
            tail.Next = &ListNode{Val: total}
            tail = tail.Next
            total = 0
        } else {
            total += current.Val
        }
        current = current.Next
    }
    return dummy.Next
}
