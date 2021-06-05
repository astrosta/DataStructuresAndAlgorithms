package queue

import (
	"fmt"
	"testing"
)

func TestArrayQueue_String(t *testing.T) {
	arrayQueue := NewArrayQueue(10)
	for i := 0; i < 10; i++ {
		arrayQueue.Enqueue(i)
	}

	fmt.Println(arrayQueue.String())

	for i := 0; i < 5; i++ {
		arrayQueue.Dequeue()
	}

	fmt.Println(arrayQueue.String())
}
