/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
    n := 0
    cur := head
    for cur != nil {
        cur = cur.Next
        n++
    }

    n /= 2
    cur = head
    for n > 0 {
        n--
        cur = cur.Next
    }
    return cur
}