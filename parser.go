package jsonparser

import "strconv"

func ToString(v interface{}) string {
	var val string = ""
	switch v.(type) {
	case int:
		val = strconv.Itoa(v.(int))
	case int64:
		val = strconv.FormatInt(v.(int64), 10)
	case float64:
		val = strconv.FormatFloat(v.(float64), 'f', 2, 64)
	default:
		val = v.(string)
	}

	return val
}
