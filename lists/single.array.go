package lists

import (
	"fmt"
	"reflect"
)

type SingleArray struct {
	arr 				[]interface{}
	size    			uint32
	genericType 		interface{}
	genericIsPointer 	bool
}

func NewSingleArray() *SingleArray {
	return &SingleArray{}
}

func (sa *SingleArray) Size() uint32 {
	return sa.size
}

func (sa *SingleArray) Add(item interface{}) {
	if sa.genericType == nil {
		vt := reflect.ValueOf(item).Type().Kind()
		if vt == reflect.Ptr {
			sa.genericType = reflect.ValueOf(item).Elem().Type()
			sa.genericIsPointer = true
		} else {
			sa.genericType = vt
		}
		fmt.Printf("\nType Kind: %+v\n", sa.genericType)
		sa.genericType = item
	}
	if sa.arr == nil {
		sa.size = 1
		sa.arr = sa.allocate(sa.size)
		sa.arr[sa.size - 1] = item
		return
	}
	sa.resize(sa.size + 1)
	sa.arr[sa.size - 1] = item
}

func (sa *SingleArray) Get(index uint32) (interface{}, error) {
	panic("implement me")
}

func (sa *SingleArray) Remove(index uint32) (interface{}, error) {
	panic("implement me")
}

func (sa *SingleArray) GetValue(index uint32, receiver interface{}) {
	if index > sa.size - 1 {
		panic("out of bounds")
	}
	if sa.genericIsPointer {
		reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(sa.arr[index]).Elem())
	} else {
		reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(sa.arr[index]))
	}
}

func (sa *SingleArray) RemoveValue(index uint32, receiver interface{}) {
	if index > sa.size {
		panic("out of bounds")
	}
	if sa.genericIsPointer {
		reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(sa.arr[index]).Elem())
	} else {
		reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(sa.arr[index]))
	}
	// Shrink the slice by slicing original slice until the index
	// and appending everything what after index + 1 with spread
	if sa.size > 1 {
		sa.arr = append(sa.arr[:index], sa.arr[index+1:]...)
	} else {
		sa.arr = nil
	}
	sa.size--
}

func (sa *SingleArray) AddAtIndex(index uint32, item interface{}) {
	if index > sa.size {
		panic("out of bounds")
	}
	sa.resize(sa.size + 1)
	sa.arr[sa.size - 1] = item
}

func (sa *SingleArray) allocate(size uint32) []interface{} {
	return make([]interface{}, size)
}

func (sa *SingleArray) resize(size uint32) {
	newArr := sa.allocate(size)
	copy(newArr, sa.arr)
	sa.arr = newArr
	sa.size = size
}
