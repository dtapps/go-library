package dorm

type Iterator struct {
	data  []interface{}
	index int
}

// NewIterator 构造函数
func NewIterator(data []interface{}) *Iterator {
	return &Iterator{data: data}
}

// HasNext 是否有下一个
func (i *Iterator) HasNext() bool {
	if i.data == nil || len(i.data) == 0 {
		return false
	}
	return i.index < len(i.data)
}

// Next 循环下一个
func (i *Iterator) Next() (Ret interface{}) {
	Ret = i.data[i.index]
	i.index = i.index + 1
	return
}
