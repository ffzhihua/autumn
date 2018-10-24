package typeconv

import (
	"strconv"
	"unsafe"
)

//类型转换  string to bytes
func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func StrToInt(s string) (i int, err error) {
	i, err = strconv.Atoi(s)

	return
}

func StrToInt64(s string) (i int64, err error) {

	i, err = strconv.ParseInt(s, 10, 0)
	return
}

func StrToFloat32(s string) (f float64, err error) {

	f, err = strconv.ParseFloat(s, 32)
	return
}

func StrToFloat64(s string) (f float64, err error) {

	f, err = strconv.ParseFloat(s, 64)
	return
}
func ArrToInterface(t []string) []interface{} {
	s := make([]interface{}, len(t))
	for i, v := range t {
		s[i] = v
	}
	return s
}

