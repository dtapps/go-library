package fly

import (
	"github.com/daodao97/fly/interval/xtype"
)

type Validator struct {
	Field  string
	Msg    string
	Handle []ValidateHandle
}

type ValidateHandle = func(m Model, row map[string]interface{}, val interface{}) (ok bool, err error)

type ValidateHandleMaker = func(field string) ValidateHandle

// Required field value must exist and not zero
func Required(field string) ValidateHandle {
	return func(m Model, row map[string]interface{}, val interface{}) (ok bool, err error) {
		v, ok := row[field]
		if !ok {
			return false, nil
		}
		return xtype.Bool(v), nil
	}
}

// IfRequired if field1 is existed, then field2 must exist and not zero
func IfRequired(field1 string) ValidateHandleMaker {
	return func(field string) ValidateHandle {
		return func(m Model, row map[string]interface{}, val interface{}) (ok bool, err error) {
			h := Required(field1)
			ok, err = h(m, row, row[field1])
			if err != nil {
				return false, err
			}
			if !ok {
				return true, nil
			}
			h = Required(field)
			return h(m, row, row[field])
		}
	}
}

// Unique field value must unique in current table
func Unique(field string) ValidateHandle {
	return func(m Model, row map[string]interface{}, val interface{}) (ok bool, err error) {
		opts := []Option{WhereEq(field, val)}
		if id, ok := row[m.PrimaryKey()]; ok {
			opts = append(opts, WhereNotEq(m.PrimaryKey(), id))
		}
		count, err := m.Count(opts...)
		if err != nil {
			return false, err
		}
		return count == 0, nil
	}
}
