package seo

import (
	"fmt"
)

func GoogleVerify(code string) (name string, body string) {
	return fmt.Sprintf("google%s.html", code), fmt.Sprintf("google-site-verification: google%s.html", code)
}
