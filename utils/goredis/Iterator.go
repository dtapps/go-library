package goredis

type Iterator struct {
	data  []interface{}
	index int
}

func NewIterator(data []interface{}) *Iterator {
	return &Iterator{data: data}
}
func (i *Iterator) HasNext() bool {
	if i.data == nil || len(i.data) == 0 {
		return false
	}
	return i.index < len(i.data)
}

func (i *Iterator) Next() (Ret interface{}) {
	Ret = i.data[i.index]
	i.index = i.index + 1
	return
}
