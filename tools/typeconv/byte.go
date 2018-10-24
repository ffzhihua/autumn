package typeconv

import "unsafe"

func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
