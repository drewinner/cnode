package common

import (
	"github.com/Tang-RoseChild/mahonia"
)

// GBK编码转换为UTF8
func GBK2UTF8(s string) (string, bool) {
	dec := mahonia.NewDecoder("gbk")

	return dec.ConvertStringOK(s)
}
