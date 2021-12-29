package dingdanxia

// Result 接口
type Result struct {
	Byte                      []byte // 内容
	Err                       error  // 错误
	JdJyOrderDetailsResponse         // 京佣订单
	WaimaiMeituanOrdersResult        // 美团联盟外卖/闪购/优选/酒店订单
}
