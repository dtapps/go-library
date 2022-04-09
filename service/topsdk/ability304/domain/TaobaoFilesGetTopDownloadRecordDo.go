package domain

import (
	"topsdk/util"
)

type TaobaoFilesGetTopDownloadRecordDo struct {
	/*
	   下载链接     */
	Url *string `json:"url,omitempty" `

	/*
	   文件创建时间     */
	Created *util.LocalTime `json:"created,omitempty" `

	/*
	   下载链接状态。1:未下载。2:已下载     */
	Status *int64 `json:"status,omitempty" `
}

func (s *TaobaoFilesGetTopDownloadRecordDo) SetUrl(v string) *TaobaoFilesGetTopDownloadRecordDo {
	s.Url = &v
	return s
}
func (s *TaobaoFilesGetTopDownloadRecordDo) SetCreated(v util.LocalTime) *TaobaoFilesGetTopDownloadRecordDo {
	s.Created = &v
	return s
}
func (s *TaobaoFilesGetTopDownloadRecordDo) SetStatus(v int64) *TaobaoFilesGetTopDownloadRecordDo {
	s.Status = &v
	return s
}
