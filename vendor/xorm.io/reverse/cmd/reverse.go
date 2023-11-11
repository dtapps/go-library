// Copyright 2019 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"bytes"
	"errors"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"xorm.io/reverse/pkg/conf"
	"xorm.io/reverse/pkg/language"
	"xorm.io/reverse/pkg/utils"

	"gitea.com/lunny/log"
	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/gobwas/glob"
	"xorm.io/xorm"
	"xorm.io/xorm/schemas"
)

var (
	defaultFuncs = template.FuncMap{
		"UnTitle": utils.UnTitle,
		"Upper":   utils.UpTitle,
	}
)

func reverseFromConfig(rFile string) error {
	configs, err := conf.NewReverseConfigFromYAML(rFile)
	if err != nil {
		return err
	}

	for _, cfg := range configs {
		for _, target := range cfg.Targets {
			if err := runReverse(&cfg.Source, &target); err != nil {
				return err
			}
		}
	}

	return nil
}

// filterTables filter by target.ExcludeTables and target.IncludeTables
func filterTables(tables []*schemas.Table, target *conf.ReverseTarget) []*schemas.Table {
	var res = make([]*schemas.Table, 0, len(tables))
	underscore.Chain(tables).
		Filter(func(tbl schemas.Table, _ int) bool {
			for _, exclude := range target.ExcludeTables {
				s, _ := glob.Compile(exclude)
				if s.Match(tbl.Name) {
					return false
				}
			}

			return true
		}).
		Filter(func(tbl schemas.Table, _ int) bool {
			// if not set, all tables by default
			if len(target.IncludeTables) == 0 {
				return true
			}

			for _, include := range target.IncludeTables {
				s, _ := glob.Compile(include)
				if s.Match(tbl.Name) {
					return true
				}
			}

			return false
		}).
		Each(func(tbl schemas.Table, _ int) {
			res = append(res, &tbl)
		})

	return res
}

func runReverse(source *conf.ReverseSource, target *conf.ReverseTarget) error {
	var (
		formatter func(string) (string, error)
		importter func([]*schemas.Table) []string
	)

	orm, err := xorm.NewEngine(source.Database, source.ConnStr)
	if err != nil {
		return err
	}

	tables, err := orm.DBMetas()
	if err != nil {
		return err
	}

	// filter tables according includes and excludes
	tables = filterTables(tables, target)

	// load configuration from language
	lang := language.GetLanguage(target.Language, target.TableName)

	// load template
	var bs []byte
	if target.Template != "" {
		bs = []byte(target.Template)
	} else if target.TemplatePath != "" {
		bs, err = ioutil.ReadFile(target.TemplatePath)
		if err != nil {
			return err
		}
	}

	var tableMapper = utils.GetMapperByName(target.TableMapper)
	var colMapper = utils.GetMapperByName(target.ColumnMapper)
	funcs := utils.MergeFuncMap(
		template.FuncMap(defaultFuncs),
		template.FuncMap{
			"TableMapper":  tableMapper.Table2Obj,
			"ColumnMapper": colMapper.Table2Obj,
		})

	if lang != nil {
		lang.BindTarget(target)

		if bs == nil {
			bs = []byte(lang.GetTemplate())
		}

		funcs = utils.MergeFuncMap(funcs, lang.GetFuncs())

		if formatter == nil {
			formatter = lang.GetFormatter()
		}

		if importter == nil {
			importter = lang.GetImportter()
		}

		target.ExtName = lang.GetExtName()
	}
	if !strings.HasPrefix(target.ExtName, ".") {
		target.ExtName = "." + target.ExtName
	}

	if bs == nil {
		return errors.New("you have to indicate template / template path or a language")
	}

	t := template.New("reverse")
	t.Funcs(funcs)

	tmpl, err := t.Parse(string(bs))
	if err != nil {
		return err
	}

	for _, table := range tables {
		if target.TablePrefix != "" {
			table.Name = strings.TrimPrefix(table.Name, target.TablePrefix)
		}
		for _, col := range table.Columns() {
			col.FieldName = colMapper.Table2Obj(col.Name)
		}
	}

	err = os.MkdirAll(target.OutputDir, os.ModePerm)
	if err != nil {
		return err
	}

	var w *os.File
	if !target.MultipleFiles {
		w, err = os.Create(filepath.Join(target.OutputDir, "models"+target.ExtName))
		if err != nil {
			return err
		}
		defer w.Close()

		imports := importter(tables)

		newbytes := bytes.NewBufferString("")
		err = tmpl.Execute(newbytes, map[string]interface{}{
			"Tables":  tables,
			"Imports": imports,
		})
		if err != nil {
			return err
		}

		tplcontent, err := ioutil.ReadAll(newbytes)
		if err != nil {
			return err
		}
		var source string
		if formatter != nil {
			source, err = formatter(string(tplcontent))
			if err != nil {
				log.Warnf("%v", err)
				source = string(tplcontent)
			}
		} else {
			source = string(tplcontent)
		}

		w.WriteString(source)
		w.Close()
	} else {
		for _, table := range tables {
			// imports
			tbs := []*schemas.Table{table}
			imports := importter(tbs)

			w, err := os.Create(filepath.Join(target.OutputDir, table.Name+target.ExtName))
			if err != nil {
				return err
			}
			defer w.Close()

			newbytes := bytes.NewBufferString("")
			err = tmpl.Execute(newbytes, map[string]interface{}{
				"Tables":  tbs,
				"Imports": imports,
			})
			if err != nil {
				return err
			}

			tplcontent, err := ioutil.ReadAll(newbytes)
			if err != nil {
				return err
			}
			var source string
			if formatter != nil {
				source, err = formatter(string(tplcontent))
				if err != nil {
					log.Warnf("%v", err)
					source = string(tplcontent)
				}
			} else {
				source = string(tplcontent)
			}

			w.WriteString(source)
			w.Close()
		}
	}

	return nil
}
