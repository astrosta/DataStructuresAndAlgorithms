package singlelinkedlist

/*
目的：判断单向链表是否是回文链表
方法：使用快慢指针，当快指针到达链表尾部时，慢指针到达链表中部，将前半部分链表反转，比较后半部分和前半部分链表是否相等
*/

func IsPalindrome(list *LinkedList) bool {
	if list.head.next == nil {
		return true
	}

	slow := list.head
	fast := list.head
	var pre *ListNode
	var temp *ListNode
	for fast != nil && fast.next != nil {
		fast = fast.next.next

		temp = slow.next
		if slow != list.head {
			slow.next = pre
			pre = slow
		}
		slow = temp
	}

	temp = slow.next
	if fast != nil {
		slow.next = pre
		pre = slow
	}

	var l1 *ListNode
	var l2 *ListNode

	l1 = pre
	l2 = temp
	//
	//fmt.Println("l1:", l1.value)
	//fmt.Println("l2:", l2.value)

	for l1 != nil && l2 != nil && l1.value == l2.value {
		l1 = l1.next
		l2 = l2.next
		//fmt.Println("l1:", l1.value)
		//fmt.Println("l2:", l2.value)
	}

	return nil == l1 && nil == l2
}
