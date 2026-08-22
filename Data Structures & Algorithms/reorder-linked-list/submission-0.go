/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    list := getNodesList(head)
	l, r := 0, len(list) / 2
	if len(list) % 2 == 0 {
		r--
	}
	curr := head 
	for l < r {
		t := curr.Next
		curr.Next = list[len(list) - 1]
		list[len(list) - 1].Next = t
		list[len(list) - 2].Next = nil
		list = list[:len(list) - 1]
		curr = t
		l++
	}
}

func getNodesList(head *ListNode) []*ListNode {
	list := make([]*ListNode, 0)
	curr := head

	for curr != nil {
		list = append(list, curr)
		curr = curr.Next
	}

	return list
}
