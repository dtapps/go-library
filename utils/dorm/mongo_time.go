package dorm

import (
	"github.com/dtapps/go-library/utils/gotime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"time"
)

// BsonTime 时间类型
type BsonTime time.Time

// MarshalJSON 实现json序列化
func (bt BsonTime) MarshalJSON() ([]byte, error) {

	b := make([]byte, 0)
	b = append(b, '"')
	b = append(b, gotime.SetCurrent(time.Time(bt)).Bson()...)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON 实现json反序列化
func (bt *BsonTime) UnmarshalJSON(data []byte) (err error) {

	if string(data) == "null" {
		return nil
	}

	bsonTime := gotime.SetCurrentParse(string(data))
	*bt = BsonTime(bsonTime.Time)

	return nil
}

// Time 转时间
func (bt BsonTime) Time() time.Time {
	return gotime.SetCurrent(time.Time(bt)).Time
}

// Format 转时间字符串
func (bt BsonTime) Format() string {
	return gotime.SetCurrent(time.Time(bt)).Format()
}

// TimePro 转时间操作
func (bt BsonTime) TimePro() gotime.Pro {
	return gotime.SetCurrent(time.Time(bt))
}

// NewBsonTimeCurrent 创建当前时间
func NewBsonTimeCurrent() BsonTime {
	return BsonTime(gotime.Current().Time)
}

// NewBsonTimeFromTime 创建某个时间
func NewBsonTimeFromTime(t time.Time) BsonTime {
	return BsonTime(t)
}

// NewBsonTimeFromString 创建某个时间 字符串
func NewBsonTimeFromString(t string) BsonTime {
	return BsonTime(gotime.SetCurrentParse(t).Time)
}

// Value 时间类型
func (bt BsonTime) Value() string {
	return gotime.SetCurrent(time.Time(bt)).Bson()
}

// MarshalBSONValue 实现bson序列化
func (bt BsonTime) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(gotime.SetCurrent(time.Time(bt)).Bson())
}

// UnmarshalBSONValue 实现bson反序列化
func (bt *BsonTime) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	t1 := gotime.SetCurrentParse(string(data))
	//if string(data) == "" {
	//	return errors.New(fmt.Sprintf("%s, %s, %s", "读取数据失败:", t, data))
	//}
	*bt = BsonTime(t1.Time)
	return nil
}
