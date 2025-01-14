package gojobs

type TaskCustomHelperTaskList struct {
	TaskID        string `json:"task_id,omitempty"`        // 任务编号
	TaskName      string `json:"task_name,omitempty"`      // 任务名称
	TaskParams    string `json:"task_params,omitempty"`    // 任务参数
	CustomID      string `json:"custom_id,omitempty"`      // 自定义编号
	CustomContent string `json:"custom_content,omitempty"` // 自定义内容
}
