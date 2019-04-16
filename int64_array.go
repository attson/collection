package collection

import (
	"errors"
	"fmt"
)

type Int64Array struct{
	AbsArray
	objs []int64
}

func NewInt64Array(objs []int64) *Int64Array {
	arr := &Int64Array{
		objs:objs,
	}
	arr.AbsArray.Parent = arr
	arr.AbsArray.compare = func(i interface{}, i2 interface{}) int {
		int1 := i.(int64)
		int2 := i2.(int64)
		if int1 > int2 {
			return 1
		}
		if int1 < int2 {
			return -1
		}
		return 0
	}
	return arr
}

func (arr *Int64Array) Insert(index int, obj interface{}) IArray {
	if arr.Err() != nil {
		return arr
	}
	if i, ok := obj.(int64); ok {
		length := len(arr.objs)
		tail := arr.objs[index:length]
		arr.objs = append(arr.objs[0:index], i)
		arr.objs = append(arr.objs, tail...)
	} else {
		return arr.SetErr(errors.New("Insert: type error"))
	}
	return arr
}

func (arr *Int64Array) Remove(i int) IArray {
	if arr.Err() != nil {
		return arr
	}

	len := arr.Count()
	if i >= len {
		return arr.SetErr(errors.New("index exceeded"))
	}
	arr.objs = append(arr.objs[0:i], arr.objs[i+1: len]...)
	return arr
}

func (arr *Int64Array) NewEmpty(err ...error) IArray {
	intArr := NewInt64Array([]int64{})
	if len(err) != 0 {
		intArr.err = err[0]
	}
	return intArr
}


func (arr *Int64Array) Index(i int) IMix {
	return NewMix(arr.objs[i])
}

func (arr *Int64Array) Count() int {
	return len(arr.objs)
}

func (arr *Int64Array) DD() {
	ret := fmt.Sprintf("IntArray(%d):{\n", arr.Count())
	for k, v := range arr.objs {
		ret = ret + fmt.Sprintf("\t%d:\t%d\n",k, v)
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}