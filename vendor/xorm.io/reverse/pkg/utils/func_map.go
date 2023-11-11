package utils

import "html/template"

func MergeFuncMap(funcMaps ...template.FuncMap) template.FuncMap {
	result := template.FuncMap{}

	for _, m := range funcMaps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}
