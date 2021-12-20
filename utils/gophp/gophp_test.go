package gophp

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestPhp(t *testing.T) {
	fmt.Println(Unserialize([]byte(`a:1:{s:4:"test";i:34343;}`)))
	serialize, _ := Serialize(map[string]interface{}{
		"test": 34343,
	})
	fmt.Println(string(serialize))
}

type student struct {
	Amount     string `json:"amount"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	UserID     string `json:"user_id"`
}

func TestUnserialize(t *testing.T) {
	unserialize, _ := Unserialize([]byte(`a:2:{i:0;a:4:{s:7:"user_id";s:5:"10118";s:6:"amount";s:5:"69.00";s:11:"create_time";s:19:"2021-01-04 16:29:03";s:11:"update_time";s:19:"2021-06-15 16:02:46";}i:1;a:4:{s:7:"user_id";s:5:"10088";s:6:"amount";s:5:"10.00";s:11:"create_time";s:19:"2021-01-04 15:46:10";s:11:"update_time";s:19:"2021-06-15 15:50:06";}}`))
	fmt.Printf("%s\n", unserialize)
	arr, _ := json.Marshal(unserialize)
	fmt.Printf("arr：%s\n", arr)
	var stu []student
	_ = json.Unmarshal(arr, &stu)
	fmt.Printf("stu：%v\n", stu)
}
