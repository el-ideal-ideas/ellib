package elconv

import "strconv"

// convert value to float64
func AsFloat(v interface{}) (res float64) {
	switch v.(type) {
	case int:
		res = float64(v.(int))
	case int8:
		res = float64(v.(int8))
	case int16:
		res = float64(v.(int16))
	case int32:
		res = float64(v.(int32))
	case int64:
		res = float64(v.(int64))
	case uint:
		res = float64(v.(uint))
	case uint8:
		res = float64(v.(uint8))
	case uint16:
		res = float64(v.(uint16))
	case uint32:
		res = float64(v.(uint32))
	case uint64:
		res = float64(v.(uint64))
	case float32:
		res = float64(v.(float32))
	case float64:
		res = v.(float64)
	case []uint8:
		f, err := strconv.ParseFloat(string(v.([]uint8)), 64)
		if err == nil {
			res = f
		} else {
			res = 0.0
		}
	case string:
		f, err := strconv.ParseFloat(v.(string), 64)
		if err == nil {
			res = f
		} else {
			res = 0.0
		}
	case bool:
		if v.(bool) {
			res = 1.0
		} else {
			res = 0.0
		}
	default:
		res = 0.0
	}
	return
}
