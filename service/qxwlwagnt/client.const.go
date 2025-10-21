package qxwlwagnt

type BaseResponse struct {
	Status  int    `json:"status"`  // 接口调用状态码
	Message string `json:"message"` // 状态描述
}

type CommonResponse[T any] struct {
	BaseResponse   // 嵌入公共基础结构体
	Result       T `json:"result,omitempty"` // 返回数据，操作成功时有值，操作失败时为空（具体类型由 T 决定）
}
