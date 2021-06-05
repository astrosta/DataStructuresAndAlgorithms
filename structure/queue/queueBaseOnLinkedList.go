package queue

import (
	"fmt"
)

type node struct {
	value interface{}
	next  *node
}

type LinkedListQueue struct {
	head   *node
	tail   *node
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (llq *LinkedListQueue) Enqueue(value interface{}) bool {
	temp := &node{
		value: value,
		next:  nil,
	}

	if llq.tail == nil {
		llq.tail = temp
		llq.head = temp
	} else {
		llq.tail.next = temp
		llq.tail = temp
	}
	llq.length++

	return true
}

func (llq *LinkedListQueue) Dequeue() interface{} {
	if llq.head == nil {
		return nil
	}

	value := llq.head.value
	llq.head = llq.head.next
	llq.length--

	return value
}

func (llq *LinkedListQueue) String() string {
	if llq.length == 0 {
		return "empty queue"
	}

	result := "head"
	cur := llq.head
	for cur != nil {
		result += fmt.Sprint("<-", cur.value)
		cur = cur.next
	}

	result += "<-tail"
	return result
}
