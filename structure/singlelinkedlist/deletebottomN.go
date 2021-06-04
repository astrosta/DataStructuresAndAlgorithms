package singlelinkedlist

/*
删除倒数第N个节点
*/
func (list LinkedList) DeleteBottomN(n int) {
	fast := list.head
	slow := list.head

	for i := 0; i < n && fast != nil; i++ {
		fast = fast.next
	}

	if fast == nil {
		return
	}

	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}

	slow.next = slow.next.next

}
