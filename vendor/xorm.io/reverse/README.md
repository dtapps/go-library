[中文](README_CN.md)

[![Build Status](https://drone.gitea.com/api/badges/xorm/reverse/status.svg)](https://drone.gitea.com/xorm/reverse) [![](http://gocover.io/_badge/xorm.io/xorm)](https://gocover.io/xorm.io/reverse)
[![](https://goreportcard.com/badge/xorm.io/reverse)](https://goreportcard.com/report/xorm.io/reverse)

# Reverse

A flexsible and powerful command line tool to convert database to codes.

## Installation

```
go install xorm.io/reverse
```

## Usage

```
reverse -f example/custom.yml
```

## Configuration File

How does the simplest configuration file look like?

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

A `language` defines some default configuration items, also you can define all yourselves.

```yml
kind: reverse
name: mydb
source:
  database: sqlite
  conn_str: ../testdata/test.db
targets:
- type: codes
  include_tables: # tables included, you can use **
    - a
    - b
  exclude_tables: # tables excluded, you can use **
    - c
  table_mapper: snake # how table name map to class or struct name
  column_mapper: snake # how column name map to class or struct field name
  table_prefix: "" # table prefix
  multiple_files: true # generate multiple files or one
  language: golang
  template: | # template for code file, it has higher perior than template_path
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

    func (m *{{TableMapper .Name}}) TableName() string {
    	return "{{$table.Name}}"
    }
    {{end}}
  template_path: ./template/goxorm.tmpl # template path for code file, it has higher perior than template field on language
  output_dir: ./models # code output directory
```

## Template Funcs

- *UnTitle*: Convert first charator of the word to lower.
- *Upper*: Convert word to all upper.
- *TableMapper*: Mapper method to convert table name to class/struct name.
- *ColumnMapper*: Mapper method to convert column name to class/struct field name.

## Golang Template Funcs

- *Type*: return column's golang type
- *Tag*: return golang struct tag for column

## Template Vars

- *Tables*: All tables.
- *Imports*: All imports needed.