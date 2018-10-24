package typeconv

import (
	"strconv"
)

//interface转为string
func Interface2String(inter interface{}) string {
	tempStr := ""
	switch inter.(type) {
	case string:
		tempStr = inter.(string)
		break
	case float64:
		tempStr = strconv.FormatFloat(inter.(float64), 'f', -1, 64)
		break
	case int64:
		tempStr = strconv.FormatInt(inter.(int64), 10)
		break
	case int:
		tempStr = strconv.Itoa(inter.(int))
		break
	}
	return tempStr
}

