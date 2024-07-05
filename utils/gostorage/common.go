package gostorage

// FileInfo 上传文件的信息
type FileInfo struct {
	Path string `json:"path"` // 文件路径
	Name string `json:"name"` // 文件名称
	Url  string `json:"url"`  // 文件地址
}
