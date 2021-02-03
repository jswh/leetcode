package adt

import (
	"errors"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func (list *LinkedList) len() int {
	length := 0
	currentNode := list.head
	for currentNode != nil {
		length++
	}
	return length
}
func (list *LinkedList) get(position int) (*ListNode, error) {
	if position < 0 {
		return nil, errors.New("position shoud >= 0")
	}
	listLen := list.len()
	if listLen < position+1 {
		return nil, errors.New("position overflow")
	}
	currentNode := list.head
	walked := 0
	for walked != position {
		currentNode = currentNode.Next
		walked++
	}

	return currentNode, nil
}

func (list *LinkedList) insert(data int, position int) error {
	node := ListNode{data, nil}

	if position == 0 {
		node.Next = list.head
		list.head = &node
		return nil
	}
	prenode, err := list.get(position - 1)
	if err != nil {
		return err
	}
	node.Next = prenode.Next
	prenode.Next = &node

	return nil
}

func (list *LinkedList) delete(position int) error {
	if position == 0 {
		list.head = list.head.Next
		return nil
	}
	prenode, err := list.get(position - 1)
	if err != nil {
		return err
	}
	node := prenode.Next
	prenode.Next = node.Next
	return nil
}
