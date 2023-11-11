// Copyright 2019 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package language

import (
	"html/template"

	"xorm.io/reverse/pkg/conf"
	"xorm.io/xorm/schemas"
)

// Language represents a languages supported when reverse codes
type Language interface {
	GetName() string
	GetTemplate() string
	GetTypes() map[string]string
	GetFuncs() template.FuncMap
	GetFormatter() func(string) (string, error)
	GetImportter() func([]*schemas.Table) []string
	GetExtName() string
	BindTarget(*conf.ReverseTarget)
}

var (
	languages = make(map[string]Language)
)

// RegisterLanguage registers a language
func RegisterLanguage(l Language) {
	languages[l.GetName()] = l
}

// GetLanguage returns a language if exists
func GetLanguage(name string, tableName bool) Language {
	language := languages[name]
	return language
}
