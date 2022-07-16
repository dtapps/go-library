package fly

import (
	"database/sql"
)

type With = func(*model)

func WithDB(db *sql.DB) With {
	return func(b *model) {
		b.client = db
	}
}

func WithConn(name string) With {
	return func(b *model) {
		b.connection = name
	}
}

func WithFakeDelKey(name string) With {
	return func(b *model) {
		b.fakeDelKey = name
	}
}

func WithPrimaryKey(name string) With {
	return func(b *model) {
		b.primaryKey = name
	}
}

func ColumnHook(columnHook ...Hook) With {
	return func(b *model) {
		if b.columnHook == nil {
			b.columnHook = make(map[string]HookData)
		}
		for _, v := range columnHook {
			f, h := v()
			b.columnHook[f] = h
		}
	}
}

// ColumnValidator while validate data by validator when create or update event
func ColumnValidator(validator ...Validator) With {
	return func(b *model) {
		if b.columnValidator == nil {
			b.columnValidator = make([]Validator, 0, len(validator))
		}
		b.columnValidator = append(b.columnValidator, validator...)
	}
}

func Validate(field, msg string, handle ...ValidateHandleMaker) (v Validator) {
	v.Field = field
	v.Msg = msg
	v.Handle = make([]ValidateHandle, 0, len(handle))
	for _, h := range handle {
		v.Handle = append(v.Handle, h(field))
	}

	return v
}

func WithSaveZero() With {
	return func(b *model) {
		b.saveZero = true
	}
}

func HasOne(opts ...HasOpts) With {
	return func(b *model) {
		if b.hasOne == nil {
			b.hasOne = make([]HasOpts, 0)
		}
		for i, v := range opts {
			if v.Conn == "" {
				opts[i].Conn = "default"
			}
			if v.LocalKey == "" {
				opts[i].LocalKey = "id"
			}
			if v.ForeignKey == "" {
				opts[i].ForeignKey = "id"
			}
		}
		b.hasOne = append(b.hasOne, opts...)
	}
}

func HasMany(opts ...HasOpts) With {
	return func(b *model) {
		if b.hasMany == nil {
			b.hasMany = make([]HasOpts, 0)
		}
		for i, v := range opts {
			if v.Conn == "" {
				opts[i].Conn = "default"
			}
			if v.LocalKey == "" {
				opts[i].LocalKey = "id"
			}
			if v.ForeignKey == "" {
				opts[i].ForeignKey = "id"
			}
		}
		b.hasMany = append(b.hasMany, opts...)
	}
}
