package solutions

import (
	"reflect"
)

func PrettifyValue(input interface{}) string {
	valueOf := reflect.ValueOf(input)

	if isPointer(valueOf) {
		if valueOf.IsNil() {
			return "null"
		}

		valueOf = reflect.Indirect(valueOf)
	}

	template := getTemplate(valueOf)
	return template.format()
}
