package lists

import (
	"errors"
	"reflect"
)

type node struct {
	next *node
	data interface{}
}

type LinkedList struct {
	head 				*node
	tail				*node
	size 				uint32
	genericType 		interface{}
	genericIsPointer 	bool
}

func (l *LinkedList) Size() uint32 {
	return l.size
}

func (l *LinkedList) Add(item interface{}) {
	if l.genericType == nil {
		vt := reflect.ValueOf(item).Type().Kind()
		if vt == reflect.Ptr {
			l.genericType = reflect.ValueOf(item).Elem().Type()
			l.genericIsPointer = true
		} else {
			l.genericType = vt
		}
		l.genericType = item
	}
	// If head is not defined then dedicate the separate procedure
	// to fulfill head and tail
	if l.head == nil {
		l.head = &node{
			next: nil,
			data: item,
		}
		l.tail = l.head
	} else {
	// Add straight to tail if head is presented
		l.tail.next = &node{
			next: nil,
			data: item,
		}
		l.tail = l.tail.next
	}
	// Increment size
	l.size++
}

func (l *LinkedList) Get(index uint32) (interface{}, error) {
	return l.getter(index, nil)
}

func (l *LinkedList) Remove(index uint32) (interface{}, error) {
	return l.remover(index, nil)
}

func (l *LinkedList) GetValue(index uint32, receiver interface{}) {
	_, err := l.getter(index, receiver)
	if err != nil {
		panic(err)
	}
}

func (l *LinkedList) RemoveValue(index uint32, receiver interface{}) {
	_, err := l.remover(index, receiver)
	if err != nil {
		panic(err)
	}
}

func (l *LinkedList) AddAtIndex(index uint32, item interface{}) {
	if l.head == nil && index == 0 {
		l.Add(item)
	}
	if index > l.size {
		panic("cannot add index which out of size bounds")
	}
	counter := uint32(0)
	prev := l.head
	next := l.head
	for counter != index {
		prev = next
		next = next.next
		counter++
	}
	prev.next = &node{
		next: next,
		data: item,
	}
	l.size++
}

func (l *LinkedList) getter(index uint32, receiver interface{})  (interface{}, error) {
	if l.head == nil || index > l.size - 1 {
		return nil, errors.New("index out of boundaries")
	}
	counter := uint32(0)
	next := l.head
	for counter != index {
		next = next.next
		counter++
	}
	if receiver != nil && l.genericIsPointer {
		reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(next.data).Elem())
	} else if receiver != nil  {
		reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(next.data))
	}
	return next.data, nil
}

func (l *LinkedList) remover(index uint32, receiver interface{})  (interface{}, error) {
	if l.head == nil || index > l.size - 1 {
		return nil, errors.New("index out of boundaries")
	}
	prev := l.head
	next := l.head
	// Check if index is zero, for that situation just replace
	// head with the next element from the head
	if index == 0 {
		next = l.head
		l.head = l.head.next
	} else {
		// If index is greater than zero we need to search until the index
		// and when found link previous element to the current next element
		counter := uint32(0)
		for counter != index {
			prev = next
			next = next.next
			counter++
		}
		prev.next = next.next
	}
	// Provide the value back to requester
	if receiver != nil && l.genericIsPointer  {
		reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(next.data).Elem())
	} else if receiver != nil  {
		reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(next.data))
	}
	// If deleted value was actual tail then replace tail with
	// previous element (obviously tail was deleted)
	if next == l.tail {
		l.tail = prev
	}
	// Decrement size
	l.size--
	// If size equals to 1 then Head will turn into Tail
	if l.size == 1 {
		l.head = l.tail
	}
	// If last element gone then reset the head and tail
	if l.size == 0 {
		l.head = nil
		l.tail = nil
	}
	return next.data, nil
}


