package dorm

import "time"

type empty struct{}

const (
	AttrExpr = "expr" // 过期时间
	AttrNx   = "nx"   // 设置Nx
)

type OperationAttr struct {
	Name  string
	Value interface{}
}

type OperationAttrs []*OperationAttr

func (a OperationAttrs) Find(name string) interface{} {
	for _, attr := range a {
		if attr.Name == name {
			return attr.Value
		}
	}
	return nil
}

// WithExpire 过期时间
func WithExpire(t time.Duration) *OperationAttr {
	return &OperationAttr{Name: AttrExpr, Value: t}
}

func WithNX() *OperationAttr {
	return &OperationAttr{Name: AttrNx, Value: empty{}}
}
