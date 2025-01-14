package cloud

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type Rpc2AiCustomV1WenxinworkshopChatYi34bChatResponse struct {
	Id               string `json:"id"`                    // 本轮对话的id
	Object           string `json:"object"`                // 回包类型
	Created          int64  `json:"created"`               // 时间戳
	SentenceId       int64  `json:"sentence_id,omitempty"` // 表示当前子句的序号。只有在流式接口模式下会返回该字段
	IsEnd            int64  `json:"is_end,omitempty"`      // 表示当前子句是否是最后一句。只有在流式接口模式下会返回该字段
	IsTruncated      bool   `json:"is_truncated"`          // 当前生成的结果是否被截断
	Result           string `json:"result"`                // 对话返回结果
	NeedClearHistory bool   `json:"need_clear_history"`    // 表示用户输入是否存在安全，是否关闭当前会话，清理历史会话信息  true：是，表示用户输入存在安全风险，建议关闭当前会话，清理历史会话信息 false：否，表示用户输入无安全风险
	BanRound         int64  `json:"ban_round"`             // 当need_clear_history为true时，此字段会告知第几轮对话有敏感信息，如果是当前问题，ban_round=-1
	Usage            struct {
		PromptTokens     int64 `json:"prompt_tokens"`     // 问题tokens数
		CompletionTokens int64 `json:"completion_tokens"` // 回答tokens数
		TotalTokens      int64 `json:"total_tokens"`      // tokens总数
	} `json:"usage"` // token统计信息
}

type Rpc2AiCustomV1WenxinworkshopChatYi34bChatResult struct {
	Result Rpc2AiCustomV1WenxinworkshopChatYi34bChatResponse // 结果
	Body   []byte                                            // 内容
	Http   gorequest.Response                                // 请求
}

func newRpc2AiCustomV1WenxinworkshopChatYi34bChatResult(result Rpc2AiCustomV1WenxinworkshopChatYi34bChatResponse, body []byte, http gorequest.Response) *Rpc2AiCustomV1WenxinworkshopChatYi34bChatResult {
	return &Rpc2AiCustomV1WenxinworkshopChatYi34bChatResult{Result: result, Body: body, Http: http}
}

// Rpc2AiCustomV1WenxinworkshopChatYi34bChat Yi-34B
// https://console.bce.baidu.com/qianfan/modelcenter/model/buildIn/detail/2259
// https://cloud.baidu.com/doc/WENXINWORKSHOP/s/vlpteyv3c
func (c *Client) Rpc2AiCustomV1WenxinworkshopChatYi34bChat(ctx context.Context, notMustParams ...gorequest.Params) (*Rpc2AiCustomV1WenxinworkshopChatYi34bChatResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, "rpc/2.0/ai_custom/v1/wenxinworkshop/chat/yi_34b_chat?access_token="+c.accessToken, params, http.MethodPost, "JSON")
	if err != nil {
		return newRpc2AiCustomV1WenxinworkshopChatYi34bChatResult(Rpc2AiCustomV1WenxinworkshopChatYi34bChatResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response Rpc2AiCustomV1WenxinworkshopChatYi34bChatResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRpc2AiCustomV1WenxinworkshopChatYi34bChatResult(response, request.ResponseBody, request), err
}
