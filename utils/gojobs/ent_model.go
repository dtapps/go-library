package gojobs

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"runtime"
	"time"
)

// EntTaskAnnotations 任务
func EntTaskAnnotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "task"},
		schema.Comment("系统任务"),
		entsql.WithComments(true),
	}
}

// EntTaskFields 任务
func EntTaskFields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment("记录编号").Annotations(entsql.WithComments(true)),
		field.String("status").Optional().Comment("状态码").Annotations(entsql.WithComments(true)),
		field.String("status_desc").Optional().Comment("状态描述").Annotations(entsql.WithComments(true)),
		field.String("spec").NotEmpty().Immutable().Comment("任务规则").Annotations(entsql.WithComments(true)),
		field.String("params").Optional().Comment("参数").Annotations(entsql.WithComments(true)),
		field.Int64("frequency").Optional().Immutable().Comment("频率(秒单位)").Annotations(entsql.WithComments(true)),
		field.Int64("number").Optional().Default(0).Comment("当前次数").Annotations(entsql.WithComments(true)),
		field.Int64("max_number").Optional().Comment("最大次数").Annotations(entsql.WithComments(true)),
		field.String("run_id").Optional().Comment("执行编号").Annotations(entsql.WithComments(true)),
		field.String("custom_id").Optional().Immutable().Comment("自定义编号").Annotations(entsql.WithComments(true)),
		field.Int64("custom_sequence").Optional().Comment("自定义顺序").Annotations(entsql.WithComments(true)),
		field.String("type").NotEmpty().Immutable().Comment("类型").Annotations(entsql.WithComments(true)),
		field.String("type_name").Optional().Immutable().Comment("类型名称").Annotations(entsql.WithComments(true)),
		field.String("created_ip").Default("0.0.0.0").Immutable().Comment("创建外网IP").Annotations(entsql.WithComments(true)),
		field.String("specify_ip").Default("0.0.0.0").Immutable().Comment("指定外网IP").Annotations(entsql.WithComments(true)),
		field.String("updated_ip").Default("0.0.0.0").Comment("更新外网IP").Annotations(entsql.WithComments(true)),
		field.String("result").Optional().Comment("结果").Annotations(entsql.WithComments(true)),
		field.Time("next_run_time").Optional().Comment("下次运行时间").Annotations(entsql.WithComments(true)),
		field.Time("created_at").Default(time.Now).Immutable().Comment("创建时间").Annotations(entsql.WithComments(true)),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间").Annotations(entsql.WithComments(true)),
	}
}

// EntTaskIndexes 任务
func EntTaskIndexes() []ent.Index {
	return []ent.Index{
		index.Fields("type"),
		index.Fields("status"),
		index.Fields("type", "status"),
		index.Fields("spec"),
		index.Fields("spec", "status"),
	}
}

// EntTaskLogAnnotations 任务日志
func EntTaskLogAnnotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "task_log"},
		schema.Comment("系统任务日志"),
		entsql.WithComments(true),
	}
}

// EntTaskLogFields 任务日志
func EntTaskLogFields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StorageKey("log_id").Comment("日志编号").Annotations(entsql.WithComments(true)), // 用 log_id 覆盖 框架的 id
		field.Time("log_time").Default(time.Now).Immutable().Comment("日志时间").Annotations(entsql.WithComments(true)),
		field.Int64("task_id").Immutable().Comment("任务编号").Annotations(entsql.WithComments(true)),
		field.String("task_run_id").Optional().Immutable().Comment("执行编号").Annotations(entsql.WithComments(true)),
		field.Int("task_result_code").Optional().Immutable().Comment("执行状态码").Annotations(entsql.WithComments(true)),
		field.String("task_result_desc").Optional().Immutable().Comment("执行结果").Annotations(entsql.WithComments(true)),
		field.Float("task_cost_time").Default(0).Immutable().Comment("消耗时长").Annotations(entsql.WithComments(true)),
		field.String("system_inside_ip").Default("0.0.0.0").Immutable().Comment("内网IP").Annotations(entsql.WithComments(true)),
		field.String("system_outside_ip").Default("0.0.0.0").Immutable().Comment("外网IP").Annotations(entsql.WithComments(true)),
		field.String("go_version").Default(runtime.Version()).Immutable().Comment("go版本").Annotations(entsql.WithComments(true)),
		field.String("sdk_version").Default(Version).Immutable().Comment("sdk版本").Annotations(entsql.WithComments(true)),
	}
}

// EntTaskLogIndexes 任务日志
func EntTaskLogIndexes() []ent.Index {
	return []ent.Index{
		index.Fields("task_id"),
	}
}
