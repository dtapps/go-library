package dayuanren

/**
  http://dyr.greatmake.cn/192.html
  https://www.kancloud.cn/boyanyun/boyanyun_huafei
  https://www.showdoc.com.cn/dyr/11137624952363255
*/

const (
	ErrnoSuccess = 0 // 成功
)

type ErrorResponse struct {
	Errno  int64  `json:"errno"`  // 错误码，0代表成功，非0代表失败
	Errmsg string `json:"errmsg"` // 错误描述
}
