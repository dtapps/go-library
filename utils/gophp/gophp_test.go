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

func TestUnserialize2(t *testing.T) {
	unserialize, err := Unserialize([]byte(`a:3:{s:4:"self";a:5:{s:2:"id";i:32;s:7:"user_id";i:10118;s:16:"merchant_user_id";i:10150;s:5:"level";i:3;s:5:"split";d:2.825;}s:5:"split";a:2:{i:0;a:7:{s:2:"id";i:12;s:7:"user_id";i:10000;s:13:"agent_user_id";i:10088;s:16:"user_golden_bean";d:5;s:5:"level";i:1;s:11:"calculation";d:0.010000000000000002;s:5:"split";d:0.57;}i:1;a:7:{s:2:"id";i:19;s:7:"user_id";i:10088;s:13:"agent_user_id";i:10118;s:16:"user_golden_bean";d:4;s:5:"level";i:2;s:11:"calculation";d:0.04;s:5:"split";d:2.26;}}s:11:"golden_bean";a:6:{s:4:"fans";d:56.5;s:12:"fan_referrer";d:16.95;s:8:"merchant";d:4.5200000000000005;s:5:"agent";d:2.825;s:5:"count";d:80.795;s:2:"bf";a:9:{s:2:"id";i:2;s:7:"user_id";i:10088;s:4:"fans";d:50;s:12:"fan_referrer";d:30;s:8:"merchant";d:8;s:6:"system";d:0;s:5:"agent";d:5;s:11:"create_time";s:19:"2020-11-20 14:47:21";s:11:"update_time";s:19:"2020-11-20 14:47:21";}}}`))
	fmt.Printf("%+v\n", unserialize)
	if err != nil {
		fmt.Printf("err：%v\n", err)
		return
	}
}
