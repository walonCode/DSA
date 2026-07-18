package main

import (
	"fmt"
	"strings"
)

type node[T comparable] struct {
	value T
	prev *node[T]
	next *node[T]
}

type linkedList[T comparable]struct {
	head *node[T]
	tail *node[T]
	size int
}

func newLinkedList[T comparable]() *linkedList[T] {
	return &linkedList[T]{}
}

func (l *linkedList[T])pushBack(value T) {
	node := &node[T]{value: value}
	if l.head == nil {
		l.head = node
		l.tail = node
	}else {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	}

	l.size++
}

func (l *linkedList[T])pushFront(value T){
	node := &node[T]{value:value, next: l.head}

	if l.head == nil {
		l.head = node
		l.tail = node
	}else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.size++
}

func (l *linkedList[T])popFront()(T, bool) {
	var zero T
	if l.head == nil {
		return zero, false
	}

	value := l.head.value
	l.head = l.head.next

	if l.head == nil {
		l.tail = nil 
	}else {
		l.head.prev = nil
	}
	
	l.size--
	return value, true
}

func (l *linkedList[T])get(idx int)(T, bool){
	var zero T
	if idx < 0 || idx >= l.size {
		return zero, false
	}

	curr := l.head
	for _ = range idx {
		curr = curr.next
	}

	return curr.value, true
}

func (l *linkedList[T])string()string {
	var result strings.Builder
	result.WriteString("[ ")
	curr := l.head
	for curr != nil {
		fmt.Fprintf(&result,"%v", curr.value)
		if curr.next != nil {
			result.WriteString(" -> ")
		}
		curr = curr.next
	}

	result.WriteString(" ]")
	return result.String()
}

func (l *linkedList[T])popBack()(T, bool){
	var zero T
	if l.tail == nil {
		return zero, false
	}

	value := l.tail.value
	l.tail = l.tail.prev

	if l.tail == nil {
		l.head = nil
	}else {
		l.tail.next = nil 
	}

	l.size--
	return value, true
}


func (l *linkedList[T])remove(value T) bool {
	curr := l.head

	for curr != nil {
		if curr.value == value {
			l.removeNode(curr)
			return true
		}
		curr = curr.next
	}

	return false
}

func(l *linkedList[T])removeNode(curr *node[T]) {
	if curr.prev != nil {
		curr.prev.next = curr.next
	}else {
		l.head = curr.next
	}

	if curr.next != nil {
		curr.next.prev = curr.prev
	}else {
		l.tail = curr.prev
	}
	l.size--
}

func (l *linkedList[T])insertAt(idx int, value T)bool {
	if idx < 0 || idx > l.size {
		return false
	}

	if idx == 0 {
		l.pushFront(value)
		return true
	}

	if idx == l.size {
		l.pushBack(value)
		return true
	}

	curr := l.head 
	for _ = range idx - 1 {
		curr = curr.next
	}

	node := &node[T]{value: value, next: curr, prev: curr.prev}
	curr.prev.next = node
	curr.prev = node

	l.size++
	return true
}