package singlelinkedlist

import "testing"

func TestIsPalindrome(t *testing.T) {
	l1 := NewLinkedList()
	l2 := NewLinkedList()
	l3 := NewLinkedList()
	var i, j int
	for i = 0; i < 10; i++ {
		l1.InsertTail(i)
		l2.InsertTail(i)
		l3.InsertTail(i)
	}

	l2.InsertTail(i)
	for j = i - 1; j >= 0; j-- {
		l1.InsertTail(j)
		l2.InsertTail(j)
		l3.InsertTail(j)
	}

	l3.InsertTail(12)

	l1.Print()
	l2.Print()
	l3.Print()

	data := []struct {
		list   *LinkedList
		result bool
	}{
		{l1, true},
		{l2, true},
		{l3, false},
	}

	for _, d := range data {
		if result := IsPalindrome(d.list); result != d.result {
			t.Errorf("got:%t, want:%t", result, d.result)
		}
	}
}
