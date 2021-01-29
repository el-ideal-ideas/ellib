package elconv

import "strconv"

// convert value to int
func AsInt(v interface{}) int {
	switch v.(type) {
	case int:
		return v.(int)
	case int8:
		return int(v.(int8))
	case int16:
		return int(v.(int16))
	case int32:
		return int(v.(int32))
	case int64:
		return int(v.(int64))
	case uint:
		return int(v.(uint))
	case uint8:
		return int(v.(uint8))
	case uint16:
		return int(v.(uint16))
	case uint32:
		return int(v.(uint32))
	case uint64:
		return int(v.(uint64))
	case float32:
		return int(v.(float32))
	case float64:
		return int(v.(float64))
	case []byte:
		f, e := strconv.ParseFloat(string(v.([]byte)), 64)
		if e == nil {
			return int(f)
		} else {
			return 0
		}
	case string:
		f, e := strconv.ParseFloat(v.(string), 64)
		if e == nil {
			return int(f)
		} else {
			return 0
		}
	default:
		return 0
	}
}

// convert value to int64
func AsInt64(v interface{}) int64 {
	switch v.(type) {
	case int:
		return int64(v.(int))
	case int8:
		return int64(v.(int8))
	case int16:
		return int64(v.(int16))
	case int32:
		return int64(v.(int32))
	case int64:
		return v.(int64)
	case uint:
		return int64(v.(uint))
	case uint8:
		return int64(v.(uint8))
	case uint16:
		return int64(v.(uint16))
	case uint32:
		return int64(v.(uint32))
	case uint64:
		return int64(v.(uint64))
	case float32:
		return int64(v.(float32))
	case float64:
		return int64(v.(float64))
	case []byte:
		f, e := strconv.ParseFloat(string(v.([]byte)), 64)
		if e == nil {
			return int64(f)
		} else {
			return 0
		}
	case string:
		f, e := strconv.ParseFloat(v.(string), 64)
		if e == nil {
			return int64(f)
		} else {
			return 0
		}
	default:
		return 0
	}
}
