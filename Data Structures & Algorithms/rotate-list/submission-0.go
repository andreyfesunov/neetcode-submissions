/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }

    arr := []int{}
    cur := head
    for cur != nil {
        arr = append(arr, cur.Val)
        cur = cur.Next
    }

    n := len(arr)
    k %= n
    cur = head
    for i := n - k; i < n; i++ {
        cur.Val = arr[i]
        cur = cur.Next
    }

    for i := 0; i < n - k; i++ {
        cur.Val = arr[i]
        cur = cur.Next
    }

    return head
}