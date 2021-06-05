package queue

/*
循环队列会损失一个数组空间
*/
import "fmt"

type CircularQueue struct {
	value    []interface{}
	tail     int
	head     int
	capacity int
}

func NewCircularQueue(n int) *CircularQueue {
	return &CircularQueue{
		value:    make([]interface{}, n),
		tail:     0,
		head:     0,
		capacity: n,
	}
}

func (cq *CircularQueue) Enqueue(value interface{}) bool {
	if (cq.tail+1)%cq.capacity == cq.head {
		return false
	}

	cq.value[cq.tail] = value
	cq.tail = (cq.tail + 1) % cq.capacity

	return true
}

func (cq *CircularQueue) Dequeue() interface{} {
	if cq.head == cq.tail {
		return nil
	}

	value := cq.value[cq.head]
	cq.head = (cq.head + 1) % cq.capacity
	return value
}

func (cq *CircularQueue) String() interface{} {
	if cq.head == cq.tail {
		return "empty queue"
	}

	result := "head"

	cur := cq.head
	for ; cur%cq.capacity != cq.tail; cur++ {
		result += fmt.Sprint("<-", cq.value[cur])
	}

	result += "<-tail"
	return result
}
