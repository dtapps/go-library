package golog

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
func EntApiLogFields(logName bool) []ent.Field {
	if logName {
		return []ent.Field{
			field.Int64("id").StorageKey("log_id").Comment("日志序号").Annotations(entsql.WithComments(true)), // 用 log_id 覆盖 框架的 id
			field.String("log_name").Immutable().Comment("日志名称").Annotations(entsql.WithComments(true)),
			field.String("trace_id").Optional().Immutable().Comment("跟踪编号").Annotations(entsql.WithComments(true)),
			field.String("request_id").Optional().Immutable().Comment("请求编号").Annotations(entsql.WithComments(true)),
			field.Time("request_time").Immutable().Comment("请求时间").Annotations(entsql.WithComments(true)),
			field.String("request_host").Optional().Immutable().Comment("请求主机").Annotations(entsql.WithComments(true)),
			field.String("request_path").Optional().Immutable().Comment("请求地址").Annotations(entsql.WithComments(true)),
			field.JSON("request_query", map[string][]string{}).Optional().Immutable().Comment("请求参数").Annotations(entsql.WithComments(true)),
			field.String("request_method").Optional().Immutable().Comment("请求方式").Annotations(entsql.WithComments(true)),
			field.JSON("request_body", map[string]any{}).Optional().Immutable().Comment("请求内容").Annotations(entsql.WithComments(true)),
			field.String("request_ip").Optional().Immutable().Comment("请求IP").Annotations(entsql.WithComments(true)),
			field.JSON("request_header", map[string][]string{}).Optional().Immutable().Comment("请求头").Annotations(entsql.WithComments(true)),
			field.Int64("request_cost_time").Optional().Immutable().Comment("请求消耗时长").Annotations(entsql.WithComments(true)),
			field.Time("response_time").Immutable().Comment("响应时间").Annotations(entsql.WithComments(true)),
			field.JSON("response_header", map[string][]string{}).Optional().Immutable().Comment("响应头").Annotations(entsql.WithComments(true)),
			field.Int("response_code").Optional().Immutable().Comment("响应状态").Annotations(entsql.WithComments(true)),
			field.JSON("response_body", map[string]any{}).Optional().Immutable().Comment("响应内容").Annotations(entsql.WithComments(true)),
		}
	} else {
		return []ent.Field{
			field.Int64("id").StorageKey("log_id").Comment("日志序号").Annotations(entsql.WithComments(true)), // 用 log_id 覆盖 框架的 id
			field.String("trace_id").Optional().Immutable().Comment("跟踪编号").Annotations(entsql.WithComments(true)),
			field.String("request_id").Optional().Immutable().Comment("请求编号").Annotations(entsql.WithComments(true)),
			field.Time("request_time").Immutable().Comment("请求时间").Annotations(entsql.WithComments(true)),
			field.String("request_host").Optional().Immutable().Comment("请求主机").Annotations(entsql.WithComments(true)),
			field.String("request_path").Optional().Immutable().Comment("请求地址").Annotations(entsql.WithComments(true)),
			field.JSON("request_query", map[string][]string{}).Optional().Immutable().Comment("请求参数").Annotations(entsql.WithComments(true)),
			field.String("request_method").Optional().Immutable().Comment("请求方式").Annotations(entsql.WithComments(true)),
			field.JSON("request_body", map[string]any{}).Optional().Immutable().Comment("请求内容").Annotations(entsql.WithComments(true)),
			field.String("request_ip").Optional().Immutable().Comment("请求IP").Annotations(entsql.WithComments(true)),
			field.JSON("request_header", map[string][]string{}).Optional().Immutable().Comment("请求头").Annotations(entsql.WithComments(true)),
			field.Int64("request_cost_time").Optional().Immutable().Comment("请求消耗时长").Annotations(entsql.WithComments(true)),
			field.Time("response_time").Immutable().Comment("响应时间").Annotations(entsql.WithComments(true)),
			field.JSON("response_header", map[string][]string{}).Optional().Immutable().Comment("响应头").Annotations(entsql.WithComments(true)),
			field.Int("response_code").Optional().Immutable().Comment("响应状态").Annotations(entsql.WithComments(true)),
			field.JSON("response_body", map[string]any{}).Optional().Immutable().Comment("响应内容").Annotations(entsql.WithComments(true)),
		}
	}
}

// EntApiLogIndexes 请求日志模型
func EntApiLogIndexes() []ent.Index {
	return []ent.Index{
		index.Fields("request_time"),
	}
}

// EntFrameLogFields 框架日志模型
func EntFrameLogFields(logName bool) []ent.Field {
	if logName {
		return []ent.Field{
			field.Int64("id").StorageKey("log_id").Comment("日志序号").Annotations(entsql.WithComments(true)), // 用 log_id 覆盖 框架的 id
			field.String("log_name").Immutable().Comment("日志名字").Annotations(entsql.WithComments(true)),
			field.String("trace_id").Optional().Immutable().Comment("跟踪编号").Annotations(entsql.WithComments(true)),
			field.String("request_id").Optional().Immutable().Comment("请求编号").Annotations(entsql.WithComments(true)),
			field.Time("request_time").Immutable().Comment("请求时间").Annotations(entsql.WithComments(true)),
			field.String("request_host").Optional().Immutable().Comment("请求主机").Annotations(entsql.WithComments(true)),
			field.String("request_path").Optional().Immutable().Comment("请求地址").Annotations(entsql.WithComments(true)),
			field.JSON("request_query", map[string]any{}).Optional().Immutable().Comment("请求参数").Annotations(entsql.WithComments(true)),
			field.String("request_method").Optional().Immutable().Comment("请求方式").Annotations(entsql.WithComments(true)),
			field.JSON("request_body", map[string]any{}).Optional().Immutable().Comment("请求内容").Annotations(entsql.WithComments(true)),
			field.String("request_ip").Optional().Immutable().Comment("请求IP").Annotations(entsql.WithComments(true)),
			field.JSON("request_header", map[string][]string{}).Optional().Immutable().Comment("请求头").Annotations(entsql.WithComments(true)),
			field.Int64("request_cost_time").Optional().Immutable().Comment("请求消耗时长").Annotations(entsql.WithComments(true)),
			field.Time("response_time").Immutable().Comment("响应时间").Annotations(entsql.WithComments(true)),
			field.JSON("response_header", map[string][]string{}).Optional().Immutable().Comment("响应头").Annotations(entsql.WithComments(true)),
			field.Int("response_code").Optional().Immutable().Comment("响应状态").Annotations(entsql.WithComments(true)),
			field.JSON("response_body", map[string]any{}).Optional().Immutable().Comment("响应内容").Annotations(entsql.WithComments(true)),
		}
	} else {
		return []ent.Field{
			field.Int64("id").StorageKey("log_id").Comment("日志序号").Annotations(entsql.WithComments(true)), // 用 log_id 覆盖 框架的 id
			field.String("log_name").Immutable().Comment("自定义类型").Annotations(entsql.WithComments(true)),
			field.String("trace_id").Optional().Immutable().Comment("跟踪编号").Annotations(entsql.WithComments(true)),
			field.String("request_id").Optional().Immutable().Comment("请求编号").Annotations(entsql.WithComments(true)),
			field.Time("request_time").Immutable().Comment("请求时间").Annotations(entsql.WithComments(true)),
			field.String("request_host").Optional().Immutable().Comment("请求主机").Annotations(entsql.WithComments(true)),
			field.String("request_path").Optional().Immutable().Comment("请求地址").Annotations(entsql.WithComments(true)),
			field.JSON("request_query", map[string]any{}).Optional().Immutable().Comment("请求参数").Annotations(entsql.WithComments(true)),
			field.String("request_method").Optional().Immutable().Comment("请求方式").Annotations(entsql.WithComments(true)),
			field.JSON("request_body", map[string]any{}).Optional().Immutable().Comment("请求内容").Annotations(entsql.WithComments(true)),
			field.String("request_ip").Optional().Immutable().Comment("请求IP").Annotations(entsql.WithComments(true)),
			field.JSON("request_header", map[string][]string{}).Optional().Immutable().Comment("请求头").Annotations(entsql.WithComments(true)),
			field.Int64("request_cost_time").Optional().Immutable().Comment("请求消耗时长").Annotations(entsql.WithComments(true)),
			field.Time("response_time").Immutable().Comment("响应时间").Annotations(entsql.WithComments(true)),
			field.JSON("response_header", map[string][]string{}).Optional().Immutable().Comment("响应头").Annotations(entsql.WithComments(true)),
			field.Int("response_code").Optional().Immutable().Comment("响应状态").Annotations(entsql.WithComments(true)),
			field.JSON("response_body", map[string]any{}).Optional().Immutable().Comment("响应内容").Annotations(entsql.WithComments(true)),
		}
	}
}

// EntFrameLogIndexes 框架日志模型
func EntFrameLogIndexes() []ent.Index {
	return []ent.Index{
		index.Fields("request_time"),
	}
}
