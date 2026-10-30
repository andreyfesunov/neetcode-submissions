/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    for headA != nil {
        cur := headB
        for cur != nil {
            if headA == cur {
                return headA
            }
            cur = cur.Next
        }
        headA = headA.Next
    }
    return nil
}