package tpl_test

import (
	"os"
	"testing"

	"github.com/itpkg/portal/base/tpl"
)

const view = "../../views/layout.html"

func TestDump(t *testing.T) {
	if err := tpl.Dump(os.Stdout, view, tpl.Model{
		Lang:        "en",
		Url:         "/index.html",
		Author:      "aaa@aaa.com",
		Keywords:    "k1, k2",
		Description: "ddd",
		Title:       "title",
		Body:        "<p>Body</p>",
	}); err != nil {
		t.Errorf("bad in dump: %v", err)
	}
}
