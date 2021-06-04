package singlelinkedlist

import "testing"

func TestLinkedList_HasCycle(t *testing.T) {
	l1 := NewLinkedList()
	l2 := NewLinkedList()
	var i int
	for i = 0; i < 10; i++ {
		l1.InsertTail(i)
		l2.InsertTail(i)
	}

	tempNode := l2.head.next

	for {
		if tempNode.next == nil {
			tempNode.next = l2.head.next
			break
		}

		tempNode = tempNode.next
	}

	data := []struct {
		list   *LinkedList
		result bool
	}{
		{l1, false},
		{l2, true},
	}

	for _, d := range data {
		if result := d.list.HasCycle(); result != d.result {
			t.Errorf("got:%t, want:%t", result, d.result)
		}
	}
}
