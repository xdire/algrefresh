package arrays

import (
	"reflect"
)

type VectorArray struct {
	arr 				[]interface{}
	size    			uint32
	vector				uint32
	genericType 		interface{}
	genericIsPointer 	bool
}

func NewVectorArray(vector uint32) *VectorArray {
	return &VectorArray{vector:vector}
}

func (sa *VectorArray) Size() uint32 {
	return sa.size
}

func (sa *VectorArray) Add(item interface{}) {
	if sa.genericType == nil {
		vt := reflect.ValueOf(item).Type().Kind()
		if vt == reflect.Ptr {
			sa.genericType = reflect.ValueOf(item).Elem().Type()
			sa.genericIsPointer = true
		} else {
			sa.genericType = vt
		}
		sa.genericType = item
	}
	if sa.arr == nil {
		sa.resize()
		sa.arr[sa.size] = item
		sa.size++
		return
	}
	if int(sa.size) == len(sa.arr) {
		sa.resize()
	}
	sa.arr[sa.size] = item
	sa.size++
}

func (sa *VectorArray) Get(index uint32, receiver interface{}) {
	if index > uint32(len(sa.arr) - 1) {
		panic("out of bounds")
	}
	if sa.genericIsPointer {
		reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(sa.arr[index]).Elem())
	} else {
		reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(sa.arr[index]))
	}
}

func (sa *VectorArray) Remove(index uint32, receiver interface{}) {
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

func (sa *VectorArray) AddAtIndex(index uint32, item interface{}) {
	if index > sa.size + sa.vector {
		panic("out of bounds")
	}
	if index < uint32(len(sa.arr) - 1) {
		sa.arr[index] = item
		return
	}
	sa.resize()
	sa.arr[index] = item
}

func (sa *VectorArray) allocate(size uint32) []interface{} {
	return make([]interface{}, size)
}

func (sa *VectorArray) resize() {
	if sa.vector == 0 {
		sa.vector = 8
	}
	newArr := sa.allocate(sa.size + sa.vector)
	if sa.arr != nil {
		copy(newArr, sa.arr)
	}
	sa.arr = newArr
}
