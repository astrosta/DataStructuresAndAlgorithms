package queue

import (
	"fmt"
	"testing"
)

func TestCircularQueue_String(t *testing.T) {
	circularQueue := NewCircularQueue(10)
	for i := 0; i < 10; i++ {
		circularQueue.Enqueue(i)
	}

	fmt.Println("circularQueue: ", circularQueue.String())
	for i := 0; i < 5; i++ {
		fmt.Println(circularQueue.Dequeue())
	}

	fmt.Println("circularQueue: ", circularQueue.String())
}
