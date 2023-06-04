package solutions

import (
	"fmt"
	"reflect"
	"strconv"
)

type Template struct {
	name  string
	value string
}

func (t Template) format() string {
	return fmt.Sprintf("%s[%s]", t.name, t.value)
}

func getTemplate(value reflect.Value) Template {
	if isInt(value) {
		return Template{value: strconv.FormatInt(value.Int(), 10), name: "int"}
	}

	if isInt8(value) {
		return Template{value: strconv.FormatInt(value.Int(), 10), name: "int8"}
	}

	if isInt32(value) {
		return Template{value: strconv.FormatInt(value.Int(), 10), name: "int32"}
	}

	if isInt64(value) {
		return Template{value: strconv.FormatInt(value.Int(), 10), name: "int64"}
	}

	if isFloat32(value) {
		return Template{value: strconv.FormatFloat(value.Float(), 'E', -1, 64), name: "float32"}
	}

	if isFloat64(value) {
		return Template{value: strconv.FormatFloat(value.Float(), 'E', -1, 64), name: "float64"}
	}

	if isBool(value) {
		return Template{value: strconv.FormatBool(value.Bool()), name: "bool"}
	}

	if isString(value) {
		return Template{value: value.String(), name: "string"}
	}

	return Template{}
}
