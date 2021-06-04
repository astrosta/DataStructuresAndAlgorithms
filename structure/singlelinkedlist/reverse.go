package singlelinkedlist

/*
	单链表反转
*/
func (list *LinkedList) Reverse() {
	if list.head == nil || list.head.next == nil || list.head.next.next == nil {
		return
	}

	var pre *ListNode
	var cur *ListNode
	var next *ListNode

	cur = list.head.next

	for cur != nil {
		next = cur.next
		cur.next = pre
		pre = cur
		cur = next
	}

	list.head.next = pre
}
