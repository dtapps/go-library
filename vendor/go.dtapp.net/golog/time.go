package golog

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"go.dtapp.net/gotime"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type TimeString struct {
	Time time.Time
}

// Value 插入数据，把时间转字符串
func (t TimeString) Value() (driver.Value, error) {
	return gotime.SetCurrent(t.Time).Format(), nil
}

// Scan 查询数据，把字符串转时间
func (t *TimeString) Scan(value interface{}) error {

	// 如果是空值，直接返回
	data, ok := value.(string)
	if !ok {
		return errors.New(fmt.Sprint("无法解析:", value))
	}

	// 解析时间
	result := gotime.SetCurrentParse(data)

	*t = TimeString{
		Time: result.Time,
	}

	return nil
}

// MarshalJSON JSON序列化
func (t TimeString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, gotime.SetCurrent(t.Time).Format())), nil
}

// UnmarshalJSON JSON反序列化
func (t *TimeString) UnmarshalJSON(data []byte) (err error) {
	// 删除双引号
	if data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}
	result := gotime.SetCurrentParse(string(data))
	*t = TimeString{
		Time: result.Time,
	}
	return
}

func (t TimeString) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	// 使用 field.Tag、field.TagSettings 获取字段的 tag
	// 查看 https://github.com/go-gorm/gorm/blob/master/schema/field.go 获取全部的选项

	// 根据不同的数据库驱动返回不同的数据类型
	switch db.Dialector.Name() {
	case "mysql", "sqlite":
		return "varchar"
	case "postgres":
		return "text"
	}
	return ""
}
