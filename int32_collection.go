package collection

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Int32Collection struct {
	AbsCollection
	objs []int32
}

func compareInt32(i interface{}, i2 interface{}) int {
	int1 := i.(int32)
	int2 := i2.(int32)
	if int1 > int2 {
		return 1
	}
	if int1 < int2 {
		return -1
	}
	return 0
}

// NewInt32Collection create a new Int32Collection
func NewInt32Collection(objs []int32) *Int32Collection {
	arr := &Int32Collection{
		objs: objs,
	}
	arr.AbsCollection.Parent = arr
	arr.SetCompare(compareInt32)
	return arr
}

// Copy copy collection
func (arr *Int32Collection) Copy() ICollection {
	objs2 := make([]int32, len(arr.objs))
	copy(objs2, arr.objs)
	arr.objs = objs2
	return arr
}

func (arr *Int32Collection) Insert(index int, obj interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}
	if i, ok := obj.(int32); ok {
		length := len(arr.objs)

		// 如果是append操作，直接调用系统的append，不新创建collection
		if index == length {
			arr.objs = append(arr.objs, i)
			return arr
		}

		new := arr.objs[0:index]
		new = append(new, i)
		new = append(new, arr.objs[index:length]...)
		arr.objs = new
	} else {
		return arr.SetErr(errors.New("Insert: type error"))
	}
	return arr
}

func (arr *Int32Collection) Remove(i int) ICollection {
	if arr.Err() != nil {
		return arr
	}

	len := arr.Count()
	if i >= len {
		return arr.SetErr(errors.New("index exceeded"))
	}
	arr.objs = append(arr.objs[0:i], arr.objs[i+1:len]...)
	return arr
}

func (arr *Int32Collection) NewEmpty(err ...error) ICollection {
	int32Arr := NewInt32Collection([]int32{})
	if len(err) != 0 {
		int32Arr.err = err[0]
	}
	return int32Arr
}

func (arr *Int32Collection) Index(i int) IMix {
	return NewMix(arr.objs[i]).SetCompare(arr.compare)
}
func (arr *Int32Collection) SetIndex(i int, val interface{}) ICollection {
	arr.objs[i] = val.(int32)
	return arr
}

func (arr *Int32Collection) Count() int {
	return len(arr.objs)
}

func (arr *Int32Collection) DD() {
	ret := fmt.Sprintf("Int32Collection(%d):{\n", arr.Count())
	for k, v := range arr.objs {
		ret = ret + fmt.Sprintf("\t%d:\t%d\n", k, v)
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}

func (arr *Int32Collection) ToJson() ([]byte, error) {
	return json.Marshal(arr.objs)
}
