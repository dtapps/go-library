package dorm

import (
	"go.dtapp.net/library/utils/gotime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"time"
)

// BsonTime 类型
type BsonTime time.Time

// Value 时间类型
func (t BsonTime) Value() string {
	return gotime.SetCurrent(time.Time(t)).Bson()
}

// MarshalJSON 实现json序列化
func (t BsonTime) MarshalJSON() ([]byte, error) {
	//log.Println("MarshalJSON")
	b := make([]byte, 0)
	b = append(b, gotime.SetCurrent(time.Time(t)).Bson()...)
	return b, nil
}

// UnmarshalJSON 实现json反序列化
func (t *BsonTime) UnmarshalJSON(data []byte) (err error) {
	//log.Println("UnmarshalJSON")
	t1 := gotime.SetCurrentParse(string(data))
	*t = BsonTime(t1.Time)
	return
}

// MarshalBSONValue 实现bson序列化
func (t BsonTime) MarshalBSONValue() (bsontype.Type, []byte, error) {
	//log.Println("MarshalBSONValue")
	targetTime := gotime.SetCurrent(time.Time(t)).Bson()
	return bson.MarshalValue(targetTime)
}

// UnmarshalBSONValue 实现bson反序列化
func (t *BsonTime) UnmarshalBSONValue(t2 bsontype.Type, data []byte) error {
	//log.Println("UnmarshalBSONValue")
	t1 := gotime.SetCurrentParse(string(data))
	//if string(data) == "" {
	//	return errors.New(fmt.Sprintf("%s, %s, %s", "读取数据失败:", t2, data))
	//}
	*t = BsonTime(t1.Time)
	return nil
}
