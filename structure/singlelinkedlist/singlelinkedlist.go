package singlelinkedlist

import "fmt"

/*
单链表的基本实现
author: astrosta
*/

type ListNode struct {
	value interface{}
	next  *ListNode
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{v, nil}
}

func (node *ListNode) GetValue() interface{} {
	return node.value
}

func (node *ListNode) GetNext() *ListNode {
	return node.next
}

type LinkedList struct {
	head   *ListNode
	length uint
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

func (list *LinkedList) InsertAfter(p *ListNode, value interface{}) bool {
	if nil == p {
		return false
	}

	node := NewListNode(value)
	oldNext := p.next
	p.next = node
	node.next = oldNext
	list.length++

	return true
}

func (list *LinkedList) InsertBefore(p *ListNode, value interface{}) bool {
	if nil == p || p == list.head {
		return false
	}

	//首先查找p的前一个节点
	cur := list.head

	for cur.next != nil {
		if cur.next == p {
			break
		}
		cur = cur.next
	}

	if cur.next == nil {
		return false
	}

	node := NewListNode(value)
	cur.next = node
	node.next = p
	list.length++

	return true
}

func (list *LinkedList) InsertHead(value interface{}) bool {
	return list.InsertAfter(list.head, value)
}

func (list *LinkedList) InsertTail(value interface{}) bool {
	cur := list.head
	for cur.next != nil {
		cur = cur.next
	}

	return list.InsertAfter(cur, value)
}

func (list *LinkedList) FindByIndex(index uint) *ListNode {
	if index > list.length {
		return nil
	}

	cur := list.head
	var i uint
	for i = 0; i < index; i++ {
		cur = cur.next
	}

	return cur
}

func (list *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}

	cur := list.head
	for cur.next != nil {
		if cur.next == p {
			break
		}

		cur = cur.next
	}

	if cur.next == nil {
		return false
	}

	cur.next = p.next
	p = nil
	list.length--
	return true
}

func (list *LinkedList) Print() {
	cur := list.head.next
	var format string
	for cur != nil {
		format += fmt.Sprintf("%v", cur.value)
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}

	fmt.Println(format)
}
