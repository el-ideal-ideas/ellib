package elstr

import (
	"errors"
	"reflect"
	"strconv"
)

var (
	ErrConvertFail = errors.New("convert data type is failure")
)

// MustString convert value to string
func MustString(v interface{}) string {
	val, _ := ToString(v)
	return val
}

// ToString convert value to string
func ToString(v interface{}) (str string, err error) {
	if v == nil {
		return
	}
	r := reflect.ValueOf(v)
	if r.Kind() == reflect.Ptr {
		v = r.Elem().Interface()
	}
	switch value := v.(type) {
	case int:
		str = strconv.Itoa(value)
	case int8:
		str = strconv.Itoa(int(value))
	case int16:
		str = strconv.Itoa(int(value))
	case int32:
		str = strconv.Itoa(int(value))
	case int64:
		str = strconv.Itoa(int(value))
	case uint:
		str = strconv.FormatUint(uint64(value), 10)
	case uint8:
		str = strconv.FormatUint(uint64(value), 10)
	case uint16:
		str = strconv.FormatUint(uint64(value), 10)
	case uint32:
		str = strconv.FormatUint(uint64(value), 10)
	case uint64:
		str = strconv.FormatUint(value, 10)
	case float32:
		str = strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		str = strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		str = strconv.FormatBool(value)
	case string:
		str = value
	case []byte:
		str = string(value)
	default:
		err = ErrConvertFail
	}
	return
}
