package gophp

import "gopkg.in/dtapps/go-library.v3/utils/gophp/serialize"

// Serialize 序列
func Serialize(value interface{}) ([]byte, error) {
	return serialize.Marshal(value)
}

// Unserialize 反序列
func Unserialize(data []byte) (interface{}, error) {
	return serialize.UnMarshal(data)
}
