package stack

import "fmt"

type node struct {
	next  *node
	value interface{}
}

type LinkedListStack struct {
	topNode *node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (ls *LinkedListStack) IsEmpty() bool {
	return ls.topNode == nil
}

func (ls *LinkedListStack) Push(value interface{}) {
	ls.topNode = &node{
		value: value,
		next:  ls.topNode,
	}
}

func (ls *LinkedListStack) Pop() interface{} {
	if ls.IsEmpty() {
		return nil
	}
	v := ls.topNode.value
	ls.topNode = ls.topNode.next

	return v
}

func (ls *LinkedListStack) Top() interface{} {
	if ls.IsEmpty() {
		return nil
	}

	return ls.topNode.value
}

func (ls *LinkedListStack) Flush() {
	ls.topNode = nil
}

func (ls *LinkedListStack) Print() {
	if ls.IsEmpty() {
		fmt.Println("empty stack")
	}

	tmpNode := ls.topNode
	for tmpNode != nil {
		fmt.Println(ls.topNode.value)
		tmpNode = tmpNode.next
	}
}
