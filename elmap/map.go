package elmap

import (
	"fmt"
	"github.com/el-ideal-ideas/ellib/elconv"
	"reflect"
)

// Flip key and value in the map.
// If `m` is not a map that will cause panic.
func MapFlip(m interface{}) map[interface{}]interface{} {
	res := make(map[interface{}]interface{})
	val := reflect.ValueOf(m)
	switch val.Kind() {
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if val.MapIndex(k).Interface() != nil && fmt.Sprintf("%v", val.MapIndex(k).Interface()) != "" {
				res[val.MapIndex(k).Interface()] = k
			}
		}
	default:
		panic("[MapFlip]m type must be map")
	}
	return res
}

// Get all keys in the map.
// If `m` is not a map that will cause panic.
func Keys(m interface{}) []interface{} {
	val := reflect.ValueOf(m)
	res := make([]interface{}, val.Len())
	switch val.Kind() {
	case reflect.Map:
		for i, k := range val.MapKeys() {
			res[i] = k
		}
	default:
		panic("[MapKeys]m type must be map")
	}
	return res
}

// Get all keys in the map as []string.
// If `m` is not a map that will cause panic.
func KeysAsString(m interface{}) []string {
	val := reflect.ValueOf(m)
	res := make([]string, val.Len())
	switch val.Kind() {
	case reflect.Map:
		for i, k := range val.MapKeys() {
			res[i] = elconv.AsStr(k.Interface())
		}
	default:
		panic("[MapKeys]m type must be map")
	}
	return res
}

// Get all values in the map
// If `m` is not a map that will cause panic.
func GetMapValues(m interface{}) []interface{} {
	var res []interface{}
	var item interface{}
	val := reflect.ValueOf(m)
	switch val.Kind() {
	case reflect.Map:
		for _, k := range val.MapKeys() {
			item = val.MapIndex(k).Interface()
			res = append(res, item)
		}
	default:
		panic("[GetMapValues]m type must be map")
	}
	return res
}
