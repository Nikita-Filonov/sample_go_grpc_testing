package solutions

import "reflect"

func isType(value reflect.Value, kindOf reflect.Kind) bool {
	return value.Type().Kind() == kindOf
}

func isPointer(value reflect.Value) bool {
	return isType(value, reflect.Pointer)
}

func isInt(value reflect.Value) bool {
	return isType(value, reflect.Int)
}

func isInt8(value reflect.Value) bool {
	return isType(value, reflect.Int8)
}

func isInt32(value reflect.Value) bool {
	return isType(value, reflect.Int32)
}

func isInt64(value reflect.Value) bool {
	return isType(value, reflect.Int64)
}

func isFloat32(value reflect.Value) bool {
	return isType(value, reflect.Float32)
}

func isFloat64(value reflect.Value) bool {
	return isType(value, reflect.Float64)
}

func isBool(value reflect.Value) bool {
	return isType(value, reflect.Bool)
}

func isString(value reflect.Value) bool {
	return isType(value, reflect.String)
}
