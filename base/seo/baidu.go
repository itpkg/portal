package seo

import (
	"fmt"
)

func BaiduVerify(code string) (name string, body string) {
	return fmt.Sprintf("baidu_verify_%s.html", code), code
}
