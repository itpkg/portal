package cms_test

import (
	"testing"

	"github.com/itpkg/portal/cms"
)

func TestMarkdown(t *testing.T) {
	hello := `
Title
---
### aaa

#### bbb

![alt text](demo.png)

	`

	t.Logf(string(cms.Md2Hm([]byte(hello))))
}
