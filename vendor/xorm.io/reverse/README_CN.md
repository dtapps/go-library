[English](README.md)

[![Build Status](https://drone.gitea.com/api/badges/xorm/reverse/status.svg)](https://drone.gitea.com/xorm/reverse) [![](http://gocover.io/_badge/xorm.io/xorm)](https://gocover.io/xorm.io/reverse)
[![](https://goreportcard.com/badge/xorm.io/reverse)](https://goreportcard.com/report/xorm.io/reverse)

# Reverse

一个灵活高效的数据库反转工具。

## 安装

```
go get xorm.io/reverse
```

## 使用

```
reverse -f example/custom.yml
```

## 配置文件

一个最简单的配置文件看起来如下：

```yml
kind: reverse
name: mydb
source:
  database: sqlite3
  conn_str: '../testdata/test.db'
targets:
- type: codes
  language: golang
  output_dir: ../models
```

`language` 定义了很多默认的配置，你也可以自己来进行配置。其中的模板是 Go 模板语法。

```yml
kind: reverse
name: mydb
source:
  database: sqlite
  conn_str: ../testdata/test.db
targets:
- type: codes
  include_tables: # 包含的表，以下可以用 **
    - a
    - b
  exclude_tables: # 排除的表，以下可以用 **
    - c
  table_mapper: snake # 表名到代码类或结构体的映射关系
  column_mapper: snake # 字段名到代码或结构体成员的映射关系
  table_prefix: "" # 表前缀
  multiple_files: true # 是否生成多个文件
  language: golang
  template: | # 生成模板，如果这里定义了，优先级比 template_path 高
    package models

    {{$ilen := len .Imports}}
    {{if gt $ilen 0}}
    import (
      {{range .Imports}}"{{.}}"{{end}}
    )
    {{end}}

    {{range .Tables}}
    type {{TableMapper .Name}} struct {
    {{$table := .}}
    {{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	{{ColumnMapper $col.Name}}	{{Type $col}} `{{Tag $table $col}}`
    {{end}}
    }
    {{end}}
  template_path: ./template/goxorm.tmpl # 生成的模板的路径，优先级比 template 低，但比 language 中的默认模板高
  output_dir: ./models # 代码生成目录
```

## 模板函数

- *UnTitle*: 将单词的第一个字母大写。
- *Upper*: 将单词转为全部大写。
- *TableMapper*: 将表名转为结构体名的映射函数。
- *ColumnMapper*: 将字段名转为结构体成员名的函数。

## Go 语言模版函数

- *Type*: 返回Go语言的类型
- *Tag*: 返回Go语言的Tag信息

## 模板变量

- *Tables*: 所有表。
- *Imports*: 所有需要的导入。