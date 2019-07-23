package cosmos

import (
	"fmt"
	"strconv"
)

func ToString(i interface{}) string {
	if i == nil {
		return ""
	}
	return fmt.Sprint(i)
}

func StringConv(s string, defaultValue interface{}) interface{} {
	var err error
	var value interface{}
	switch defaultValue.(type) {
	case bool:
		value, err = strconv.ParseBool(s)
	case int:
		value, err = strconv.ParseInt(s, 10, 0)
		value = int(value.(int64))
	case int8:
		value, err = strconv.ParseInt(s, 10, 8)
		value = int8(value.(int64))
	case int16:
		value, err = strconv.ParseInt(s, 10, 16)
		value = int16(value.(int64))
	case int32:
		value, err = strconv.ParseInt(s, 10, 32)
		value = int8(value.(int64))
	case int64:
		value, err = strconv.ParseInt(s, 10, 64)
	case float64:
		value, err = strconv.ParseFloat(s, 64)
	case string:
		value = s
		if value == "" {
			value = defaultValue.(string)
		}
	default:
		panic(fmt.Sprintf("nonexistent value type:%v", defaultValue))
	}
	if err != nil {
		return defaultValue
	}
	return value
}
