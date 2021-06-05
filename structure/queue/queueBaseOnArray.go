package queue

import "fmt"

type ArrayQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{
		q:        make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

func (aq *ArrayQueue) Enqueue(value interface{}) bool {
	if aq.tail == aq.capacity {
		return false
	}

	aq.q[aq.tail] = value
	aq.tail++

	return true
}

func (aq *ArrayQueue) Dequeue() interface{} {
	if aq.head == aq.tail {
		return nil
	}

	value := aq.q[aq.head]
	aq.head++

	return value
}

func (aq *ArrayQueue) String() interface{} {
	if aq.head == aq.tail {
		return "empty queue"
	}

	result := "head"

	for i := aq.head; i < aq.tail; i++ {
		result += fmt.Sprint("<-", aq.q[i])
	}

	result += "<-tail"
	return result
}
