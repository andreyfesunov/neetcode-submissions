/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
    stack := []int{}
    cur := head

    for cur != nil {
        stack = append(stack, cur.Val)
        cur = cur.Next
    }

    cur = head
    for cur != nil && cur.Val == stack[len(stack)-1] {
        stack = stack[:len(stack)-1]
        cur = cur.Next
    }

    return cur == nil
}
