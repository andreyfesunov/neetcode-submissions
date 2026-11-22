/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    dummy := &ListNode{Val: 0, Next: head}
    prev := dummy

    for i := 0; i < left-1; i++ {
        prev = prev.Next
    }

    sublistHead := prev.Next
    sublistTail := sublistHead
    for i := 0; i < right-left; i++ {
        sublistTail = sublistTail.Next
    }

    nextNode := sublistTail.Next
    sublistTail.Next = nil
    prev.Next = reverseList(sublistHead)
    sublistHead.Next = nextNode

    return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head

    for curr != nil {
        temp := curr.Next
        curr.Next = prev
        prev = curr
        curr = temp
    }
    return prev
}