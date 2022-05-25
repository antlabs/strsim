package similarity

import (
	"encoding/base64"
	"reflect"
	"unsafe"
)

const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func StringToBytes(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}

// Base64Encode encodes a byte slice to a base64 string.
func Base64Encode(s string) string {
	base := base64.NewEncoding(base64Table)
	bytes := StringToBytes(s)
	return base.EncodeToString(bytes)
}

// StrToStrs 字符串转化字符数组
func StrToStrs(s string) []string {
	base := make([]string, 0)
	for i := 0; i < len(s); i++ {
		base = append(base, string(s[i]))
	}
	return base
}
