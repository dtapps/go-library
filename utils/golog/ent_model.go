package golog

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// EntApiLogFields 请求日志模型
func EntApiLogFields() []ent.Field {
	return []ent.Field{
		field.String("trace_id").Optional().Comment("跟踪编号"),

		field.String("id").StorageKey("request_id").Comment("请求编号"), // 用 request_id 覆盖 框架的 id
		field.Time("request_time").Comment("请求时间"),
		field.String("request_host").Optional().Comment("请求主机"),
		field.String("request_path").Optional().Comment("请求地址"),
		field.String("request_query").Optional().Comment("请求参数"),
		field.String("request_method").Optional().Comment("请求方式"),
		field.String("request_scheme").Optional().Comment("请求协议"),
		field.String("request_content_type").Optional().Comment("请求类型"),
		field.String("request_body").Optional().Comment("请求内容"),
		field.String("request_client_ip").Optional().Comment("请求IP"),
		field.String("request_user_agent").Optional().Comment("请求UA"),
		field.String("request_header").Optional().Comment("请求头"),
		field.Int64("request_cost_time").Optional().Comment("请求消耗时长"),

		field.Time("response_time").Optional().Comment("响应时间"),
		field.String("response_header").Optional().Comment("响应头"),
		field.Int("response_status_code").Optional().Comment("响应状态"),
		field.String("response_body").Optional().Comment("响应内容"),

		field.String("go_version").Optional().Comment("go版本"),
		field.String("sdk_version").Optional().Comment("sdk版本"),
	}
}

// EntApiLogIndexes 请求日志模型
func EntApiLogIndexes() []ent.Index {
	return []ent.Index{
		index.Fields("trace_id"),
		//index.Fields("request_id"),
		index.Fields("request_time"),
		index.Fields("request_path"),
		index.Fields("request_method"),
		index.Fields("response_time"),
	}
}

// EntGinLogFields Gin框架日志模型
func EntGinLogFields() []ent.Field {
	return []ent.Field{
		field.String("trace_id").Optional().Comment("跟踪编号"),

		field.String("id").StorageKey("request_id").Comment("请求编号"), // 用 request_id 覆盖 框架的 id
		field.Time("request_time").Comment("请求时间"),
		field.String("request_host").Optional().Comment("请求主机"),
		field.String("request_path").Optional().Comment("请求地址"),
		field.String("request_query").Optional().Comment("请求参数"),
		field.String("request_method").Optional().Comment("请求方式"),
		field.String("request_scheme").Optional().Comment("请求协议"),
		field.String("request_content_type").Optional().Comment("请求类型"),
		field.String("request_body").Optional().Comment("请求内容"),
		field.String("request_client_ip").Optional().Comment("请求IP"),
		field.String("request_user_agent").Optional().Comment("请求UA"),
		field.String("request_header").Optional().Comment("请求头"),
		field.Int64("request_cost_time").Optional().Comment("请求消耗时长"),

		field.Time("response_time").Optional().Comment("响应时间"),
		field.String("response_header").Optional().Comment("响应头"),
		field.Int("response_status_code").Optional().Comment("响应状态"),
		field.String("response_body").Optional().Comment("响应内容"),

		field.String("go_version").Optional().Comment("go版本"),
		field.String("sdk_version").Optional().Comment("sdk版本"),
	}
}

// EntGinLogIndexes Gin框架日志模型
func EntGinLogIndexes() []ent.Index {
	return []ent.Index{
		index.Fields("trace_id"),
		//index.Fields("request_id"),
		index.Fields("request_time"),
		index.Fields("request_path"),
		index.Fields("request_method"),
		index.Fields("response_time"),
	}
}

// EntHertzLogFields Hertz框架日志模型
func EntHertzLogFields() []ent.Field {
	return []ent.Field{
		field.String("trace_id").Optional().Comment("跟踪编号"),

		field.String("id").StorageKey("request_id").Comment("请求编号"), // 用 request_id 覆盖 框架的 id
		field.Time("request_time").Comment("请求时间"),
		field.String("request_host").Optional().Comment("请求主机"),
		field.String("request_path").Optional().Comment("请求地址"),
		field.String("request_query").Optional().Comment("请求参数"),
		field.String("request_method").Optional().Comment("请求方式"),
		field.String("request_scheme").Optional().Comment("请求协议"),
		field.String("request_content_type").Optional().Comment("请求类型"),
		field.String("request_body").Optional().Comment("请求内容"),
		field.String("request_client_ip").Optional().Comment("请求IP"),
		field.String("request_user_agent").Optional().Comment("请求UA"),
		field.String("request_header").Optional().Comment("请求头"),
		field.Int64("request_cost_time").Optional().Comment("请求消耗时长"),

		field.Time("response_time").Optional().Comment("响应时间"),
		field.String("response_header").Optional().Comment("响应头"),
		field.Int("response_status_code").Optional().Comment("响应状态"),
		field.String("response_body").Optional().Comment("响应内容"),

		field.String("go_version").Optional().Comment("go版本"),
		field.String("sdk_version").Optional().Comment("sdk版本"),
	}
}

// EntHertzLogIndexes Hertz框架日志模型
func EntHertzLogIndexes() []ent.Index {
	return []ent.Index{
		index.Fields("trace_id"),
		//index.Fields("request_id"),
		index.Fields("request_time"),
		index.Fields("request_path"),
		index.Fields("request_method"),
		index.Fields("response_time"),
	}
}
