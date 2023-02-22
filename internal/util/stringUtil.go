package util

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
)

func Md5(source string) string {
	md5h := md5.New()
	md5h.Write([]byte(source))
	return hex.EncodeToString(md5h.Sum(nil))
}

func UUID() string {
	return uuid.New().String()
}

func Substring(str string, start, end int) string {
	var r = []rune(str)
	len := len(r)
	if start < 0 || end > len || start > end {
		return ""
	}
	if start == 0 && end == len {
		return str
	}
	return string(r[start:end])
}
