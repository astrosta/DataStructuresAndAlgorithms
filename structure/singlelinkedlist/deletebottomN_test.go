package singlelinkedlist

import "testing"

func TestLinkedList_DeleteBottomN(t *testing.T) {
	l1 := NewLinkedList()
	l2 := NewLinkedList()
	l3 := NewLinkedList()
	for i := 0; i < 10; i++ {
		l1.InsertTail(i)
		l2.InsertTail(i)
		l3.InsertTail(i)
	}

	l1.DeleteBottomN(3)
	l2.DeleteBottomN(10)
	l3.DeleteBottomN(11)
	l1.Print()
	l2.Print()
	l3.Print()
}
