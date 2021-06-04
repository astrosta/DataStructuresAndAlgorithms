package singlelinkedlist

import "testing"

func TestLinkedList_Reverse(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertHead(i + 1)
	}
	l.Print()
	l.Reverse()
	l.Print()
}
