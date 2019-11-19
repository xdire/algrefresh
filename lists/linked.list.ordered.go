package lists

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

type ifRepresentation struct {
	Type, Data unsafe.Pointer
	typ uintptr
	word unsafe.Pointer
}

type emptyInterface interface {}
type AbstractComparator func(a interface{}, b interface{}) (int8, error)

type nodeComparable struct {
	cmpIndex	interface{}
	next 		*nodeComparable
	data 		interface{}
}

func OrderedIntComparator(a interface{}, b interface{}) (int8, error) {
	aInt, ok := a.(int)
	bInt, ok1 := b.(int)
	if ok && ok1 {
		if aInt > bInt {
			return 1, nil
		} else if aInt < bInt {
			return -1, nil
		}
		return 0, nil
	}
	return 0, fmt.Errorf("non comparable values provided")
}

type LinkedListOrdered struct {
	head 				*nodeComparable
	tail				*nodeComparable
	size 				uint32
	genericType 		interface{}
	genericIsPointer 	bool
	comparator			func(a interface{}, b interface{}) (int8, error)
}



func NewLinkedListOrdered(cmp AbstractComparator) *LinkedListOrdered {
	return &LinkedListOrdered{
		head:             nil,
		tail:             nil,
		size:             0,
		genericType:      nil,
		genericIsPointer: false,
		comparator:       cmp,
	}
}

func (l *LinkedListOrdered) Once(item interface{}, order interface{}) interface{} {
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
	return l.addComparable(item, order, true)
}

func (l *LinkedListOrdered) GetOrderValue(order interface{}, receiver interface{}) error {
	_, err := l.orderGetter(order, receiver)
	return err
}

func (l *LinkedListOrdered) SetComparator(cmp func(a interface{}, b interface{}) (int8, error))  {
	l.comparator = cmp
}

func (l *LinkedListOrdered) Size() uint32 {
	return l.size
}

func (l *LinkedListOrdered) Add(item interface{}, order interface{}) {
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
	if l.comparator == nil {
		panic("no comparator presented for ordered list")
	}
	// Use method which can work with comparable
	l.addComparable(&item, order, false)
}

func (l *LinkedListOrdered) addComparable(item interface{}, order interface{}, once bool) interface{} {
	if l.head == nil {
		l.head = &nodeComparable{
			cmpIndex: order,
			next: nil,
			data: item,
		}
		l.tail = l.head
		l.size++
		return l.head.data
	} else {
		// Walk the list
		counter := 0
		next := l.head
		prev := l.head
		for next != nil {
			// Compare order of values
			res, err := l.comparator(next.cmpIndex, order)
			if err != nil {
				panic(err)
			}
			// Check if value is equal to order
			// and return back the real value to interface
			if once && res == 0 {
				return next.data
			}
			// If comparable value lesser than current value then
			// switch values
			if res == 1 {
				// If at head then switch head
				if counter == 0 {
					l.head = &nodeComparable{
						cmpIndex: 	order,
						next: 		prev,
						data: 		item,
					}
					l.size++
					return l.head.data
				} else {
				// Switch prev next to new node and link next
					prev.next = &nodeComparable{
						cmpIndex: 	order,
						next: 		next,
						data: 		item,
					}
					l.size++
					return prev.next.data
				}
			}
			// iterate
			prev = next
			next = next.next
			counter++
		}
		// If iteration didn't return any results
		// just add to the tail
		l.tail.next = &nodeComparable{
			cmpIndex: 	order,
			next: nil,
			data: item,
		}
		l.tail = l.tail.next
		l.size++
		return l.tail.data
	}
}

func (l *LinkedListOrdered) GetValue(index uint32, receiver interface{}) {
	_, err := l.getter(index, receiver)
	if err != nil {
		panic(err)
	}
}

func (l *LinkedListOrdered) Get(index uint32) (interface{}, error) {
	return l.getter(index, nil)
}

func (l *LinkedListOrdered) GetOrder(order interface{}) (interface{}, error) {
	return l.orderGetter(order, nil)
}

func (l *LinkedListOrdered) Remove(index uint32) (interface{}, error) {
	return l.remover(index, nil)
}

func (l *LinkedListOrdered) RemoveValue(index uint32, receiver interface{}) {
	_, err := l.remover(index, receiver)
	if err != nil {
		panic(err)
	}
}

func (l *LinkedListOrdered) remover(index uint32, receiver interface{}) (interface{}, error) {
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
	if receiver != nil && l.genericIsPointer {
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

func (l *LinkedListOrdered) getter(index uint32, receiver interface{}) (interface{}, error) {
	if l.head == nil || index > l.size - 1 {
		return nil, errors.New("index out of boundaries")
	}
	counter := uint32(0)
	next := l.head
	for counter != index {
		next = next.next
		counter++
	}
	if receiver != nil && l.genericIsPointer  {
		reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(next.data).Elem())
	} else if receiver != nil  {
		reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(next.data))
	}
	return next.data, nil
}

func (l *LinkedListOrdered) orderGetter(order interface{}, receiver interface{}) (interface{}, error) {
	next := l.head
	for next != nil {
		if next.cmpIndex == order {
			if receiver != nil && l.genericIsPointer {
				reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(next.data).Elem())
			} else if receiver != nil  {
				reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(next.data))
			}
			return next.data, nil
		}
		next = next.next
	}
	return nil, errors.New("not found")
}

func (l *LinkedListOrdered) Push(item interface{}) {
	l.Add(item, 0)
}

func (l *LinkedListOrdered) Poll() (interface{}, error) {
	if l.size == 0 {
		return nil, errors.New("empty")
	}
	return l.remover(0, nil)
}

func (l *LinkedListOrdered) Peek() (interface{}, error) {
	if l.size == 0 {
		return nil, errors.New("empty")
	}
	return l.getter(0, nil)
}

func (l *LinkedListOrdered) PollValue(receiver interface{}) error {
	if l.size == 0 {
		return errors.New("empty")
	}
	l.RemoveValue(0, receiver)
	return nil
}

func (l *LinkedListOrdered) PeekValue(receiver interface{}) error {
	if l.size == 0 {
		return errors.New("empty")
	}
	l.GetValue(0, receiver)
	return nil
}
