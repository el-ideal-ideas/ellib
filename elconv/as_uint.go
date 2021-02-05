package elconv

import (
	"reflect"
	"strconv"
)

// convert value to uint
func AsUint(v interface{}) uint {
	if v == nil {
		return 0
	}
	v = AsValueRef(reflect.ValueOf(v)).Interface()
	switch v.(type) {
	case int, int8, int16, int32, int64, float32, float64:
		if AsFloat64(v) <= 0 {
			return 0
		}
	}
	switch v.(type) {
	case int:
		return uint(v.(int))
	case int8:
		return uint(v.(int8))
	case int16:
		return uint(v.(int16))
	case int32:
		return uint(v.(int32))
	case int64:
		return uint(v.(int64))
	case uint:
		return v.(uint)
	case uint8:
		return uint(v.(uint8))
	case uint16:
		return uint(v.(uint16))
	case uint32:
		return uint(v.(uint32))
	case uint64:
		return uint(v.(uint64))
	case float32:
		return uint(v.(float32))
	case float64:
		return uint(v.(float64))
	case []byte:
		f, e := strconv.ParseFloat(string(v.([]byte)), 64)
		if e == nil {
			return uint(f)
		} else {
			return 0
		}
	case string:
		f, e := strconv.ParseFloat(v.(string), 64)
		if e == nil {
			return uint(f)
		} else {
			return 0
		}
	case bool:
		if v.(bool) {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
}

// convert value to uint64
func AsUint64(v interface{}) uint64 {
	switch v.(type) {
	case int, int8, int16, int32, int64, float32, float64:
		if AsFloat64(v) <= 0 {
			return 0
		}
	}
	switch v.(type) {
	case int:
		return uint64(v.(int))
	case int8:
		return uint64(v.(int8))
	case int16:
		return uint64(v.(int16))
	case int32:
		return uint64(v.(int32))
	case int64:
		return uint64(v.(int64))
	case uint:
		return uint64(v.(uint))
	case uint8:
		return uint64(v.(uint8))
	case uint16:
		return uint64(v.(uint16))
	case uint32:
		return uint64(v.(uint32))
	case uint64:
		return v.(uint64)
	case float32:
		return uint64(v.(float32))
	case float64:
		return uint64(v.(float64))
	case []byte:
		f, e := strconv.ParseFloat(string(v.([]byte)), 64)
		if e == nil {
			return uint64(f)
		} else {
			return 0
		}
	case string:
		f, e := strconv.ParseFloat(v.(string), 64)
		if e == nil {
			return uint64(f)
		} else {
			return 0
		}
	case bool:
		if v.(bool) {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
}
