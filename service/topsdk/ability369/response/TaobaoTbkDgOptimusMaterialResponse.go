package response

import (
	"topsdk/ability369/domain"
)

type TaobaoTbkDgOptimusMaterialResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   resultList
	*/
	ResultList []domain.TaobaoTbkDgOptimusMaterialMapData `json:"result_list,omitempty" `
	/*
	   推荐信息-是否抄底
	*/
	IsDefault string `json:"is_default,omitempty" `
	/*
	   商品总数-目前只有全品库商品查询有该字段
	*/
	TotalCount int64 `json:"total_count,omitempty" `
}
