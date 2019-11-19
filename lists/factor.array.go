package lists

import (
	"reflect"
)

type FactorArray struct {
	arr 				[]interface{}
	size    			uint32
	factor				uint32
	cap 				uint32
	initialCap			uint32
	genericType 		interface{}
	genericIsPointer 	bool
}



type FactorOptions struct {
	Factor 		uint32
	InitialCap	uint32
}

func NewFactorArray(opt *FactorOptions) *FactorArray {
	f := uint32(50)
	c := uint32(8)
	if opt != nil {
		if opt.Factor != 0 { f = opt.Factor }
		if opt.InitialCap != 0 { c = opt.InitialCap }
	}
	return &FactorArray{factor: f, cap: c, initialCap: c}
}

func (sa *FactorArray) Size() uint32 {
	return sa.size
}

func (sa *FactorArray) Add(item interface{}) {
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
	if sa.size == uint32(len(sa.arr)) {
		sa.resize()
	}
	sa.arr[sa.size] = item
	sa.size++
}

func (sa *FactorArray) Get(index uint32) (interface{}, error) {
	panic("implement me")
}

func (sa *FactorArray) Remove(index uint32) (interface{}, error) {
	panic("implement me")
}

func (sa *FactorArray) GetValue(index uint32, receiver interface{}) {
	if index > uint32(len(sa.arr) - 1) {
		panic("out of bounds")
	}
	if sa.genericIsPointer {
		reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(sa.arr[index]).Elem())
	} else {
		reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(sa.arr[index]))
	}
}

func (sa *FactorArray) RemoveValue(index uint32, receiver interface{}) {
	if index > sa.size {
		panic("out of bounds")
	}
	if receiver != nil && sa.genericIsPointer {
		reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(sa.arr[index]).Elem())
	} else if receiver != nil {
		reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(sa.arr[index]))
	}
	// Shrink the slice by slicing original slice until the index
	// and appending everything what after index + 1 with spread
	if sa.size > 1 {
		sa.arr = append(sa.arr[:index], sa.arr[index+1:]...)
		sa.cap--
	} else {
		sa.arr = nil
		sa.cap = sa.initialCap
	}
	sa.size--
}

func (sa *FactorArray) AddAtIndex(index uint32, item interface{}) {
	if index > sa.size + sa.factor {
		panic("out of bounds")
	}
	if index < uint32(len(sa.arr) - 1) {
		sa.arr[index] = item
		return
	}
	sa.resize()
	sa.arr[index] = item
}

func (sa *FactorArray) allocate(size uint32) []interface{} {
	return make([]interface{}, size)
}

func (sa *FactorArray) resize() {
	if sa.cap < sa.initialCap {
		sa.cap = sa.initialCap
	}
	newArr := sa.allocate(sa.size + sa.cap * sa.factor / 100)
	if sa.arr != nil {
		copy(newArr, sa.arr)
	}
	sa.arr = newArr
}
