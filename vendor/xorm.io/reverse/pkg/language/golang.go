// Copyright 2019 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package language

import (
	"fmt"
	"go/format"
	"html/template"
	"sort"
	"strings"

	"xorm.io/reverse/pkg/conf"
	"xorm.io/xorm/schemas"
)

func init() {
	RegisterLanguage(NewGoLanguage())
}

var (
	defaultGolangTemplate = fmt.Sprintf(`package models

{{$ilen := len .Imports}}{{if gt $ilen 0}}import (
	{{range .Imports}}"{{.}}"{{end}}
){{end}}

{{range .Tables}}
type {{TableMapper .Name}} struct {
{{$table := .}}{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	{{ColumnMapper $col.Name}}	{{Type $col}} %s{{Tag $table $col}}%s
{{end}}
}
{{end}}
`, "`", "`")
	defaultGolangTemplateTable = fmt.Sprintf(`package models

{{$ilen := len .Imports}}{{if gt $ilen 0}}import (
	{{range .Imports}}"{{.}}"{{end}}
){{end}}

{{range .Tables}}
type {{TableMapper .Name}} struct {
{{$table := .}}{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	{{ColumnMapper $col.Name}}	{{Type $col}} %s{{Tag $table $col}}%s
{{end}}
}

func (m *{{TableMapper .Name}}) TableName() string {
	return "{{$table.Name}}"
}
{{end}}
`, "`", "`")
)

// Golang represents a golang language
type GoLanguage struct {
	target *conf.ReverseTarget
	types  map[string]string
}

func NewGoLanguage() *GoLanguage {
	return &GoLanguage{
		target: nil,
		types:  make(map[string]string),
	}
}

func (g *GoLanguage) GetName() string {
	return "golang"
}

func (g *GoLanguage) GetTemplate() string {
	if g.target.TableName {
		return defaultGolangTemplateTable
	}

	return defaultGolangTemplate
}

func (g *GoLanguage) GetTypes() map[string]string {
	return g.types
}

func (g *GoLanguage) GetFuncs() template.FuncMap {
	return template.FuncMap{
		"Type": g.TypeString,
		"Tag":  g.Tag,
	}
}

func (g *GoLanguage) GetFormatter() func(string) (string, error) {
	return g.FormatGo
}

func (g *GoLanguage) GetImportter() func([]*schemas.Table) []string {
	return g.GenGoImports
}

func (g *GoLanguage) GetExtName() string {
	return ".go"
}

func (g *GoLanguage) BindTarget(target *conf.ReverseTarget) {
	g.target = target
}

func (g *GoLanguage) TypeString(col *schemas.Column) string {
	st := col.SQLType
	t := schemas.SQLType2Type(st)
	s := t.String()
	if s == "[]uint8" {
		return "[]byte"
	}
	return s
}

func (g *GoLanguage) Tag(table *schemas.Table, col *schemas.Column) template.HTML {
	isNameId := col.FieldName == "Id"
	isIdPk := isNameId && g.TypeString(col) == "int64"

	var res []string
	if !col.Nullable {
		if !isIdPk {
			res = append(res, "not null")
		}
	}
	if col.IsPrimaryKey {
		res = append(res, "pk")
	}
	if col.Default != "" {
		res = append(res, "default "+col.Default)
	}
	if col.IsAutoIncrement {
		res = append(res, "autoincr")
	}

	/*if col.SQLType.IsTime() && include(created, col.Name) {
		res = append(res, "created")
	}

	if col.SQLType.IsTime() && include(updated, col.Name) {
		res = append(res, "updated")
	}

	if col.SQLType.IsTime() && include(deleted, col.Name) {
		res = append(res, "deleted")
	}*/

	if /*supportComment &&*/ col.Comment != "" {
		res = append(res, fmt.Sprintf("comment('%s')", col.Comment))
	}

	names := make([]string, 0, len(col.Indexes))
	for name := range col.Indexes {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		index := table.Indexes[name]
		var uistr string
		if index.Type == schemas.UniqueType {
			uistr = "unique"
		} else if index.Type == schemas.IndexType {
			uistr = "index"
		}
		if len(index.Cols) > 1 {
			uistr += "(" + index.Name + ")"
		}
		res = append(res, uistr)
	}

	nstr := col.SQLType.Name
	if col.Length != 0 {
		if col.Length2 != 0 {
			nstr += fmt.Sprintf("(%v,%v)", col.Length, col.Length2)
		} else {
			nstr += fmt.Sprintf("(%v)", col.Length)
		}
	} else if len(col.EnumOptions) > 0 { //enum
		nstr += "("
		opts := ""

		enumOptions := make([]string, 0, len(col.EnumOptions))
		for enumOption := range col.EnumOptions {
			enumOptions = append(enumOptions, enumOption)
		}
		sort.Strings(enumOptions)

		for _, v := range enumOptions {
			opts += fmt.Sprintf(",'%v'", v)
		}
		nstr += strings.TrimLeft(opts, ",")
		nstr += ")"
	} else if len(col.SetOptions) > 0 { //enum
		nstr += "("
		opts := ""

		setOptions := make([]string, 0, len(col.SetOptions))
		for setOption := range col.SetOptions {
			setOptions = append(setOptions, setOption)
		}
		sort.Strings(setOptions)

		for _, v := range setOptions {
			opts += fmt.Sprintf(",'%v'", v)
		}
		nstr += strings.TrimLeft(opts, ",")
		nstr += ")"
	}
	res = append(res, nstr)
	if len(res) > 0 {
		if g.target.ColumnName {
			return template.HTML(fmt.Sprintf(`xorm:"'%s' %s"`, col.Name, strings.Join(res, " ")))
		}

		return template.HTML(fmt.Sprintf(`xorm:"%s"`, strings.Join(res, " ")))
	}
	return ""
}

func (g *GoLanguage) FormatGo(src string) (string, error) {
	source, err := format.Source([]byte(src))
	if err != nil {
		return "", err
	}
	return string(source), nil
}

func (g *GoLanguage) GenGoImports(tables []*schemas.Table) []string {
	imports := make(map[string]string)
	results := make([]string, 0)
	for _, table := range tables {
		for _, col := range table.Columns() {
			if g.TypeString(col) == "time.Time" {
				if _, ok := imports["time"]; !ok {
					imports["time"] = "time"
					results = append(results, "time")
				}
			}
		}
	}
	return results
}
