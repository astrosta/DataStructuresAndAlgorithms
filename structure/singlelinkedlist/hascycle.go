package singlelinkedlist

/*
单链表是否有环
*/
func (list *LinkedList) HasCycle() bool {
	if list.head == nil {
		return false
	}

	fast := list.head
	slow := list.head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			return true
		}
	}

	return false
}
