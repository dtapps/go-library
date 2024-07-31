package golog

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"runtime"
)

func EntAnnotations(table string, comment string) []schema.Annotation {
	if table == "" && comment == "" {
		return []schema.Annotation{}
	}
	if table != "" && comment == "" {
		return []schema.Annotation{
			entsql.Annotation{Table: table},
			entsql.WithComments(true),
		}
	}
	return []schema.Annotation{
		entsql.Annotation{Table: table},
		schema.Comment(comment),
		entsql.WithComments(true),
	}
}

// EntApiLogFields 请求日志模型
func EntApiLogFields() []ent.Field {
	return []ent.Field{
		field.String("trace_id").Optional().Immutable().Comment("跟踪编号").Annotations(entsql.WithComments(true)),
		field.String("id").StorageKey("request_id").Comment("请求编号").Annotations(entsql.WithComments(true)), // 用 request_id 覆盖 框架的 id
		field.Time("request_time").Immutable().Comment("请求时间").Annotations(entsql.WithComments(true)),
		field.String("request_host").Optional().Immutable().Comment("请求主机").Annotations(entsql.WithComments(true)),
		field.String("request_path").Optional().Immutable().Comment("请求地址").Annotations(entsql.WithComments(true)),
		field.String("request_query").Optional().Immutable().Comment("请求参数").Annotations(entsql.WithComments(true)),
		field.String("request_method").Optional().Immutable().Comment("请求方式").Annotations(entsql.WithComments(true)),
		field.String("request_scheme").Optional().Immutable().Comment("请求协议").Annotations(entsql.WithComments(true)),
		field.String("request_content_type").Optional().Immutable().Comment("请求类型").Annotations(entsql.WithComments(true)),
		field.String("request_body").Optional().Immutable().Comment("请求内容").Annotations(entsql.WithComments(true)),
		field.String("request_client_ip").Optional().Immutable().Comment("请求IP").Annotations(entsql.WithComments(true)),
		field.String("request_user_agent").Optional().Immutable().Comment("请求UA").Annotations(entsql.WithComments(true)),
		field.String("request_header").Optional().Immutable().Comment("请求头").Annotations(entsql.WithComments(true)),
		field.Int64("request_cost_time").Optional().Immutable().Comment("请求消耗时长").Annotations(entsql.WithComments(true)),
		field.Time("response_time").Optional().Immutable().Comment("响应时间").Annotations(entsql.WithComments(true)),
		field.String("response_header").Optional().Immutable().Comment("响应头").Annotations(entsql.WithComments(true)),
		field.Int("response_status_code").Optional().Immutable().Comment("响应状态").Annotations(entsql.WithComments(true)),
		field.String("response_body").Optional().Immutable().Comment("响应内容").Annotations(entsql.WithComments(true)),
		field.String("go_version").Optional().Default(runtime.Version()).Immutable().Comment("go版本").Annotations(entsql.WithComments(true)),
		field.String("sdk_version").Optional().Default(Version).Immutable().Comment("sdk版本").Annotations(entsql.WithComments(true)),
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
		index.Fields("request_time", "request_method"),
	}
}

// EntGinLogFields Gin框架日志模型
func EntGinLogFields() []ent.Field {
	return []ent.Field{
		field.String("trace_id").Optional().Immutable().Comment("跟踪编号").Annotations(entsql.WithComments(true)),
		field.String("id").StorageKey("request_id").Comment("请求编号").Annotations(entsql.WithComments(true)), // 用 request_id 覆盖 框架的 id
		field.Time("request_time").Immutable().Comment("请求时间").Annotations(entsql.WithComments(true)),
		field.String("request_host").Optional().Immutable().Comment("请求主机").Annotations(entsql.WithComments(true)),
		field.String("request_path").Optional().Immutable().Comment("请求地址").Annotations(entsql.WithComments(true)),
		field.String("request_query").Optional().Immutable().Comment("请求参数").Annotations(entsql.WithComments(true)),
		field.String("request_method").Optional().Immutable().Comment("请求方式").Annotations(entsql.WithComments(true)),
		field.String("request_scheme").Optional().Immutable().Comment("请求协议").Annotations(entsql.WithComments(true)),
		field.String("request_content_type").Optional().Immutable().Comment("请求类型").Annotations(entsql.WithComments(true)),
		field.String("request_body").Optional().Immutable().Comment("请求内容").Annotations(entsql.WithComments(true)),
		field.String("request_client_ip").Optional().Immutable().Comment("请求IP").Annotations(entsql.WithComments(true)),
		field.String("request_user_agent").Optional().Immutable().Comment("请求UA").Annotations(entsql.WithComments(true)),
		field.String("request_header").Optional().Immutable().Comment("请求头").Annotations(entsql.WithComments(true)),
		field.Int64("request_cost_time").Optional().Immutable().Comment("请求消耗时长").Annotations(entsql.WithComments(true)),
		field.Time("response_time").Optional().Immutable().Comment("响应时间").Annotations(entsql.WithComments(true)),
		field.String("response_header").Optional().Immutable().Comment("响应头").Annotations(entsql.WithComments(true)),
		field.Int("response_status_code").Optional().Immutable().Comment("响应状态").Annotations(entsql.WithComments(true)),
		field.String("response_body").Optional().Immutable().Comment("响应内容").Annotations(entsql.WithComments(true)),
		field.String("go_version").Optional().Default(runtime.Version()).Immutable().Comment("go版本").Annotations(entsql.WithComments(true)),
		field.String("sdk_version").Optional().Default(Version).Immutable().Comment("sdk版本").Annotations(entsql.WithComments(true)),
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
		index.Fields("request_time", "request_method"),
	}
}

// EntHertzLogFields Hertz框架日志模型
func EntHertzLogFields() []ent.Field {
	return []ent.Field{
		field.String("trace_id").Optional().Immutable().Comment("跟踪编号").Annotations(entsql.WithComments(true)),
		field.String("id").StorageKey("request_id").Comment("请求编号").Annotations(entsql.WithComments(true)), // 用 request_id 覆盖 框架的 id
		field.Time("request_time").Immutable().Comment("请求时间").Annotations(entsql.WithComments(true)),
		field.String("request_host").Optional().Immutable().Comment("请求主机").Annotations(entsql.WithComments(true)),
		field.String("request_path").Optional().Immutable().Comment("请求地址").Annotations(entsql.WithComments(true)),
		field.String("request_query").Optional().Immutable().Comment("请求参数").Annotations(entsql.WithComments(true)),
		field.String("request_method").Optional().Immutable().Comment("请求方式").Annotations(entsql.WithComments(true)),
		field.String("request_scheme").Optional().Immutable().Comment("请求协议").Annotations(entsql.WithComments(true)),
		field.String("request_content_type").Optional().Immutable().Comment("请求类型").Annotations(entsql.WithComments(true)),
		field.String("request_body").Optional().Immutable().Comment("请求内容").Annotations(entsql.WithComments(true)),
		field.String("request_client_ip").Optional().Immutable().Comment("请求IP").Annotations(entsql.WithComments(true)),
		field.String("request_user_agent").Optional().Immutable().Comment("请求UA").Annotations(entsql.WithComments(true)),
		field.String("request_header").Optional().Immutable().Comment("请求头").Annotations(entsql.WithComments(true)),
		field.Int64("request_cost_time").Optional().Immutable().Comment("请求消耗时长").Annotations(entsql.WithComments(true)),
		field.Time("response_time").Optional().Immutable().Comment("响应时间").Annotations(entsql.WithComments(true)),
		field.String("response_header").Optional().Immutable().Comment("响应头").Annotations(entsql.WithComments(true)),
		field.Int("response_status_code").Optional().Immutable().Comment("响应状态").Annotations(entsql.WithComments(true)),
		field.String("response_body").Optional().Immutable().Comment("响应内容").Annotations(entsql.WithComments(true)),
		field.String("go_version").Optional().Default(runtime.Version()).Immutable().Comment("go版本").Annotations(entsql.WithComments(true)),
		field.String("sdk_version").Optional().Default(Version).Immutable().Comment("sdk版本").Annotations(entsql.WithComments(true)),
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
		index.Fields("request_time", "request_method"),
	}
}
