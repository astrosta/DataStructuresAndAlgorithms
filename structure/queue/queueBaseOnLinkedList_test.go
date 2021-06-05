package queue

import (
	"fmt"
	"testing"
)

func TestLinkedListQueue_String(t *testing.T) {
	linkedListQueue := NewLinkedListQueue()
	for i := 0; i < 20; i++ {
		linkedListQueue.Enqueue(i)
	}

	fmt.Println(linkedListQueue.String())

	for i := 0; i < 5; i++ {
		linkedListQueue.Dequeue()
	}

	fmt.Println(linkedListQueue.String())
}
