package api2

import (
	"reflect"
	"strings"
)

var fieldsByTag = make(map[reflect.Type]map[string]string)
var tagsByField = make(map[reflect.Type]map[string]string)

func buildFieldsToTagsMapping(s interface{}) {
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		panic("bad type")
	}

	if fieldsByTag[rt] == nil {
		fieldsByTag[rt] = make(map[string]string)
	}

	if tagsByField[rt] == nil {
		tagsByField[rt] = make(map[string]string)
	}

	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		v := strings.Split(f.Tag.Get("json"), ",")[0] // use split to ignore tag "options"
		if v == "" || v == "-" {
			continue
		}
		fieldsByTag[rt][v] = f.Name
		tagsByField[rt][f.Name] = v
	}
}

func init() {
	buildFieldsToTagsMapping(Storage{})
}

func getFieldNameByTag(tag string, s interface{}) (fieldname string) {
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		panic("bad type")
	}
	return fieldsByTag[rt][tag]
}

func getTagByFieldName(fieldname string, s interface{}) (tag string) {
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		panic("bad type")
	}
	return tagsByField[rt][fieldname]
}
