/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    prev:=list1;for i:=1;i<a;i++{prev=prev.Next};current:=prev.Next;for i:=a;i<=b;i++{next:=current.Next;current.Next=nil;current=next};prev.Next=list2;tail:=list2;for tail.Next!=nil{tail=tail.Next};tail.Next=current;return list1
}
