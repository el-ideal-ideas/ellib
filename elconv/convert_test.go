package elconv

import (
	"reflect"
	"testing"
)

func tMap(arr []interface{}, f func(interface{}) interface{}) []interface{} {
	res := make([]interface{}, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[i] = f(arr[i])
	}
	return res
}

func tIsArrayOrSlice(v interface{}) bool {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Array, reflect.Slice:
		return true
	default:
		return false
	}
}

func tIsSame(data ...interface{}) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check type
	for i := 0; i < count; i++ {
		if !tIsArrayOrSlice(data[i]) {
			return false
		}
	}
	// check length
	length := reflect.ValueOf(data[0]).Len()
	for i := 1; i < count; i++ {
		if reflect.ValueOf(data[i]).Len() != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		v := reflect.ValueOf(data[0])
		first := v.Index(i).Interface()
		for j := 1; j < count; j++ {
			switch first.(type) {
			case int, int8, int16, int32, int64:
				if AsInt64(first) != AsInt64(reflect.ValueOf(data[j]).Index(i).Interface()) {
					return false
				}
			case uint, uint8, uint16, uint32, uint64:
				if AsUint64(first) != AsUint64(reflect.ValueOf(data[j]).Index(i).Interface()) {
					return false
				}
			case float32:
				if AsFloat32(first) != AsFloat32(reflect.ValueOf(data[j]).Index(i).Interface()) {
					return false
				}
			case float64:
				if AsFloat64(first) != AsFloat64(reflect.ValueOf(data[j]).Index(i).Interface()) {
					return false
				}
			case string:
				if AsStr(first) != AsStr(reflect.ValueOf(data[j]).Index(i).Interface()) {
					return false
				}
			default:
				if !reflect.DeepEqual(first, reflect.ValueOf(data[j]).Index(i).Interface()) {
					return false
				}
			}
		}
	}
	return true
}

func TestConv(t *testing.T) {
	data := []interface{}{1, 2.3, "3", -8, "true", true}
	boolWant := []interface{}{true, true, false, false, true, true}
	floatWant := []interface{}{1, 2.3, 3, -8, 0, 1}
	intWant := []interface{}{1, 2, 3, -8, 0, 1}
	uintWant := []interface{}{1, 2, 3, 0, 0, 1}
	strWant := []interface{}{"1", "2.3", "3", "-8", "true", "true"}
	if res := tMap(data, func(v interface{}) interface{} { return AsBool(v) }); !tIsSame(res, boolWant) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := tMap(data, func(v interface{}) interface{} { return AsFloat64(v) }); !tIsSame(res, floatWant) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := tMap(data, func(v interface{}) interface{} { return AsFloat32(v) }); !tIsSame(res, floatWant) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := tMap(data, func(v interface{}) interface{} { return AsInt(v) }); !tIsSame(res, intWant) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := tMap(data, func(v interface{}) interface{} { return AsInt64(v) }); !tIsSame(res, intWant) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := tMap(data, func(v interface{}) interface{} { return AsUint(v) }); !tIsSame(res, uintWant) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := tMap(data, func(v interface{}) interface{} { return AsUint64(v) }); !tIsSame(res, uintWant) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := tMap(data, func(v interface{}) interface{} { return AsStr(v) }); !tIsSame(res, strWant) {
		t.Errorf("Invalid result: %v", res)
	}
}
